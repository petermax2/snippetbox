package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name  string
		input time.Time
		want  string
	}{
		{
			name:  "UTC",
			input: time.Date(2020, 12, 17, 10, 12, 0, 0, time.UTC),
			want:  "17 Dec 2020 at 10:12",
		},
		{
			name:  "Empty",
			input: time.Time{},
			want:  "",
		},
		{
			name:  "CET",
			input: time.Date(2020, 12, 17, 10, 0, 0, 0, time.FixedZone("CET", 1*60*60)),
			want:  "17 Dec 2020 at 09:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := humanDate(tt.input)
			if res != tt.want {
				t.Errorf("want %q; got %q", tt.want, res)
			}
		})
	}
}
