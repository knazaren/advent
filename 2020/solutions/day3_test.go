package solutions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_readInputAsForest(t *testing.T) {
	treeLines, lineSize := readInputAsForest("../test_data/day3_input.dat")
	assert.Equal(t, 4, len(treeLines))
	assert.Equal(t, lineSize, 11)
	assert.Equal(t, 1, countTrees(1, 3, treeLines, lineSize))
}

func TestDay3P1(t *testing.T) {
	treeLines, triesInLine := readInputAsForest("../input/31.dat")
	fmt.Println(countTrees(1, 3, treeLines, triesInLine))
}

func TestDay3P2(t *testing.T) {
	slopes := [5][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}
	// Right 1, down 1.
	// Right 3, down 1.
	// Right 5, down 1.
	// Right 7, down 1.
	// Right 1, down 2.

	treeLines, triesInLine := readInputAsForest("../input/31.dat")
	mult := 1
	for _, slope := range slopes {
		mult *= countTrees(slope[1], slope[0], treeLines, triesInLine)
	}

	fmt.Printf("Day3 Part 2 result: %d\n", mult)
}
