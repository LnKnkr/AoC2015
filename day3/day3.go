package main

import (
	"fmt"
	"os"
)

type santa struct {
	posX int
	posY int
}

func createSanta(pos int) *santa {
	return &santa{posX: pos, posY: pos}
}

func (s *santa) move(b byte) {
	if b == '^' {
		s.posX++
	}
	if b == '<' {
		s.posY++
	}
	if b == 'v' {
		s.posX--
	}
	if b == '>' {
		s.posY--
	}
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	oneSanta(file)
	twoSanta(file)
}

func oneSanta(file []byte) {
	area := make([][]int, len(file))
	for i := range area {
		area[i] = make([]int, len(file))
	}

	santa := createSanta(len(file) / 2)
	for _, b := range file {
		santa.move(b)
		area[santa.posX][santa.posY]++
	}

	amount := 0
	for _, l := range area {
		for _, r := range l {
			if r > 0 {
				amount++
			}
		}
	}
	fmt.Printf("%d", amount)
}

func twoSanta(file []byte) {
	area := make([][]int, len(file))
	for i := range area {
		area[i] = make([]int, len(file))
	}

	santa := createSanta(len(file) / 2)
	robosanta := createSanta(len(file) / 2)
	for i, b := range file {
		if i%2 == 1 {
			santa.move(b)
			area[santa.posX][santa.posY]++
		} else {
			robosanta.move(b)
			area[robosanta.posX][robosanta.posY]++
		}
	}

	amount := 0
	for _, l := range area {
		for _, r := range l {
			if r > 0 {
				amount++
			}
		}
	}
	fmt.Printf("%d", amount)
}
