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


The missing part wasn't the only issue - one of the gears in the engine is wrong. A gear is any * symbol that is adjacent to exactly two part numbers. Its gear ratio is the result of multiplying those two numbers together.

This time, you need to find the gear ratio of every gear and add them all up so that the engineer can figure out which gear needs to be replaced.

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

In this schematic, there are two gears. The first is in the top left; it has part numbers 467 and 35, so its gear ratio is 16345. The second gear is in the lower right; its gear ratio is 451490. (The * adjacent to 617 is not a gear because it is only adjacent to one part number.) Adding up all of the gear ratios produces 467835.

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

	gears := make(map[string]bool)
	gearNeighbours := make(map[string][]part)
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

				if cell == "*" {
					gears[fmt.Sprintf("%d,%d", rowIdx, colIdx)] = true
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
			// now iterate through cols in each row, looking for match
			for row := p.row - 1; row <= p.row+1; row += 1 {
				for col := (p.col - 1) - i; col <= p.col+1; col += 1 {
					// fmt.Println("looking at row x col", row, col)
					if gears[fmt.Sprintf("%d,%d", row, col)] {
						gearNeighbours[fmt.Sprintf("%d,%d", row, col)] = append(gearNeighbours[fmt.Sprintf("%d,%d", row, col)], p)
						// fmt.Println("found match")
						parts = append(parts, p)
						break thing
					}
				}
			}
		}
	}

	sum := 0

	for _, value := range gearNeighbours {
		if len(value) == 2 {
			val := 1
			for _, part := range value {
				int, err := strconv.Atoi(part.num)
				if err != nil {
					log.Fatal(err)
				}
				val *= int
			}

			sum += val
		}
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
