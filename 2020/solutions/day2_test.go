package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadDay2Pretest(t *testing.T) {
	got := readInputAsPasswords("../test_data/day2_input.dat")
	assert.Equal(t, 3, len(got))
	assert.Equal(t, true, got[0].match())
	assert.Equal(t, false, got[1].match())
	assert.Equal(t, true, got[2].match())
}

func TestDay2(t *testing.T) {
	got := readInputAsPasswords("../input/21.dat")
	matched := 0

	for _, pass := range got {
		if pass.match2() {
			matched++
		}
	}

	fmt.Println(matched)
}
