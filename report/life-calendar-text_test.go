package report

import (
	"bytes"
	"testing"
	"time"

	"github.com/aquilax/biograph"
)

func TestLifeCalendarText_Generate(t *testing.T) {
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
				time.Date(2001, 1, 1, 1, 1, 1, 1, time.Local),
				&bytes.Buffer{},
			},
			args{biograph.Events{
				biograph.NewHome("City", "Country", d("2000-02-01"), d("2022-01-01"), nil),
				biograph.NewHome("City", "Country", d("2000-02-02"), d("2022-01-01"), nil),
				biograph.NewHome("City", "Country", d("2000-03-02"), d("2022-01-01"), nil),
			}},
			`2000  00 00 00 00 02 00 00 00 01 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
2001  00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00
`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LifeCalendarText{
				startDate: tt.fields.startDate,
				endDate:   tt.fields.endDate,
				output:    tt.fields.output,
			}
			if err := l.Generate(tt.args.events); (err != nil) != tt.wantErr {
				t.Errorf("LifeCalendarText.Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
			output := tt.fields.output.String()
			if tt.wantOutput != output {
				t.Errorf("LifeCalendarText.Generate() output = %v, wants %v", output, tt.wantOutput)
			}
		})
	}
}
