package report

import (
	"bytes"
	"testing"
	"time"

	"github.com/aquilax/biograph"
)

func TestLifeCalendarOwnTime_Generate(t *testing.T) {
	type fields struct {
		startDate time.Time
		endDate   time.Time
		output    *bytes.Buffer
	}
	type args struct {
		events biograph.Events
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			"",
			fields{
				time.Date(2000, 1, 1, 1, 1, 1, 1, time.Local),
				time.Date(2000, 1, 1, 1, 1, 1, 1, time.Local),
				&bytes.Buffer{},
			},
			args{biograph.Events{
				biograph.NewHome("City", "Country", d("2000-02-01"), d("2022-01-01"), nil),
				biograph.NewHome("City", "Country", d("2000-02-02"), d("2022-01-01"), nil),
				biograph.NewHome("City", "Country", d("2000-03-02"), d("2022-01-01"), nil),
			}},
			`    01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30 31 32 33 34 35 36 37 38 39 40 41 42 43 44 45 46 47 48 49 50 51 52
  0  -- -- -- -- 02 -- -- -- 01 -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- -- --
`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LifeCalendarOwnTime{
				startDate: tt.fields.startDate,
				endDate:   tt.fields.endDate,
				output:    tt.fields.output,
			}
			if err := l.Generate(tt.args.events); (err != nil) != tt.wantErr {
				t.Errorf("LifeCalendarOwnTime.Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
			output := tt.fields.output.String()
			if tt.wantOutput != output {
				t.Errorf("LifeCalendarText.Generate() output = %v, wants %v", output, tt.wantOutput)
			}
		})
	}
}
