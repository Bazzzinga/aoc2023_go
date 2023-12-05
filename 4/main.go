package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const inputFileName = "input"

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	res1 := 0

	scanner := bufio.NewScanner(file)

	i := 0
	costs := make(map[int]int)
	for scanner.Scan() {
		line := scanner.Text()

		cost := processCard(line)

		costs[i] = cost
		if cost > 0 {
			res1 += 1 << (cost - 1)
		}
		i++
	}

	cardCounts := make([]int, len(costs))
	for i := range costs {
		cardCounts[i] = 1
	}

	for i, n := range cardCounts {
		cost := costs[i]
		for j := 1; j <= cost; j++ {
			cardCounts[i+j] += n
		}
	}

	res2 := 0

	for _, n := range cardCounts {
		res2 += n
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("1: ", res1)
	fmt.Println("2: ", res2)
}

func processCard(line string) int {
	p1 := strings.Split(line, ": ")
	p2 := strings.Split(p1[1], " | ")

	need := p2[0]
	have := p2[1]

	need = strings.Replace(strings.TrimSpace(need), "  ", " ", -1)
	have = strings.Replace(strings.TrimSpace(have), "  ", " ", -1)

	needS := strings.Split(need, " ")
	haveS := strings.Split(have, " ")

	i := intersect(needS, haveS)
	fmt.Println(i)
	return len(i)
}

func intersect(slice1, slice2 []string) []string {
	var res []string
	for _, element1 := range slice1 {
		for _, element2 := range slice2 {
			if element1 == element2 {
				res = append(res, element1)
				break
			}
		}
	}
	return res //return slice after intersection
}
