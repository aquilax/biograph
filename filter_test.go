package biograph

import (
	"testing"
	"time"
)

func d(date string) time.Time {
	time, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	return time
}

func TestNewBetween(t *testing.T) {
	type args struct {
		from  time.Time
		to    time.Time
		event LifeEvent
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Same dates returns true",
			args{
				d("2017-01-01"),
				d("2017-01-01"),
				NewItem("test", d("2017-01-01"), d("2017-01-01"), nil),
			},
			true,
		},
		{
			"Outside of interval date returns false",
			args{
				d("2017-01-01"),
				d("2017-01-01"),
				NewItem("test", d("2017-02-01"), d("2017-02-01"), nil),
			},
			false,
		},
		{
			"Outside of interval date returns false",
			args{
				d("2017-02-01"),
				d("2017-02-01"),
				NewItem("test", d("2017-01-01"), d("2017-03-01"), nil),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := NewBetween(tt.args.from, tt.args.to)

			if got := filter(tt.args.event); got != tt.want {
				t.Errorf("NewBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
