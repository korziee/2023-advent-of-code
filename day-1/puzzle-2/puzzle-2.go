package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

/**
Your calculation isn't quite right. It looks like some of the digits are actually
spelled out with letters: one, two, three, four, five, six, seven, eight, and nine also count as valid "digits".

Equipped with this new information, you now need to find the real first and last digit on each line. For example:

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

In this example, the calibration values are 29, 83, 13, 24, 42, 14, and 76. Adding these together produces 281.

**/

func main() {
	file, err := fs.ReadFile(os.DirFS(filepath.Dir(".")), "input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	calibrationSum := 0

	for _, line := range lines {
		firstIdx := 100
		secondIdx := -1
		first, second := 0, 0

		matchables := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

		for _, match := range matchables {
			matchIdxFirst := strings.Index(line, match)
			matchIdxLast := strings.LastIndex(line, match)

			if matchIdxFirst == -1 {
				continue
			}

			num := 0

			switch match {
			case "one":
				num = 1
			case "two":
				num = 2
			case "three":
				num = 3
			case "four":
				num = 4
			case "five":
				num = 5
			case "six":
				num = 6
			case "seven":
				num = 7
			case "eight":
				num = 8
			case "nine":
				num = 9
			default:
				conv, err := strconv.Atoi(match)
				if err != nil {
					log.Fatal(err)
				}

				num = conv
			}

			if matchIdxFirst < firstIdx {
				first = num
				firstIdx = matchIdxFirst
			}

			if matchIdxLast != -1 && matchIdxLast > secondIdx {
				second = num
				secondIdx = matchIdxLast
			}
		}

		num, err := strconv.Atoi(fmt.Sprint(first) + fmt.Sprint(second))
		if err != nil {
			log.Fatal(err)
		}

		calibrationSum += num
	}

	fmt.Printf("Calibration sum is: %v\n", calibrationSum)
}
