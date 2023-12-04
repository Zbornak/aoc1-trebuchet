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
)

func main() {
	// import calibration document
	content, err := os.ReadFile("calibration-document.txt")
	if err != nil {
		log.Fatal(err)
	}

	var filteredContent string

	// remove letter runes from calibration document
	for _, code := range content {
		if code < 64 || code > 122 {
			filteredContent += string(code)
		}
	}

	fmt.Println(filteredContent)
}
