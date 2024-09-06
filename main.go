package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	width  = 200
	height = 40
)

type Universe [][]bool

func NewUniverse() Universe {
	sls := make([][]bool, height)
	for i := range sls {
		sls[i] = make([]bool, width)
	}
	return Universe(sls)
}

func (u Universe) Seed() {
	for i := 0; i < (width * height / 4); i++ {
		u.Set(rand.Intn(height), rand.Intn(width), true)
	}
}

func (u Universe) String() string {
	str := ""
	str += fmt.Sprintf("┏")
	for i := 0; i < width; i++ {
		str += fmt.Sprintf("━")
	}
	str += fmt.Sprintf("┓")
	str += fmt.Sprintf("\n")
	for i := range u {
		str += fmt.Sprintf("┃")
		for j := range u[i] {
			if u.Alive(i, j) {
				//str += fmt.Sprintf("⏹")
				str += fmt.Sprintf("๏")
			} else {
				str += fmt.Sprintf(" ")
			}
		}
		str += fmt.Sprintf("┃")
		str += fmt.Sprintf("\n")
	}
	str += fmt.Sprintf("┗")
	for i := 0; i < width; i++ {
		str += fmt.Sprintf("━")
	}
	str += fmt.Sprintf("┛")
	return str
}

func (u Universe) Show() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	_ = err
	fmt.Println(u.String())
}

func (u Universe) Neighbors(x, y int) int {
	neighbors := 0
	for h := -1; h <= 1; h++ {
		for v := -1; v <= 1; v++ {
			if !(h == 0 && v == 0) && u.Alive(x+h, y+v) {
				neighbors++
			}
		}
	}
	return neighbors
}

func (u Universe) Alive(x, y int) bool {
	X := (x + height) % height
	Y := (y + width) % width
	return u[X][Y]
}

func (u Universe) Set(x, y int, value bool) {
	u[x][y] = value
}

func (u Universe) Next(x, y int) bool {
	neighbors := u.Neighbors(x, y)
	return neighbors == 3 || neighbors == 2 && u.Alive(x, y)
}

func Step(a, b Universe) {
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			a.Set(x, y, b.Next(x, y))
		}
	}
}

func main() {

	universe, universeCopy := NewUniverse(), NewUniverse()
	universe.Seed()

	for {
		universe.Show()
		Step(universeCopy, universe)
		universe, universeCopy = universeCopy, universe
		time.Sleep(time.Second / 30)
	}

}
