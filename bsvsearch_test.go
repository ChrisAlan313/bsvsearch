package main

import (
	"reflect"
	"testing"
)

func Test_createVerseFromLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name      string
		args      args
		wantVerse verse
	}{
		{
			name:      "createVerseFromLine creates Verse struct from string",
			args:      args{line: "Gen|1|1|In principio creavit Deus c&#230;lum et terram."},
			wantVerse: verse{book: "Gen", chapter: 1, number: 1, content: "In principio creavit Deus c&#230;lum et terram."},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVerse := createVerseFromLine(tt.args.line)
			if !reflect.DeepEqual(gotVerse, tt.wantVerse) {
				t.Errorf("createVerseFromLine() = %v, want %v", gotVerse, tt.wantVerse)
			}
		})
	}
}

func Test_loadVersesFromFile(t *testing.T) {
	type args struct {
		bibleFileLocation string
	}
	tests := []struct {
		name       string
		args       args
		wantVerses []string
		wantErr    bool
	}{
		{
			name:       "Successfully loaded",
			args:       args{bibleFileLocation: "./vuldat_test.txt"},
			wantVerses: []string{"Gen|1|1|In principio creavit Deus c&#230;lum et terram.", "Gen|1|2|Terra autem erat inanis et vacua, et tenebr&#230; erant super faciem abyssi: et spiritus Dei ferebatur super aquas.", "Gen|1|3|Dixitque Deus: Fiat lux. Et facta est lux.", "Gen|1|4|Et vidit Deus lucem quod esset bona: et divisit lucem a tenebris.", "Gen|1|5|Appellavitque lucem Diem, et tenebras Noctem: factumque est vespere et mane, dies unus.", "Gen|1|6|Dixit quoque Deus: Fiat firmamentum in medio aquarum: et dividat aquas ab aquis.", "Gen|1|7|Et fecit Deus firmamentum, divisitque aquas, qu&#230; erant sub firmamento, ab his, qu&#230; erant super firmamentum. Et factum est ita.", "Gen|1|8|Vocavitque Deus firmamentum, C&#230;lum: et factum est vespere et mane, dies secundus.", "Gen|1|9|Dixit vero Deus: Congregentur aqu&#230;, qu&#230; sub c&#230;lo sunt, in locum unum: et appareat arida. Et factum est ita.", "Gen|1|10|Et vocavit Deus aridam Terram, congregationesque aquarum appellavit Maria. Et vidit Deus quod esset bonum."},
			wantErr:    false,
		},
		{
			name:       "File not found",
			args:       args{bibleFileLocation: "./someNonExistentFile.txt"},
			wantVerses: []string{},
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotVerses, err := loadVersesFromFile(tt.args.bibleFileLocation)
			if (err != nil) != tt.wantErr {
				t.Errorf("loadVersesFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotVerses, tt.wantVerses) {
				t.Errorf("loadVersesFromFile() = %v, want %v", gotVerses, tt.wantVerses)
			}
		})
	}
}
