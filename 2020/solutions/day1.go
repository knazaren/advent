package solutions

import (
	"errors"

	"github.com/knazaren/advent/utils"
)

func sum2020() (term1, term2 int64, err error) {
	input := utils.ReadInputAsInts("../input/11.dat")
	for i, term1 := range input {
		for _, term2 := range input[i:] {
			if term1+term2 == 2020 {
				return term1, term2, nil
			}
		}
	}
	return 0, 0, errors.New("can't find terms")
}

func sum2020_2() (term1, term2, term3 int64, err error) {
	input := utils.ReadInputAsInts("../input/11.dat")
	for i, term1 := range input {
		for ii, term2 := range input[i:] {
			if term1+term2 < 2020 {
				for _, term3 := range input[ii:] {
					if term1+term2+term3 == 2020 {
						return term1, term2, term3, nil
					}
				}
			}
		}
	}
	return 0, 0, 0, errors.New("can't find terms")
}
