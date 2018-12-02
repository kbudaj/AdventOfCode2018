package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func part1(inp []int) int {
	offset := 0
	for _, v := range inp {
		offset += v
	}
	return offset
}

func part2(inp []int) int {
	offset := 0
	values := make([]int, 0, 256)
	for {
		for _, v := range inp {
			offset += v

			idx := sort.SearchInts(values, offset)
			if (idx < len(values)) && (values[idx] == offset) {
				return offset
			}

			values = append(values, 0)
			copy(values[idx+1:], values[idx:])
			values[idx] = offset
		}
	}
}

func main() {
	inp := make([]int, 0, 256)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v, _ := strconv.Atoi(s.Text())
		inp = append(inp, v)
	}
	fmt.Println("1: ", part1(inp))
	fmt.Println("2: ", part2(inp))
}
