package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDay1(t *testing.T) {
	term1, term2, err := sum2020()
	assert.NoError(t, err)
	fmt.Printf("Day1 puzzle1: %d\n", term1*term2)

	term1, term2, term3, err := sum2020_2()
	assert.NoError(t, err)
	fmt.Printf("Day1 puzzle2: %d\n", term1*term2*term3)
}
