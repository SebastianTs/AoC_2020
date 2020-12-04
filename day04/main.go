package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type pass struct {
	content map[string]string
}

type validateFn func(p pass) bool

func main() {
	passes, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(validatePassList(passes, isPresent))
	fmt.Println(validatePassList(passes, isPresentAndValid))
	fmt.Println(validatePassList(passes, isPresentAndValidRegex))
}

func validatePassList(passes []pass, fn validateFn) (count int) {
	for _, pass := range passes {
		if fn(pass) {
			count++
		}
	}
	return count
}

func isPresent(p pass) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	for _, field := range fields {
		if _, ok := p.content[field]; !ok {
			return false
		}
	}
	return true
}

func isPresentAndValid(p pass) bool {
	if !isPresent(p) {
		return false
	}
	result, err := isValid(p)
	if err != nil {
		fmt.Printf("error: %v %v\n", err, p.content)
	}
	return result

}

func isPresentAndValidRegex(p pass) bool {
	if !isPresent(p) {
		return false
	}
	result, err := isValidRegex(p)
	if err != nil {
		fmt.Printf("error: %v %v\n", err, p.content)
	}
	return result

}

func isValidRegex(p pass) (bool, error) {
	validFieldRegex := map[string]string{
		"byr": "^(19[2-9][0-9]|200[0-2])$",
		"iyr": "^20(1[0-9]|20)$",
		"eyr": "^20(2[0-9]|30)$",
		"hgt": "^(1[5-8][0-9]cm|19[0-3]cm|59in|6[0-9]in|7[0-6]in)",
		"hcl": "^#[0-9a-f]{6}",
		"ecl": "^(amb|blu|brn|gry|grn|hzl|oth){1}$",
		"pid": "^[0-9]{9}$",
	}
	for k, v := range p.content {
		match, err := regexp.MatchString(validFieldRegex[k], v)
		if err != nil {
			return false, err
		}
		if !match {
			return false, nil
		}
	}
	return true, nil
}

func isValid(p pass) (bool, error) {
	//byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byr, err := strconv.Atoi(p.content["byr"])
	if err != nil {
		return false, err
	}
	if byr < 1920 || byr > 2002 {
		return false, nil
	}
	//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr, err := strconv.Atoi(p.content["iyr"])
	if err != nil {
		return false, err
	}
	if iyr < 2010 || iyr > 2020 {
		return false, nil
	}
	//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr, err := strconv.Atoi(p.content["eyr"])
	if err != nil {
		return false, err
	}
	if eyr < 2020 || eyr > 2030 {
		return false, nil
	}
	//hgt (Height) - a number followed by either cm or in:
	var hgt int
	var unit string
	_, err = fmt.Sscanf(p.content["hgt"], "%d%s", &hgt, &unit)
	if err != nil {
		return false, err
	}
	if unit != "cm" && unit != "in" {
		return false, nil
		//If cm, the number must be at least 150 and at most 193.
	} else if unit == "cm" {
		if hgt < 150 || hgt > 193 {
			return false, nil
		}
		//If in, the number must be at least 59 and at most 76.
	} else if unit == "in" {
		if hgt < 59 || hgt > 76 {
			return false, nil
		}
	}
	//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hcl := p.content["hcl"]
	match, err := regexp.MatchString("^#[0-9a-f]{6}$", hcl)
	if err != nil {
		return false, err
	}
	if !match {
		return false, nil
	}
	//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	ecl := p.content["ecl"]
	validEyeColors := map[string]struct{}{"amb": {}, "blu": {}, "brn": {}, "gry": {}, "grn": {}, "hzl": {}, "oth": {}}
	if _, ok := validEyeColors[ecl]; !ok {
		return false, nil
	}
	/*
		match, err = regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth){1}$", ecl)
		if err != nil {
			return false, err
		}
		if !match {
			return false, nil
		}
		return true, nil
	*/
	//pid (Passport ID) - a nine-digit number, including leading zeroes.
	pid := p.content["pid"]
	match, err = regexp.MatchString("^[0-9]{9}$", pid)
	if err != nil {
		return false, err
	}
	if !match {
		return false, nil
	}
	return true, nil
}

func readInput(filename string) (passes []pass, err error) {
	passes = make([]pass, 0)
	p := pass{}
	p.content = make(map[string]string)

	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return passes, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 2 {
			splits := strings.Split(line, " ")
			for i := 0; i < len(splits); i++ {
				entries := strings.Split(splits[i], ":")
				p.content[entries[0]] = entries[1]
			}
		} else {
			passes = append(passes, p)
			p = pass{}
			p.content = make(map[string]string)
		}
	}
	passes = append(passes, p)
	return passes, nil
}
