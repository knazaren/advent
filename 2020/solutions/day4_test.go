package solutions

import (
	"fmt"
	"testing"

	"github.com/knazaren/advent/utils"
	"github.com/stretchr/testify/assert"
)

func Test_readData(t *testing.T) {
	lines := utils.ReadFileLines("../test_data/day4_input.dat")
	got := readData(lines)
	assert.Equal(t, 4, len(got))
	assert.Equal(t, got[0]["ecl"], "gry")
	assert.Equal(t, got[1]["cid"], "350")
	assert.Equal(t, got[2]["byr"], "1931")
	assert.Equal(t, got[3]["hcl"], "#cfa07d")

	matchingPassports := 0
	for _, p := range got {
		if passportIsValid(p) {
			matchingPassports++
		}
	}
	assert.Equal(t, 2, matchingPassports)
}

func TestDay4P1(t *testing.T) {
	lines := utils.ReadFileLines("../input/41.dat")
	got := readData(lines)

	matchingPassports := 0
	for _, p := range got {
		if passportIsValid(p) {
			matchingPassports++
		}
	}
	fmt.Printf("The number of matching passports: %d\n", matchingPassports)
}

func Test_is4DigitNumInRange(t *testing.T) {
	type args struct {
		strNum string
		min    int64
		max    int64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test1",
			args: args{strNum: "2002", min: 1920, max: 2002},
			want: true,
		},
		{
			name: "Test2",
			args: args{strNum: "2003", min: 1920, max: 2002},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := is4DigitNumInRange(tt.args.strNum, tt.args.min, tt.args.max); got != tt.want {
				t.Errorf("is4DigitNumInRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHeightValid(t *testing.T) {
	type args struct {
		strHeight string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test1", args: args{strHeight: "60in"}, want: true},
		{name: "Test2", args: args{"190cm"}, want: true},
		{name: "Test3", args: args{"190in"}, want: false},
		{name: "Test2", args: args{"190"}, want: false},
		{name: "171cm", args: args{"171cm"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHeightValid(tt.args.strHeight); got != tt.want {
				t.Errorf("isHeightValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isHairColorValid(t *testing.T) {
	type args struct {
		hairColor string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test1", args{"#123abc"}, true},
		{"Test2", args{"#123abz"}, false},
		{"Test3", args{"123abc"}, false},
		{"Test4", args{"#123abcd"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHairColorValid(tt.args.hairColor); got != tt.want {
				t.Errorf("isHairColorValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEyeColorValid(t *testing.T) {
	type args struct {
		eyeColor string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test1", args{"brn"}, true},
		{"Test2", args{"wat"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEyeColorValid(tt.args.eyeColor); got != tt.want {
				t.Errorf("isEyeColorValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPassportIDValid(t *testing.T) {
	type args struct {
		passportID string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{

		{"Test1", args{"000000001"}, true},
		{"Test2", args{"0123456789"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPassportIDValid(tt.args.passportID); got != tt.want {
				t.Errorf("isPassportIDValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_passportValues(t *testing.T) {
	lines := utils.ReadFileLines("../test_data/day42_input.dat")
	got := readData(lines)

	matchingPassports := 0
	for _, p := range got {
		if passportIsValidWithValues(p) {
			matchingPassports++
		}
	}

	assert.Equal(t, 4, matchingPassports)
}
func TestDay4P2(t *testing.T) {
	lines := utils.ReadFileLines("../input/42.dat")
	got := readData(lines)

	matchingPassports := 0
	for _, p := range got {
		if passportIsValidWithValues(p) {
			matchingPassports++
		}
	}
	fmt.Printf("The number of matching passports: %d\n", matchingPassports)
}
