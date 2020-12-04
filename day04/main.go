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
	//fmt.Println(validatePassList(passes, isPresent))
	fmt.Println(validatePassList(passes, isPresentAndValid))
	//fmt.Println(isPresentAndValid(passes[1]))
}

func validatePassList(passes []pass, fn validateFn) (count int) {
	for _, pass := range passes {
		if fn(pass) {
			count++
			for _, field := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
				fmt.Printf("%v\t", pass.content[field])
			}
			fmt.Println(count)
		}
	}
	return count
}

func isPresent(p pass) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	//cid }
	for _, field := range fields {
		if _, ok := p.content[field]; !ok {
			//fmt.Println(p.content)
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

func isValid(p pass) (bool, error) {
	//byr (Birth Year) - four digits; at least 1920 and at most 2002.
	byr, err := strconv.Atoi(p.content["byr"])
	if err != nil {
		return false, err
	}
	if byr < 1920 || byr > 2002 {
		//fmt.Printf("failed byr %v\n", byr)
		return false, nil
	}
	//fmt.Printf("passed byr %v\n", byr)
	//iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	iyr, err := strconv.Atoi(p.content["iyr"])
	if err != nil {
		return false, err
	}
	if iyr < 2010 || iyr > 2020 {
		//fmt.Printf("failed iyr %v\n", iyr)
		return false, nil
	}
	//fmt.Printf("passed iyr %v\n", iyr)
	//eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	eyr, err := strconv.Atoi(p.content["eyr"])
	if err != nil {
		return false, err
	}
	if eyr < 2020 || eyr > 2030 {
		//fmt.Printf("failed eyr %v\n", eyr)
		return false, nil
	}
	//fmt.Printf("passed eyr %v\n", eyr)
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
			//fmt.Printf("failed hgt %v %v\n", hgt, unit)
			return false, nil
		}
		//If in, the number must be at least 59 and at most 76.
	} else if unit == "in" {
		if hgt < 59 || hgt > 76 {
			//fmt.Printf("failed hgt %v %v\n", hgt, unit)
			return false, nil
		}
	}
	//fmt.Printf("passed hgt %v %v\n", hgt, unit)
	//hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	hcl := p.content["hcl"]
	match, err := regexp.MatchString("^#[0-9a-f]{6}$", hcl)
	if err != nil {
		return false, err
	}
	if !match {
		//fmt.Printf("failed hcl %v\n", hcl)
		return false, nil
	}
	//fmt.Printf("passed hcl %v\n", hcl)
	//ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	ecl := p.content["ecl"]
	validEyeColors := map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
	if _, ok := validEyeColors[ecl]; !ok {
		//fmt.Printf("failed ecl %v\n", ecl)
		return false, nil
	}
	//fmt.Printf("passed ecl %v\n", ecl)
	//pid (Passport ID) - a nine-digit number, including leading zeroes.
	pid := p.content["pid"]
	match, err = regexp.MatchString("^[0-9]{9}$", pid)
	if err != nil {
		return false, err
	}
	if !match {
		//fmt.Printf("failed pid %v\n", pid)
		return false, nil
	}
	//fmt.Printf("passed pid %v\n", pid)
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
	//fmt.Println(passes[len(passes)-1])
	return passes, nil
}
