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


The engine schematic (your puzzle input) consists of a visual representation of the engine.
There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol,
even diagonally, is a "part number" and should be included in your sum.
(Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.r
...$.*....
.664.598..

In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right).
Every other number is adjacent to a symbol and so is a part number; their sum is 4361

**/

type part struct {
	num string
	row int
	col int
}

func main() {
	file, err := fs.ReadFile(os.DirFS(filepath.Dir(".")), "input")
	if err != nil {
		panic(err)
	}

	symbols := make(map[string]bool)
	var potentialParts []part

	for rowIdx, line := range strings.Split(string(file), "\n") {
		prevStringNum := ""

		for colIdx, cell := range strings.Split(line, "") {
			// fmt.Println(rowIdx, colIdx, cell)
			_, err := strconv.Atoi(cell)

			// must be either period or symbol if string to int conv failed?
			if err != nil {
				// if there was prev num, cut here
				// cut

				if prevStringNum != "" {
					potentialParts = append(potentialParts, part{
						num: prevStringNum,
						row: rowIdx,
						col: colIdx - 1,
					})

					prevStringNum = ""
				}

				// must be symbol?
				if cell != "." {
					symbols[fmt.Sprintf("%d,%d", rowIdx, colIdx)] = true
				}
			} else {
				// if not nil, append cell string val to prev num
				prevStringNum += cell
			}
		}

		// if we're at the end of the row check to see if we need to cleanup
		if prevStringNum != "" {
			potentialParts = append(potentialParts, part{
				num: prevStringNum,
				row: rowIdx,
				col: len(line) - 1,
			})

			prevStringNum = ""
		}
	}

	var parts []part

	for _, p := range potentialParts {
		// for each of the cells that we found a potential part in
		// look around it in a star pattern to find neighbours
		// if we find symbol neighbour, add to parts list and break
	thing:
		for i := len(p.num) - 1; i >= 0; i -= 1 {
			fmt.Println(p, i)

			// now iterate through cols in each row, looking for match
			for row := p.row - 1; row <= p.row+1; row += 1 {
				for col := (p.col - 1) - i; col <= p.col+1; col += 1 {
					// fmt.Println("looking at row x col", row, col)
					if symbols[fmt.Sprintf("%d,%d", row, col)] {
						// fmt.Println("found match")
						parts = append(parts, p)
						break thing
					}
				}
			}
		}
	}

	sum := 0

	for _, p := range parts {
		int, err := strconv.Atoi(p.num)
		if err != nil {
			log.Fatal(err)
		}
		sum += int
	}

	fmt.Printf("sum: %+v\n", sum)

}

// iterating through the lines to collect engine information
//	if isNum(l[i]) str join prev num
//	else write num and the num coords (current coord - 1 on the x plane)
//		thinking: this is struct{ part int, coords Coord} where coord is coord of last digit
//	if EOL and prev num != nil, terminate
//	if isSymbol(l[i]) write coord
//		thinking: this is a map[string]bool where default is false

// finding parts (iterating through the engine info struct list)
// 	find coords start point by going coord - len(int) - 1
//	for each coord from start point to struct[coord]
//		look around it in a star pattern to see if their is a symbol near by
//		* * *
//		*	c *
//		* * *
//		if so, add to sum, continue onto next partrroff
