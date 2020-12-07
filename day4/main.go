package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rowNum := 0
	passportInfo := []string{}
	validPassport := 0
	for scanner.Scan() {
		row := scanner.Text()
		if row != "" {
			fields := strings.Split(row, " ")
			passportInfo = append(passportInfo, fields...)
		} else {
			if validatePassport(passportInfo) {
				validPassport++
			}
			passportInfo = []string{}
		}
		rowNum++
	}

	fmt.Println(validPassport)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func validatePassport(passportInfo []string) bool {
	requiredFields := map[string]bool{"byr": false, "iyr": false, "eyr": false, "hgt": false, "hcl": false, "ecl": false, "pid": false, "cid": true}
	if len(passportInfo) < 7 || len(passportInfo) > 8 {
		return false
	}

	for _, field := range passportInfo {
		kv := strings.Split(field, ":")

		requiredFields[kv[0]] = true

		switch kv[0] {
		case "byr":
			requiredFields[kv[0]] = validateBirthYear(kv[1])
		case "iyr":
			requiredFields[kv[0]] = validateIssueYear(kv[1])
		case "eyr":
			requiredFields[kv[0]] = validateExpirationYear(kv[1])
		case "hgt":
			requiredFields[kv[0]] = validateHeight(kv[1])
		case "hcl":
			requiredFields[kv[0]] = validateHairColor(kv[1])
		case "ecl":
			requiredFields[kv[0]] = validateEyeColor(kv[1])
		case "pid":
			requiredFields[kv[0]] = validatePassportID(kv[1])
		case "cid":
			continue
		}
	}

	for _, present := range requiredFields {
		if !present {
			return false
		}
	}

	return true
}

// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
func validateEyeColor(value string) bool {
	validValues := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	return validValues[value]
}

// byr (Birth Year) - four digits; at least 1920 and at most 2002.
func validateBirthYear(value string) bool {
	val, err := strconv.Atoi(value)
	return err == nil && val >= 1920 && val <= 2002
}

// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
func validateIssueYear(value string) bool {
	val, err := strconv.Atoi(value)
	return err == nil && val >= 2010 && val <= 2020
}

// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
func validateExpirationYear(value string) bool {
	val, err := strconv.Atoi(value)
	return err == nil && val >= 2020 && val <= 2030
}

// hgt (Height) - a number followed by either cm or in:
//   If cm, the number must be at least 150 and at most 193.
//   If in, the number must be at least 59 and at most 76.
func validateHeight(value string) bool {
	if strings.HasSuffix(value, "cm") {
		numString := strings.TrimRight(value, "cm")
		num, err := strconv.Atoi(numString)
		return err == nil && num >= 150 && num <= 193
	} else if strings.HasSuffix(value, "in") {
		numString := strings.TrimRight(value, "in")
		num, err := strconv.Atoi(numString)
		return err == nil && num >= 59 && num <= 76
	}

	return false
}

// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
func validateHairColor(value string) bool {
	matches, err := regexp.MatchString("^#[a-f0-9]{6}$", value)
	return err == nil && matches
}

// pid (Passport ID) - a nine-digit number, including leading zeroes.
func validatePassportID(value string) bool {
	matches, err := regexp.MatchString("^[0-9]{9}$", value)
	return err == nil && matches
}
