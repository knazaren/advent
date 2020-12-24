package solutions

import (
	"strings"

	"github.com/knazaren/advent/utils"
)

func readInputAsForest(filename string) ([]string, int) {
	got := utils.ReadInput(filename)
	items := strings.Fields(got)

	return items, len(items[0])
}

func getNextCoord(maxV, maxH, stepV, stepH, v, h int) (int, int, bool) {
	newH := h + stepH
	if newH >= maxH {
		newH = newH - maxH
	}
	newV := v + stepV
	if newV == maxV-1 {
		return newV, newH, true
	}
	return newV, newH, false
}

func countTrees(stepV, stepH int, treeLines []string, triesInLine int) int {
	treeCount := 0
	v, h, stopNow := getNextCoord(len(treeLines), triesInLine, stepV, stepH, 0, 0)

	for {
		if treeLines[v][h] == 35 {
			treeCount++
		}
		if stopNow {
			break
		}
		v, h, stopNow = getNextCoord(len(treeLines), triesInLine, stepV, stepH, v, h)
	}
	return treeCount
}
