package main

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		translation  string
		fileLocation string
	}
	tests := []struct {
		name string
		args args
		want Bible
	}{
		{
			name: "creates Bible with empty struct",
			args: args{
				translation:  "some translation",
				fileLocation: "./vuldat_test.txt",
			},
			want: Bible{
				Translation: "some translation",
				Verses: []Verse{
					{"Gen", 1, 1, "In principio creavit Deus cælum et terram."},
					{"Gen", 1, 2, "Terra autem erat inanis et vacua, et tenebræ erant super faciem abyssi: et spiritus Dei ferebatur super aquas."},
					{"Ma2", 1, 21, "Et jussit eos haurire, et afferre sibi: et sacrificia quæ imposita erant, jussit sacerdos Nehemias aspergi ipsa aqua: et ligna, et quæ erant superposita."},
					{"Ma2", 1, 22, "Utque hoc factum est, et tempus affuit quo sol refulsit, qui prius erat in nubilo, accensus est ignis magnus, ita ut omnes mirarentur."},
					{"Joh", 1, 1, "[In principio erat Verbum,<BR> et Verbum erat apud Deum,<BR> et Deus erat Verbum.<BR>"},
					{"Joh", 1, 2, "Hoc erat in principio apud Deum.<BR>"},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.translation, tt.args.fileLocation); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_Bible_load(t *testing.T) {
	type args struct {
		content []string
	}
	tests := []struct {
		name    string
		b       Bible
		args    args
		want    Bible
		wantErr bool
	}{
		{
			name: "loads from content slice of strings representing lines of a file",
			b:    Bible{Translation: "some translation", Verses: make([]Verse, 0, 36000)},
			args: args{[]string{
				"Gen|1|1|In principio creavit Deus cælum et terram.",
				"Gen|1|2|Terra autem erat inanis et vacua, et tenebræ erant super faciem abyssi: et spiritus Dei ferebatur super aquas.",
				"Ma2|1|21|Et jussit eos haurire, et afferre sibi: et sacrificia quæ imposita erant, jussit sacerdos Nehemias aspergi ipsa aqua: et ligna, et quæ erant superposita.",
				"Ma2|1|22|Utque hoc factum est, et tempus affuit quo sol refulsit, qui prius erat in nubilo, accensus est ignis magnus, ita ut omnes mirarentur.",
				"Joh|1|1|[In principio erat Verbum,<BR> et Verbum erat apud Deum,<BR> et Deus erat Verbum.<BR>",
				"Joh|1|2|Hoc erat in principio apud Deum.<BR>"},
			},
			want: Bible{
				Translation: "some translation",
				Verses: []Verse{
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
			got, err := tt.b.load(tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("Bible.load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bible.load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_loadLinesFromFile(t *testing.T) {
	type args struct {
		fileLocation string
	}
	tests := []struct {
		name      string
		args      args
		wantLines []string
	}{
		{
			name: "Returns expected lines from test file",
			args: args{fileLocation: "./vuldat_test.txt"},
			wantLines: []string{
				"Gen|1|1|In principio creavit Deus cælum et terram.",
				"Gen|1|2|Terra autem erat inanis et vacua, et tenebræ erant super faciem abyssi: et spiritus Dei ferebatur super aquas.",
				"Ma2|1|21|Et jussit eos haurire, et afferre sibi: et sacrificia quæ imposita erant, jussit sacerdos Nehemias aspergi ipsa aqua: et ligna, et quæ erant superposita.",
				"Ma2|1|22|Utque hoc factum est, et tempus affuit quo sol refulsit, qui prius erat in nubilo, accensus est ignis magnus, ita ut omnes mirarentur.",
				"Joh|1|1|[In principio erat Verbum,<BR> et Verbum erat apud Deum,<BR> et Deus erat Verbum.<BR>",
				"Joh|1|2|Hoc erat in principio apud Deum.<BR>",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLines := loadLinesFromFile(tt.args.fileLocation); !reflect.DeepEqual(gotLines, tt.wantLines) {
				t.Errorf("loadLinesFromFile() = %v, want %v", gotLines, tt.wantLines)
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
			name:        "parses line of Bible txt input",
			args:        args{line: "Gen|1|1|In principio creavit Deus cælum et terram."},
			wantBook:    "Gen",
			wantChapter: 1,
			wantNumber:  1,
			wantContent: "In principio creavit Deus cælum et terram.",
		},
		{
			name:        "parses line of Bible txt input with '[' and '<br>'",
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
