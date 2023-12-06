package main

import (
	"testing"

	"github.com/matryer/is"
)

func TestDay2Part2(t *testing.T) {
	is := is.New(t)

	for i, game := range first15 {
		power := Part2GameResults(game, i)
		is.Equal(power, expectation[i])
	}

}

var expectation []int = []int{3, 2, 4, 24, 120, 16, 36, 48, 12, 2, 2, 1, 90, 84, 3}

var first15 []string = []string{
	"Game 1: 1 blue, 8 green; 14 green, 15 blue; 3 green, 9 blue; 8 green, 8 blue, 1 red; 1 red, 9 green, 10 blue",
	"Game 2: 3 blue, 1 green, 2 red; 2 red, 2 green, 5 blue; 3 green, 10 blue; 8 red, 1 blue; 3 red, 1 green, 5 blue; 1 blue, 5 red, 3 green",
	"Game 3: 4 green, 1 blue; 6 blue, 5 green, 1 red; 11 green, 10 blue",
	"Game 4: 12 blue, 12 green, 3 red; 15 blue, 1 green, 10 red; 8 blue, 3 red, 2 green; 14 red, 8 blue",
	"Game 5: 7 blue, 8 red, 5 green; 15 blue, 16 red, 14 green; 3 blue, 14 red, 10 green",
	"Game 6: 4 blue, 13 red; 1 green, 13 blue, 11 red; 4 red, 19 blue; 18 blue, 10 red, 1 green",
	"Game 7: 8 green, 3 blue, 3 red; 2 red, 7 green, 10 blue; 6 green, 11 red, 3 blue",
	"Game 8: 10 red, 6 green, 1 blue; 15 green, 10 red, 3 blue; 8 red, 10 green, 5 blue",
	"Game 9: 2 green, 8 blue, 1 red; 6 blue, 10 red; 13 blue, 12 red, 7 green",
	"Game 10: 2 blue, 8 red, 10 green; 1 green, 2 blue; 1 red, 1 green; 7 red, 2 blue, 1 green",
	"Game 11: 8 green, 1 blue; 6 green; 2 green, 1 blue; 2 blue, 11 green; 1 red, 12 green",
	"Game 12: 3 red, 2 green, 15 blue; 1 blue, 1 green, 4 red; 1 green, 12 blue, 3 red; 1 red, 10 blue; 3 red, 2 green, 14 blue; 3 red, 13 blue",
	"Game 13: 7 blue, 5 red; 7 red, 3 green, 9 blue; 9 green, 7 blue, 7 red; 6 blue, 8 red; 11 red; 3 green, 7 blue, 8 red",
	"Game 14: 4 blue, 6 green, 7 red; 8 red, 4 green, 11 blue; 3 green, 9 red, 13 blue",
	"Game 15: 3 green, 1 blue, 5 red; 2 red; 1 red, 4 green",
}
