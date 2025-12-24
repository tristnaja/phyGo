package utility

import (
	"testing"
)

func TestIntPow(t *testing.T) {
	tests := []struct {
		name        string
		baseNum     int
		exponentNum int
		want        int
		wantErr     bool
	}{
		{name: "Normal Case", baseNum: 5, exponentNum: 2, want: 25, wantErr: false},
		{name: "Exponent 0", baseNum: 10, exponentNum: 0, want: 1, wantErr: false},
		{name: "Base 0", baseNum: 0, exponentNum: 7, want: 0, wantErr: false},
		{name: "Base 1", baseNum: 1, exponentNum: 1000, want: 1, wantErr: false},
		{name: "Base -1 #1", baseNum: -1, exponentNum: 1000, want: 1, wantErr: false},
		{name: "Base -1 #2", baseNum: -1, exponentNum: 999, want: -1, wantErr: false},
		{name: "Negative Power", baseNum: -5, exponentNum: 3, want: -125, wantErr: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := intPow(test.baseNum, test.exponentNum)
			if result != test.want {
				t.Errorf("intPow(%d, %d) results in %d, but wanted %d", test.baseNum, test.exponentNum, result, test.want)
			}
		})
	}
}

func TestNotHypotenusePhyCalc(t *testing.T) {
	tests := []struct {
		name    string
		sides   map[string]int
		want    int
		wantErr bool
	}{
		{
			name: "Normal Case",
			sides: map[string]int{
				"LegA":       8,
				"Hypotenuse": 10,
			},
			want:    6,
			wantErr: false,
		},
		{
			name: "Leg Value == Hypotenuse Value",
			sides: map[string]int{
				"LegA":       8,
				"Hypotenuse": 8,
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "Leg Value > Hypotenuse Value",
			sides: map[string]int{
				"LegA":       8,
				"Hypotenuse": 6,
			},
			want:    0,
			wantErr: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := notHypotenusePhyCalc(test.sides)

			if test.wantErr {
				if err == nil {
					t.Errorf("Test expecting error, but error is nil.")
				}
			} else {
				if err != nil {
					t.Errorf("Test is not expecting error, but returns an error.")
				}
			}

			if !test.wantErr && result != test.want {
				t.Errorf("Expected %d, but got %d", test.want, result)
			}
		})
	}
}

func TestHypotenusePhyCalc(t *testing.T) {
	tests := []struct {
		name    string
		sides   map[string]int
		want    int
		wantErr bool
	}{
		{
			name: "Normal Case",
			sides: map[string]int{
				"LegA": 8,
				"LegB": 6,
			},
			want:    10,
			wantErr: false,
		},
		{
			name: "Big Number",
			sides: map[string]int{
				"LegA": 5000,
				"LegB": 5000,
			},
			want:    7071,
			wantErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := hypotenusePhyCalc(test.sides)

			if test.wantErr {
				if err == nil {
					t.Errorf("Test expecting error, but error is nil.")
				}
			} else {
				if err != nil {
					t.Errorf("Test is not expecting error, but returns an error.")
				}
			}

			if !test.wantErr && result != test.want {
				t.Errorf("Expected %d, but got %d", test.want, result)
			}
		})
	}
}
