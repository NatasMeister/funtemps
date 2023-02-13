package conv

import (
//	"reflect"
	"testing"
	"math"
)

/*
*
	Mal for testfunksjoner
	Du skal skrive alle funksjonene basert på denne malen
	For alle konverteringsfunksjonene (tilsammen 6)
	kan du bruke malen som den er (du må selvsagt endre
	funksjonsnavn og testverdier)
*/


func TestFarhenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
		{input: 32, want: 0},
	}

	tolerance := 0.01

	for _, tc := range tests {
		got := FarhenheitToCelsius(tc.input)
//		if !reflect.DeepEqual(tc.want, got) {
		if math.Abs(tc.want-got) > tolerance {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 0},
		{input: 70, want: -203.15},
	}

	tolerance := 0.01

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if math.Abs(tc.want-got) > tolerance {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 273.15},
		{input: 70, want: 343.15},
	}

	tolerance := 0.01

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if math.Abs(tc.want-got) > tolerance {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestCelsiusToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 32},
		{input: 56.67, want: 134},
	}

	tolerance := 0.01

	for _, tc := range tests {
		got := CelsiusToFarhenheit(tc.input)
		if math.Abs(tc.want-got) > tolerance {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestKelvinToFarhenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 255.37, want: 0},
		{input: 1000, want: 1340.33},
	}

	tolerance := 0.01

	for _, tc := range tests {
		got := KelvinToFarhenheit(tc.input)
		if math.Abs(tc.want-got) > tolerance {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFarhenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 255.37},
		{input: 1340.33, want: 1000},
	}

	tolerance := 0.01

	for _, tc := range tests {
		got := FarhenheitToKelvin(tc.input)
		if math.Abs(tc.want-got) > tolerance {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}