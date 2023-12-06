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
	"strconv"
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

	// part 2: convert written numbers into digits
	calDoc = replaceNums(calDoc, "one", "o1ne")
	calDoc = replaceNums(calDoc, "two", "t2wo")
	calDoc = replaceNums(calDoc, "three", "t3ree")
	calDoc = replaceNums(calDoc, "four", "f4our")
	calDoc = replaceNums(calDoc, "five", "f5ive")
	calDoc = replaceNums(calDoc, "six", "s6ix")
	calDoc = replaceNums(calDoc, "seven", "s7even")
	calDoc = replaceNums(calDoc, "eight", "e8ight")
	calDoc = replaceNums(calDoc, "nine", "n9ine")

	// extract numbers from calDoc
	numString := extractNums(calDoc)

	// split numbers into separate slice entries
	splitNumString := strings.Fields(numString)

	// get first and last number of each line and put together
	var intSlice []int
	var firstDigit string
	var lastDigit string
	for _, digit := range splitNumString {
		firstDigit = digit[:1]
		lastDigit = digit[len(digit)-1:]
		joinedDigits := firstDigit + lastDigit

		// convert strings to ints and append to int slice
		intNum, err := strconv.Atoi(joinedDigits)
		if err != nil {
			log.Fatal(err)
		}
		intSlice = append(intSlice, intNum)
	}

	fmt.Println(intSlice)

	// add all the numbers together
	var calibrationValuesTotal int
	for i := 0; i < len(intSlice); i++ {
		calibrationValuesTotal += intSlice[i]
	}

	// print total
	fmt.Printf("The answer is %d.\n", calibrationValuesTotal)
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

func replaceNums(str string, wordToReplace string, replaceWith string) string {
	newCalDoc := strings.Replace(str, wordToReplace, replaceWith, -1)
	return newCalDoc
}
