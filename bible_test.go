package main

import (
	"reflect"
	"testing"
)

func TestNewBible(t *testing.T) {
	type args struct {
		translation string
	}
	tests := []struct {
		name string
		args args
		want bible
	}{
		{
			name: "creates bible empt struct",
			args: args{"some translation"},
			want: bible{
				translation: "some translation", verses: []verse{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBible(tt.args.translation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBible() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name        string
		args        args
		wantBook    string
		wantChapter int
		wantNumber  int
		wantContent string
	}{
		{
			name:        "parses line of bible txt input",
			args:        args{line: "Gen|1|1|In principio creavit Deus cælum et terram."},
			wantBook:    "Gen",
			wantChapter: 1,
			wantNumber:  1,
			wantContent: "In principio creavit Deus cælum et terram.",
		},
		{
			name:        "parses line of bible txt input with '[' and '<br>'",
			args:        args{line: "Joh|1|1|[In principio erat Verbum,<BR> et Verbum erat apud Deum,<BR> et Deus erat Verbum.<BR>"},
			wantBook:    "Joh",
			wantChapter: 1,
			wantNumber:  1,
			wantContent: "[In principio erat Verbum,<BR> et Verbum erat apud Deum,<BR> et Deus erat Verbum.<BR>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBook, gotChapter, gotNumber, gotContent := parseLine(tt.args.line)
			if gotBook != tt.wantBook {
				t.Errorf("parseLine() gotBook = %v, want %v", gotBook, tt.wantBook)
			}
			if gotChapter != tt.wantChapter {
				t.Errorf("parseLine() gotChapter = %v, want %v", gotChapter, tt.wantChapter)
			}
			if gotNumber != tt.wantNumber {
				t.Errorf("parseLine() gotNumber = %v, want %v", gotNumber, tt.wantNumber)
			}
			if gotContent != tt.wantContent {
				t.Errorf("parseLine() gotContent = %v, want %v", gotContent, tt.wantContent)
			}
		})
	}
}

func Test_bible_Load(t *testing.T) {
	type args struct {
		content []string
	}
	tests := []struct {
		name    string
		b       bible
		args    args
		want    bible
		wantErr bool
	}{
		{
			name: "loads from content slice of strings representing lines of a file",
			b:    bible{translation: "some translation", verses: make([]verse, 0, 36000)},
			args: args{[]string{
				"Gen|1|1|In principio creavit Deus cælum et terram.",
				"Gen|1|2|Terra autem erat inanis et vacua, et tenebræ erant super faciem abyssi: et spiritus Dei ferebatur super aquas.",
				"Ma2|1|21|Et jussit eos haurire, et afferre sibi: et sacrificia quæ imposita erant, jussit sacerdos Nehemias aspergi ipsa aqua: et ligna, et quæ erant superposita.",
				"Ma2|1|22|Utque hoc factum est, et tempus affuit quo sol refulsit, qui prius erat in nubilo, accensus est ignis magnus, ita ut omnes mirarentur.",
				"Joh|1|1|[In principio erat Verbum,<BR> et Verbum erat apud Deum,<BR> et Deus erat Verbum.<BR>",
				"Joh|1|2|Hoc erat in principio apud Deum.<BR>"},
			},
			want: bible{
				translation: "some translation",
				verses: []verse{
					{"Gen", 1, 1, "In principio creavit Deus cælum et terram."},
					{"Gen", 1, 2, "Terra autem erat inanis et vacua, et tenebræ erant super faciem abyssi: et spiritus Dei ferebatur super aquas."},
					{"Ma2", 1, 21, "Et jussit eos haurire, et afferre sibi: et sacrificia quæ imposita erant, jussit sacerdos Nehemias aspergi ipsa aqua: et ligna, et quæ erant superposita."},
					{"Ma2", 1, 22, "Utque hoc factum est, et tempus affuit quo sol refulsit, qui prius erat in nubilo, accensus est ignis magnus, ita ut omnes mirarentur."},
					{"Joh", 1, 1, "[In principio erat Verbum,<BR> et Verbum erat apud Deum,<BR> et Deus erat Verbum.<BR>"},
					{"Joh", 1, 2, "Hoc erat in principio apud Deum.<BR>"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.b.Load(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("bible.Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bible.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
