package funfacts

import (
	"reflect"
	"testing"
)

func TestGetFunFacts(t *testing.T) {
	tests := []struct {
		about string
		want  []string
	}{
		{
			about: "sun",
			want: []string{
				"Sola er ei stjerne.",
          		"Sola er nesten 5 milliarder år gammel.",
			},
		},
		{
			about: "luna",
			want: []string{
				"Månen går i bane rundt jorda.",
          		"Månen er faktisk ikke lagd av ost.",
			},
		},
		{
			about: "terra",
			want: []string{
				"Vi bor på jorda.",
          		"Jorda er kul og sånt.",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.about, func(t *testing.T) {
			if got := GetFunFacts(tt.about); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFunFacts() = %v, want %v", got, tt.want)
			}
		})
	}
}