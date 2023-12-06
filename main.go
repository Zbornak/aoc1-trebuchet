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
	calDoc = replaceNums(calDoc, "one", "1")
	calDoc = replaceNums(calDoc, "two", "2")
	calDoc = replaceNums(calDoc, "three", "3")
	calDoc = replaceNums(calDoc, "four", "4")
	calDoc = replaceNums(calDoc, "five", "5")
	calDoc = replaceNums(calDoc, "six", "6")
	calDoc = replaceNums(calDoc, "seven", "7")
	calDoc = replaceNums(calDoc, "eight", "8")
	calDoc = replaceNums(calDoc, "nine", "9")

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
