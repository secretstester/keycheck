package main

import (
	"fmt"
	"os"
	"regexp"
)

func matches(args []string) (m []string) {
	patterns := []string{
		"^(A3T[A-Z0-9]|AKIA|AGPA|AIDA|AROA|AIPA|ANPA|ANVA|ASIA)[A-Z0-9]{16}$",
	}
	for _, pattern := range patterns {
		re, _ := regexp.Compile(pattern)
		for _, arg := range args {
			if re.MatchString(arg) {
				m = append(m, arg)
			}
		}
	}
	return m
}

func main() {
	for _, m := range matches(os.Args[1:]) {
		fmt.Println(m)
	}
}
