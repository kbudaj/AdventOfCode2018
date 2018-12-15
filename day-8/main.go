package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	qChildNode, qMeta int
	childNodes        []*node
	metaEntries       []int
}

func parse(data []string) (int, int, []string) {
	qChildren, _ := strconv.Atoi(data[0])
	qMeta, _ := strconv.Atoi(data[1])
	data = data[2:]
	totalMeta := 0
	childValues := make([]int, qChildren, qChildren)
	value := 0

	for i := 0; i < qChildren; i++ {
		deltaMeta := 0
		deltaMeta, value, data = parse(data)
		totalMeta += deltaMeta
		childValues[i] = value
	}

	totalDeltaMeta, deltaMeta := 0, 0
	for i := 0; i < qMeta; i++ {
		deltaMeta, _ = strconv.Atoi(data[i])
		totalDeltaMeta += deltaMeta
	}
	totalMeta += totalDeltaMeta

	if qChildren == 0 {
		return totalMeta, totalDeltaMeta, data[qMeta:]
	}
	value = 0
	for i := 0; i < qMeta; i++ {
		idx, _ := strconv.Atoi(data[i])
		if idx <= len(childValues) {
			value += childValues[idx-1]
		}
	}
	return totalMeta, value, data[qMeta:]
}

func main() {
	inp := make([]string, 0, 256)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := s.Text()
		inp = append(inp, v)
	}
	total, value, _ := parse(strings.Fields(inp[0]))
	fmt.Println("1: ", total)
	fmt.Println("2: ", value)
}
