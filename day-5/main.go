package main

import (
	"bufio"
	"fmt"
	"os"
)

func partOne(inp []rune) int {
	inpLength := len(inp)
	for i := 0; i+1 < inpLength; i++ {
		if i < 0 {
			i = 0
		}
		first := int(inp[i])
		second := int(inp[i+1])
		delta := first - second
		if delta == 32 || delta == -32 {
			inp = append(inp[:i], inp[i+2:]...)
			inpLength -= 2
			i -= 2
		}
	}
	return len(inp)
}

func partTwo(inp []rune) int {
	inpLen := len(inp)
	var bestResult int
	for i := 97; i <= 122; i++ {
		reduced := make([]rune, 0, inpLen)
		for _, v := range inp {
			if int(v) != i && int(v) != i-32 {
				reduced = append(reduced, v)
			}
		}
		result := partOne(reduced)
		if bestResult == 0 {
			bestResult = result
		} else {
			if result < bestResult {
				bestResult = result
			}
		}
	}
	return bestResult
}

func main() {
	inp := make([]string, 0, 256)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := s.Text()
		inp = append(inp, v)
	}
	inpChars := []rune(inp[0])
	fmt.Println("1: ", partOne(inpChars))
	fmt.Println("2: ", partTwo(inpChars))
}
