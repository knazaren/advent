package solutions

import (
	"strconv"
	"strings"

	"github.com/knazaren/advent/utils"
)

type password struct {
	min, max int
	ch       byte
	payload  string
}

func (pass *password) match() bool {
	count := strings.Count(pass.payload, string(pass.ch))
	return (pass.min <= count) && (count <= pass.max)
}

func (pass *password) match2() bool {
	count := utils.BoolToInt(pass.payload[pass.min-1] == pass.ch) + utils.BoolToInt(pass.payload[pass.max-1] == pass.ch)
	return count == 1
}

//readInputAsPasswords read input as a slice of passwords
func readInputAsPasswords(filename string) []password {
	got := utils.ReadInput(filename)
	items := strings.Fields(got)
	passwordCount := len(items) / 3
	passwords := make([]password, passwordCount)
	for i := 0; i < passwordCount; i++ {
		minMaxStr := strings.Split(items[i*3], "-")
		ch := items[i*3+1][0]
		min, _ := strconv.Atoi(minMaxStr[0])
		max, _ := strconv.Atoi(minMaxStr[1])
		passwords[i] = password{min: min, max: max, ch: ch, payload: items[i*3+2]}
	}
	return passwords
}
