package solutions

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func readData(lines []string) []map[string]string {
	var res []map[string]string
	passportFields := make(map[string]string)
	for _, line := range lines {
		if line == "" {
			res = append(res, passportFields)
			passportFields = make(map[string]string)
			continue
		}
		items := strings.Fields(line)
		for _, item := range items {
			passportFields[item[0:3]] = item[4:]
		}
	}
	res = append(res, passportFields)
	return res
}

//allPassportFields := [...]string{"byr",	"iyr", "eyr", "hgt", "hcl", "ecl", "pid","cid"}
func passportIsValid(passport map[string]string) bool {
	if len(passport) < 7 {
		return false
	}
	for _, field := range [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		if _, ok := passport[field]; !ok {
			return false
		}
	}
	return true
}

func is4DigitNumInRange(strNum string, min, max int64) bool {
	if len(strNum) != 4 {
		return false
	}
	var (
		i   int64
		err error
	)

	if i, err = strconv.ParseInt(strNum, 10, 64); err != nil {
		return false
	}
	return i >= min && i <= max
}

//four digits; at least 1920 and at most 2002.
func isBirthYearValid(strNum string) bool {
	return is4DigitNumInRange(strNum, 1920, 2002)
}

//four digits; at least 2010 and at most 2020.
func isIssueYearValid(strNum string) bool {
	return is4DigitNumInRange(strNum, 2010, 2020)
}

//four digits; at least 2020 and at most 2030.
func isExpirationYearValid(strNum string) bool {
	return is4DigitNumInRange(strNum, 2020, 2030)
}

func isHeightValid(strHeight string) bool {
	strLen := len(strHeight)
	if strLen < 3 {
		return false
	}
	var (
		i   int64
		err error
	)
	dim := strHeight[strLen-2:]
	if dim == "cm" {
		if i, err = strconv.ParseInt(strHeight[:strLen-2], 10, 64); err != nil {
			return false
		}
		return i >= 150 && i <= 193
	} else if dim == "in" {
		if i, err = strconv.ParseInt(strHeight[:strLen-2], 10, 64); err != nil {
			return false
		}
		return i >= 59 && i <= 76
	} else {
		return false
	}
}

//# followed by exactly six characters 0-9 or a-f
var validHairColor = regexp.MustCompile(`^#([0-9|a-f]){6}$`)

//exactly one of: amb blu brn gry grn hzl oth.
var validEyeColor = regexp.MustCompile(`(amb|blu|brn|gry|grn|hzl|oth)`)

//a nine-digit number, including leading zeroes.
var validPassportID = regexp.MustCompile(`^([0-9]){9}$`)

func isHairColorValid(hairColor string) bool {
	return validHairColor.MatchString(hairColor)
}

func isEyeColorValid(eyeColor string) bool {
	return validEyeColor.MatchString(eyeColor)
}

func isPassportIDValid(passportID string) bool {
	return validPassportID.MatchString(passportID)
}

func isCountryIDValid(countryID string) bool {
	return true
}

type passportFieldCheck func(string) bool

var passportFieldValidator = map[string]passportFieldCheck{
	"byr": isBirthYearValid,
	"iyr": isIssueYearValid,
	"eyr": isExpirationYearValid,
	"hgt": isHeightValid,
	"hcl": isHairColorValid,
	"ecl": isEyeColorValid,
	"pid": isPassportIDValid,
	"cid": isCountryIDValid,
}

func passportIsValidWithValues(passport map[string]string) bool {
	if len(passport) < 7 {
		return false
	}
	for _, field := range [...]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
		var (
			val string
			ok  bool
		)
		if val, ok = passport[field]; !ok {
			//fmt.Printf("Field %s is missing\n", field)
			return false
		}
		if !passportFieldValidator[field](val) {
			fmt.Printf("Field %s has invalid value: %s\n", field, val)
			return false
		}
	}
	return true
}
