/*
Day 1 of Advent of Code 2023
Trebuchet?!
Author: Mark Strijdom (zbornak)
Date: 04/12/2023
*/

package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// import calibration document
	contents, err := os.ReadFile("calibration-document.txt")
	if err != nil {
		log.Fatal(err)
	}

	// convert []byte to string
	calDoc := string(contents)

	// extract numbers from calDoc
	numString := extractNums(calDoc)

	// take first and last rune from numString and convert to int
	splitNumString := strings.Fields(numString)

	fmt.Println(splitNumString)
}

func extractNums(str string) string {
	var numberString string
	for _, digit := range str {
		if digit < 64 || digit > 122 {
			numberString += string(digit)
		}
	}

	return numberString
}
