package main

import (
	"reflect"
	"testing"
)

func TestBookSpecification_IsSatisfied(t *testing.T) {
	type args struct {
		v Verse
	}
	tests := []struct {
		name string
		b    BookSpecification
		args args
		want bool
	}{
		{
			name: "is satisfied",
			b:    BookSpecification{"Jo3"},
			args: args{Verse{"Jo3", 1, 1, "Something profound"}},
			want: true,
		},
		{
			name: "is NOT satisfied",
			b:    BookSpecification{"Jo3"},
			args: args{Verse{"Gen", 1, 1, "Something profound"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.b.IsSatisfied(tt.args.v); got != tt.want {
				t.Errorf("BookSpecification.IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChapterSpecification_IsSatisfied(t *testing.T) {
	type args struct {
		v Verse
	}
	tests := []struct {
		name string
		c    ChapterSpecification
		args args
		want bool
	}{
		{
			name: "is satisfied",
			c:    ChapterSpecification{1},
			args: args{Verse{"Gen", 1, 1, "Something profound"}},
			want: true,
		},
		{
			name: "is NOT satisfied",
			c:    ChapterSpecification{1},
			args: args{Verse{"Gen", 2, 1, "Something profound"}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.IsSatisfied(tt.args.v); got != tt.want {
				t.Errorf("ChapterSpecification.IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBetterFilter_Filter(t *testing.T) {
	type args struct {
		verses []Verse
		spec   Specification
	}
	tests := []struct {
		name string
		f    *BetterFilter
		args args
		want []Verse
	}{
		{
			name: "filters for books",
			f:    &BetterFilter{},
			args: args{
				[]Verse{
					{"Gen", 1, 1, "In principio creavit Deus cælum et terram."},
					{"Gen", 1, 2, "Terra autem erat inanis et vacua, et tenebræ erant super faciem abyssi: et spiritus Dei ferebatur super aquas."},
					{"Ma2", 1, 21, "Et jussit eos haurire, et afferre sibi: et sacrificia quæ imposita erant, jussit sacerdos Nehemias aspergi ipsa aqua: et ligna, et quæ erant superposita."},
					{"Ma2", 1, 22, "Utque hoc factum est, et tempus affuit quo sol refulsit, qui prius erat in nubilo, accensus est ignis magnus, ita ut omnes mirarentur."},
					{"Joh", 1, 1, "[In principio erat Verbum,<BR> et Verbum erat apud Deum,<BR> et Deus erat Verbum.<BR>"},
					{"Joh", 1, 2, "Hoc erat in principio apud Deum.<BR>"},
				},
				BookSpecification{"Joh"},
			},
			want: []Verse{
				{"Joh", 1, 1, "[In principio erat Verbum,<BR> et Verbum erat apud Deum,<BR> et Deus erat Verbum.<BR>"},
				{"Joh", 1, 2, "Hoc erat in principio apud Deum.<BR>"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.f.Filter(tt.args.verses, tt.args.spec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BetterFilter.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
