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

func main() {
	file, err := os.Open(inputFileName)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	schema := NewSchema()

	y := 0
	for scanner.Scan() {
		line := scanner.Text()

		schema.preprocessLine(line, y)
		y++
	}

	schema.MarkConnected()
	res1 := schema.GetConnectedNumbersSum()
	res2 := schema.GetConnectedGearsSum()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("1: ", res1)
	fmt.Println("2: ", res2)
}

type Coord struct {
	X int
	Y int
}

type CoordString string

func CoordFromString(s CoordString) Coord {
	cc := strings.Split(string(s), "_")
	x, _ := strconv.Atoi(cc[0])
	y, _ := strconv.Atoi(cc[1])

	return Coord{
		X: x,
		Y: y,
	}
}

func (c Coord) ToCoordString() CoordString {
	return CoordString(fmt.Sprintf("%d_%d", c.X, c.Y))
}

type Number struct {
	Id          int
	Coords      []Coord
	Value       int
	IsConnected bool
}

type Schema struct {
	nextId     int
	Numbers    []*Number
	NumbersMap map[CoordString]*Number
	SymbolsMap map[CoordString]rune
	GearsMap   map[CoordString][]*Number
}

func NewSchema() *Schema {
	return &Schema{
		NumbersMap: make(map[CoordString]*Number),
		SymbolsMap: make(map[CoordString]rune),
		Numbers:    make([]*Number, 0),
		GearsMap:   make(map[CoordString][]*Number),
	}
}

func (s *Schema) MarkConnected() {
	for c, sm := range s.SymbolsMap {
		coord := CoordFromString(c)
		for i := -1; i <= 1; i++ {
			for j := -1; j <= 1; j++ {
				s.markCoordsConnected(Coord{coord.X + i, coord.Y + j}, coord, string([]rune{sm}))
			}
		}
	}
}

func (s *Schema) markCoordsConnected(c, gc Coord, sm string) {
	n, ok := s.NumbersMap[c.ToCoordString()]
	if ok {
		n.IsConnected = true
		if sm == "*" {
			s.connectGear(gc, n)
		}
	}
}

func (s *Schema) connectGear(c Coord, n *Number) {
	cs := c.ToCoordString()
	gm, ok := s.GearsMap[cs]
	if !ok {
		s.GearsMap[cs] = make([]*Number, 0)
	}

	alreadyConnected := false
	for _, gn := range gm {
		if gn.Id == n.Id {
			alreadyConnected = true
			break
		}
	}

	if !alreadyConnected {
		s.GearsMap[cs] = append(s.GearsMap[cs], n)
	}
}

func (s *Schema) GetConnectedGearsSum() int {
	res := 0
	for _, ns := range s.GearsMap {
		if len(ns) == 2 {
			res += ns[0].Value * ns[1].Value
		}
	}
	return res
}

func (s *Schema) GetConnectedNumbersSum() int {
	res := 0

	for _, n := range s.Numbers {
		if n.IsConnected {
			res += n.Value
		}
	}

	return res
}

func (s *Schema) preprocessLine(line string, y int) {
	/*
		0: 48
		9: 57
		.: 46
	*/
	currentNumber := 0
	numberCoords := make([]Coord, 0)
	runes := []rune(line)
	for x, c := range runes {
		cc := Coord{X: x, Y: y}

		//is a digit
		if c >= 48 && c <= 57 {
			currentNumber = currentNumber*10 + int(c-48)
			numberCoords = append(numberCoords, cc)
			continue
		} else {
			if c != 46 {
				s.SymbolsMap[cc.ToCoordString()] = c
			}

			if currentNumber > 0 {
				nn := Number{
					Value:  currentNumber,
					Coords: numberCoords,
					Id:     s.nextId,
				}

				s.nextId++

				s.Numbers = append(s.Numbers, &nn)
				for _, cc := range nn.Coords {
					s.NumbersMap[cc.ToCoordString()] = &nn
				}

				currentNumber = 0
				numberCoords = make([]Coord, 0)
			}
		}
	}

	if currentNumber > 0 {
		nn := Number{
			Value:  currentNumber,
			Coords: numberCoords,
		}

		s.Numbers = append(s.Numbers, &nn)
		for _, cc := range nn.Coords {
			s.NumbersMap[cc.ToCoordString()] = &nn
		}

		currentNumber = 0
		numberCoords = make([]Coord, 0)
	}
}
