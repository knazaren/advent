package solutions

import (
	"testing"

	"github.com/knazaren/advent/utils"
)

func Test_countIncreasing(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "problemStatemnt",
			args: args{fileName: "../inputs/day1_test.txt"},
			want: 7,
		},
		{
			name: "day1_1",
			args: args{fileName: "../inputs/day1_1.txt"},
			want: 1529,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			numbers := utils.ReadInputAsInts(tt.args.fileName)
			if got := countIncreasing(numbers); got != tt.want {
				t.Errorf("countIncreasing() = %v, want %v", got, tt.want)
			}
		})
	}
}
