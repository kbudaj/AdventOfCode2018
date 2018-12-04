package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x, y int
}

func lineToParams(line string) (int, int, int, int, int) {
	r := regexp.MustCompile(`#(\d+)\s@\s(\d+),(\d+)\:\s(\d+)x(\d+)`)
	res := r.FindStringSubmatch(line)
	claim, _ := strconv.Atoi(res[1])
	lOff, _ := strconv.Atoi(res[2])
	tOff, _ := strconv.Atoi(res[3])
	width, _ := strconv.Atoi(res[4])
	height, _ := strconv.Atoi(res[5])
	return claim, lOff, tOff, width, height
}

func createBoard(inp []string) map[point]int {
	board := make(map[point]int)
	for _, line := range inp {
		_, lOff, tOff, width, height := lineToParams(line)

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				p := point{x + lOff, y + tOff}
				board[p]++
			}
		}
	}
	return board
}

func part1(board map[point]int) int {
	overlaps := 0
	for _, value := range board {
		if value > 1 {
			overlaps++
		}
	}
	return overlaps
}

func part2(inp []string, board map[point]int) int {
	for _, line := range inp {
		claim, lOff, tOff, width, height := lineToParams(line)
		overlaps := false

		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				p := point{x + lOff, y + tOff}
				if board[p] > 1 {
					overlaps = true
					break
				}
			}
			if overlaps {
				break
			}
		}
		if !overlaps {
			return claim
		}
	}
	return 0
}

func main() {
	inp := make([]string, 0, 256)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := s.Text()
		inp = append(inp, v)
	}
	board := createBoard(inp)
	fmt.Println("1: ", part1(board))
	fmt.Println("2: ", part2(inp, board))
}
