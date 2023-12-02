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

/*
*

In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration.
However, game 3 would have been impossible because at one point the Elf showed you 20 red cubes at once;
similarly, game 4 would also have been impossible because the Elf showed you 15 blue cubes at once.
If you add up the IDs of the games that would have been possible, you get 8.

Determine which games would have been possible if the bag had been loaded with only
12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?

*
*/
func main() {
	file, err := fs.ReadFile(os.DirFS(filepath.Dir(".")), "input")
	if err != nil {
		panic(err)
	}

	m := make(map[string]int)
	m["red"] = 12
	m["green"] = 13
	m["blue"] = 14

	//Game 1: 6 green, 3 blue; 3 red, 1 green; 4 green, 3 red, 5 blu
	// Game 2: 2 red, 7 green; 13 green, 2 blue, 4 red; 4 green, 5 red, 1 blue; 1 blue, 9 red, 1 green
	// Game 3: 2 green, 3 blue, 9 red; 3 red, 2 green; 6 red, 4 blue; 6 red

	// split by : to get ["Game x", "6 green, 3 blue; 3 red, 1 green; 4 green, 3 red, 5 blue"]
	//		split on [0] by space to get ["Game", "x"]
	// 		split on [1] by ; to get ["6 green, 3 blue", "3 red, 1 green", "4 green, 3 red, 5 blue"]
	//				split on [m] by , to get [["6 green", "3 blue"], ["3 red", "1 green"], ["4 green", "3 red", "5 blue"]]
	//					split on [n] by space to get [[["6" , "green"], ["3", "blue"]], ...]

	possibleGames := 0

	for _, line := range strings.Split(string(file), "\n") {
		res := strings.Split(line, ":")
		if len(res) == 1 {
			continue
		}
		if len(res) != 2 {
			log.Fatalf("expected only one colon in line: %q", line)
		}

		gameId, err := strconv.Atoi(strings.Split(res[0], " ")[1])
		if err != nil {
			log.Fatal(err)
		}

		gamePossible := true

	gameLoop:
		for _, set := range strings.Split(res[1], ";") {
			for _, result := range strings.Split(set, ",") {
				t := strings.Split(strings.Trim(result, " "), " ") // 6 green
				amount, err := strconv.Atoi(t[0])
				if err != nil {
					log.Fatal(err)
				}

				colour := t[1]

				if amount > m[colour] {
					gamePossible = false
					break gameLoop
				}
				// todo: probably need to use more colour specific information for p 2
			}
		}

		if gamePossible {
			possibleGames += gameId
		}
	}

	fmt.Printf("res = %v\n", possibleGames)
}
