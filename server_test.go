package main

import "testing"

func Test_withHtmlBoilerplate(t *testing.T) {
	type args struct {
		verses []Verse
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "",
			args: args{
				[]Verse{
					{"Gen", 1, 1, "In principio creavit Deus cælum et terram."},
					{"Gen", 1, 2, "Terra autem erat inanis et vacua, et tenebræ erant super faciem abyssi: et spiritus Dei ferebatur super aquas."}},
			},
			want: "<html><head><title>BSV Search Results</title></head><body><h1>BSV Search Results</h1><p>In principio creavit Deus cælum et terram.<br/>Terra autem erat inanis et vacua, et tenebræ erant super faciem abyssi: et spiritus Dei ferebatur super aquas.<br/></p></body><html>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := withHtmlBoilerplate(tt.args.verses); got != tt.want {
				t.Errorf("withHtmlBoilerplate() = %v, want %v", got, tt.want)
			}
		})
	}
}
