package main

import (
	"reflect"
	"testing"
)

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
