package report

import (
	"bytes"
	"io"
	"testing"

	"github.com/aquilax/biograph"
)

func TestMarkWhen_Generate(t *testing.T) {
	tests := []struct {
		name    string
		events  biograph.Events
		want    string
		wantErr bool
	}{
		{
			"generates markwhen timeline",
			biograph.Events{
				biograph.NewHome("City", "Country", d("2000-02-01"), d("2022-01-01"), nil),
				biograph.NewHome("City", "Country", d("2000-02-02"), d("2022-01-01"), nil),
				biograph.NewHome("City", "Country", d("2000-03-02"), d("2022-01-01"), nil),
			},
			`2000-02-01 - 2022-01-01: City #home
2000-02-02 - 2022-01-01: City #home
2000-03-02 - 2022-01-01: City #home
`,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var b bytes.Buffer
			w := io.Writer(&b)
			mw := MarkWhen{w}
			if err := mw.Generate(tt.events); (err != nil) != tt.wantErr {
				t.Errorf("MarkWhen.Generate() error = %v, wantErr %v", err, tt.wantErr)
			}
			got := b.String()
			if got != tt.want {
				t.Errorf("MarkWhen.Generate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
