package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const inputFileName = "input"

func getNumbers() []string {
	return []string{
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}
}

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	res1 := 0
	res2 := 0

	for scanner.Scan() {
		line := scanner.Text()

		res1 += getLineNumber1(line)
		res2 += getLineNumber2(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("1: ", res1)
	fmt.Println("2: ", res2)
}

func getLineNumber2(line string) int {
	return 10*getLineNumber2Left(line) + getLineNumber2Right(line)
}

func getLineNumber2Left(line string) int {
	firstDigitIdx := -1
	firstDigit := 0
	for i := 0; i < len(line); i++ {
		c := line[i]
		if c >= 0x30 && c <= 0x39 {
			firstDigitIdx = i
			firstDigit = int(c - 0x30)
			break
		}
	}

	for i, number := range getNumbers() {
		firstDigitIdx, firstDigit = checkLeft(line, number, firstDigitIdx, firstDigit, i+1)
	}

	if firstDigit > 0 {
		return firstDigit
	}

	return 0
}

func checkLeft(s, numS string, currIdx, currNum, num int) (int, int) {
	tmp := strings.Index(s, numS)
	if tmp >= 0 {
		if (currIdx >= 0 && tmp < currIdx) || currIdx < 0 {
			currIdx = tmp
			currNum = num
		}
	}
	return currIdx, currNum
}

func getLineNumber2Right(line string) int {
	firstDigitIdx := -1
	firstDigit := 0
	for i := len(line) - 1; i >= 0; i-- {
		c := line[i]
		if c >= 0x30 && c <= 0x39 {
			firstDigitIdx = i
			firstDigit = int(c - 0x30)
			break
		}
	}

	for i, number := range getNumbers() {
		firstDigitIdx, firstDigit = checkRight(line, number, firstDigitIdx, firstDigit, i+1)
	}

	if firstDigit > 0 {
		return firstDigit
	}

	return 0
}

func checkRight(s, numS string, currIdx, currNum, num int) (int, int) {
	tmp := strings.LastIndex(s, numS)
	if tmp >= 0 {
		if (currIdx >= 0 && tmp > currIdx) || currIdx < 0 {
			currIdx = tmp
			currNum = num
		}
	}
	return currIdx, currNum
}

func getLineNumber1(line string) int {
	re := regexp.MustCompile(`[a-zA-Z]+`)
	s := re.ReplaceAllString(line, "")
	cs := strings.Split(s, "")

	if len(cs) > 0 {
		left := cs[0]
		right := cs[len(cs)-1]

		ln, err := strconv.Atoi(left)
		if err != nil {
			log.Fatal(err)
		}
		rn, err := strconv.Atoi(right)
		if err != nil {
			log.Fatal(err)
		}
		return 10*ln + rn
	}
	return 0
}
