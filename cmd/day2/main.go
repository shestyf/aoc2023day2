package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var cubes map[string]int = map[string]int{
	"red":   12,
	"blue":  14,
	"green": 13,
}

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		panic(fmt.Sprintf("couldn't open file: %s", err))
	}
	defer input.Close()

	fileContent := bufio.NewScanner(input)
	fileContent.Split(bufio.ScanLines)

	fileLines := []string{}

	for fileContent.Scan() {
		fileLines = append(fileLines, fileContent.Text())
	}

	fmt.Printf("successfully read %d, lines from file\n", len(fileLines))

	totalPossibleGames := 0
	totalPower := 0

	for i, game := range fileLines {
		possible := Part1GameResults(game, i+1)
		if possible {
			totalPossibleGames += i + 1
		}
	}

	fmt.Printf("Part 1:\n total after adding index of all possible games is %d\n", totalPossibleGames)

	for i, game := range fileLines {
		power := Part2GameResults(game, i+1)
		totalPower += power
	}

	fmt.Printf("Part 2:\n total after adding the power of all games is %d\n", totalPower)
}

func Part2GameResults(game string, index int) int {

	prefix := fmt.Sprintf("Game %d: ", index)
	trimmed := strings.TrimPrefix(game, prefix)
	games := strings.Split(trimmed, ";")

	power := 0

	blue := 0
	red := 0
	green := 0

	for i, g := range games {
		i++
		colours := strings.Split(g, " ")
		pairs := make(map[string]int)

		for j := 0; j < len(colours); j++ {
			number, err := strconv.Atoi(colours[j])
			if err != nil {
				continue
			}

			trimmed := strings.TrimSuffix(colours[j+1], ",")
			pairs[trimmed] = number
		}

		if green < pairs["green"] && pairs["green"] != 0 {
			green = pairs["green"]
		}

		if blue < pairs["blue"] && pairs["blue"] != 0 {
			blue = pairs["blue"]
		}

		if red < pairs["red"] && pairs["red"] != 0 {
			red = pairs["red"]
		}

	}

	power = red * green * blue

	return power
}

func Part1GameResults(game string, index int) bool {

	prefix := fmt.Sprintf("Game %d: ", index)
	trimmed := strings.TrimPrefix(game, prefix)
	games := strings.Split(trimmed, ";")

	for i, g := range games {
		i++
		colours := strings.Split(g, " ")
		pairs := make(map[string]int)

		for j := 0; j < len(colours); j++ {
			number, err := strconv.Atoi(colours[j])
			if err != nil {
				continue
			}

			trimmed := strings.TrimSuffix(colours[j+1], ",")
			pairs[trimmed] = number
		}

		if pairs["blue"] > cubes["blue"] || pairs["red"] > cubes["red"] || pairs["green"] > cubes["green"] {
			return false
		}
	}

	return true
}
