package main

import (
	"container/list"
	"fmt"
)

func partOne(players int, mValue int) int {
	scores := make([]int, players, players)
	list := list.New()
	c := list.PushBack(0)

	for i := 1; i <= mValue; i++ {
		if i%23 == 0 {
			toDel := c
			for j := 0; j < 7; j++ {
				toDel = toDel.Prev()
				if toDel == nil {
					toDel = list.Back()
				}
			}
			scores[i%players] += i + toDel.Value.(int)
			c = toDel.Next()
			list.Remove(toDel)
		} else {
			n := c.Next()
			if n == nil {
				n = list.Front()
			}
			c = list.InsertAfter(i, n)
		}
	}

	topScore := scores[0]
	for _, s := range scores {
		if s > topScore {
			topScore = s
		}
	}
	return topScore
}

func main() {
	fmt.Println("1: ", partOne(458, 72019))
	fmt.Println("2: ", partOne(458, 7201900))
}
