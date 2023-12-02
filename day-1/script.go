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

// The newly-improved calibration document consists of lines of text;
//each line originally contained a specific calibration value that the Elves now need to recover.
// On each line, the calibration value can be found by combining the first digit and the last digit
// (in that order) to form a single two-digit number.

// For example:

// 1abc2 -> 12
// pqr3stu8vwx -> 38
// a1b2c3d4e5f -> 15
// treb7uchet -> 77

func main() {
	file, err := fs.ReadFile(os.DirFS(filepath.Dir(".")), "input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	calibrationSum := 0

	for _, line := range lines {
		first, second := 0, 0

		for i := 0; i < len(line); i += 1 {
			if line[i] >= 48 && line[i] <= 57 {
				num, err := strconv.Atoi(string(line[i]))
				if err != nil {
					log.Fatal(err)
				}
				if first == 0 {
					first = num
					second = num
				} else {
					second = num
				}
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
