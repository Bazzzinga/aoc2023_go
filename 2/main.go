package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const inputFileName = "input"

const Red = "red"
const Green = "green"
const Blue = "blue"

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	res1 := 0
	res2 := 0

	limits := map[string]int{Red: 12, Green: 13, Blue: 14}

	for scanner.Scan() {
		line := scanner.Text()

		res1 += getGameIdOrZero(line, limits)
		res2 += getCubePower(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("1: ", res1)
	fmt.Println("2: ", res2)
}

func getCubePower(game string) int {
	p1 := strings.Split(game, ": ")
	p2 := strings.Split(p1[1], "; ")

	minSet := CubeSet{}

	for i, round := range p2 {
		set := minimumRequiredSet(round)
		if i == 0 {
			minSet = set
			continue
		}

		if set.Red > minSet.Red {
			minSet.Red = set.Red
		}

		if set.Green > minSet.Green {
			minSet.Green = set.Green
		}

		if set.Blue > minSet.Blue {
			minSet.Blue = set.Blue
		}
	}

	return minSet.Red * minSet.Green * minSet.Blue
}

type CubeSet struct {
	Red   int
	Green int
	Blue  int
}

func minimumRequiredSet(round string) CubeSet {
	res := CubeSet{}

	cubes := strings.Split(round, ", ")

	for _, cube := range cubes {
		parts := strings.Split(cube, " ")

		n, _ := strconv.Atoi(parts[0])
		c := parts[1]

		switch c {
		case Red:
			if n > res.Red {
				res.Red = n
			}
		case Green:
			if n > res.Green {
				res.Green = n
			}
		case Blue:
			if n > res.Blue {
				res.Blue = n
			}
		}
	}

	return res
}

func getGameIdOrZero(game string, limits map[string]int) int {
	p1 := strings.Split(game, ": ")
	idS := p1[0][5:]
	id, _ := strconv.Atoi(idS)

	p2 := strings.Split(p1[1], "; ")

	for _, round := range p2 {
		if !isRoundPossible(round, limits) {
			return 0
		}
	}

	return id
}

func isRoundPossible(round string, limits map[string]int) bool {
	cubes := strings.Split(round, ", ")

	for _, cube := range cubes {
		parts := strings.Split(cube, " ")

		n, _ := strconv.Atoi(parts[0])
		c := parts[1]

		if n > limits[c] {
			return false
		}
	}

	return true
}
