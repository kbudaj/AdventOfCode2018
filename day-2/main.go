package main

import (
	"bufio"
	"fmt"
	"os"
)

func checksum(inp string) (bool, bool) {
	t := make([]int, 26, 26) // a = 97, z = 122
	for _, char := range inp {
		t[char-97]++
	}
	var doubles, triples bool
	for _, v := range t {
		if v == 2 {
			doubles = true
		}
		if v == 3 {
			triples = true
		}
	}
	return doubles, triples
}

func part1(inp []string) int {
	doubles, triples := 0, 0

	for _, i := range inp {
		d, t := checksum(i)
		if d {
			doubles++
		}
		if t {
			triples++
		}
	}
	return doubles * triples
}

func part2(inp []string) string {
	length := len(inp[0])
	for lineNr, line := range inp {
		for _, comp := range inp[lineNr+1:] {
			differ, differIdx := 0, 0
			for i := 0; i < length; i++ {
				if differ >= 2 {
					break
				}
				if line[i] != comp[i] {
					differ++
					differIdx = i
				}
			}
			if differ == 1 {
				return line[:differIdx] + line[differIdx+1:]
			}
		}
	}
	return ""
}

func main() {
	inp := make([]string, 0, 256)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := s.Text()
		inp = append(inp, v)
	}
	fmt.Println("1: ", part1(inp))
	fmt.Println("2: ", part2(inp))
}
