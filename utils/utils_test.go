package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInput(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantInt []int64
	}{
		{name: "test1",
			args:    args{filename: "./test_data/input1.dat"},
			want:    "1\n2\n3\n4",
			wantInt: []int64{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			if got = ReadInput(tt.args.filename); got != tt.want {
				t.Errorf("ReadInput() = %v, want %v", got, tt.want)
			}
			items := strings.Fields(got)
			assert.Equal(t, 4, len(items))

			gotInt := ReadInputAsInts(tt.args.filename)
			for i, v := range gotInt {
				if v != tt.wantInt[i] {
					t.Errorf("Index %d is not matching", i)
				}
			}
		})
	}
}

func TestBoolToInt(t *testing.T) {
	assert.Equal(t, 1, BoolToInt(true))
	assert.Equal(t, 0, BoolToInt(false))
}

func TestReadFileLines(t *testing.T) {
	lines := ReadFileLines("../2020/test_data/day4_input.dat")
	assert.Equal(t, 13, len(lines))
}
