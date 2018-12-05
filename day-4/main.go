package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partOne(inp []string) int {
	inpLength := len(inp)
	fmt.Println(inpLength)
	reductions := 0
	for i := 0; i+1 < inpLength; i++ {
		first, _ := strconv.Atoi(inp[i])
		second, _ := strconv.Atoi(inp[i+1])
		delta := first - second
		if delta == 32 || delta == -32 {
			inp = append(inp[:i], inp[i+2:]...)
			reductions++
		}
	}
	return inpLength - 2*reductions
}

func main() {
	inp := make([]string, 0, 256)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := s.Text()
		inp = append(inp, v)
	}
	fmt.Println("1: ", partOne(inp))
	// fmt.Println("2: ", part2(inp, board))
}
