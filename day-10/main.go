package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type coord struct {
	x, y int
}

type point struct {
	x, y       int
	velX, velY int
}

func lineToPoint(l string) point {
	re := regexp.MustCompile(`position=<\s?(\-?\d+),\s+(\-?\d+)>.+<\s?(\-?\d+),\s+(\-?\d+)>`)
	inpArr := re.FindStringSubmatch(l)
	x, _ := strconv.Atoi(inpArr[1])
	y, _ := strconv.Atoi(inpArr[2])
	velX, _ := strconv.Atoi(inpArr[3])
	velY, _ := strconv.Atoi(inpArr[4])
	return point{x, y, velX, velY}
}

func drawBoard(points []point, pointMap map[coord]*point, s int) {
	minX, maxX, minY, maxY := points[0].x, points[0].x, points[0].y, points[0].y
	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	draw := false
	maxCounter := 0
	for _, p := range points {
		counter := 0
		for i := 1; i < 10; i++ {
			c := coord{p.x, p.y + i}
			if pointMap[c] != nil {
				counter++
			} else {
				break
			}
		}
		if counter > maxCounter {
			maxCounter = counter
			if maxCounter > 4 {
				draw = true
				break
			}
		}
	}

	if draw == false {
		return
	}
	for y := minY; y <= maxY; y++ {
		for x := minX; x < maxX; x++ {
			c := coord{x, y}
			if pointMap[c] == nil {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Println(s)
}

func partOne(inp []string) {
	points := make([]point, 0, len(inp))
	pointMap := make(map[coord]*point)
	for _, l := range inp {
		p := lineToPoint(l)
		points = append(points, p)
		c := coord{p.x, p.y}
		pointMap[c] = &p
	}
	for s := 1; s < 10350; s++ {
		for i := 0; i < len(points); i++ {
			p := points[i]
			c := coord{p.x, p.y}
			delete(pointMap, c)
			p.x += p.velX
			p.y += p.velY
			points[i] = p
			c = coord{p.x, p.y}
			pointMap[c] = &p

		}
		drawBoard(points, pointMap, s)
	}
}

func main() {
	inp := make([]string, 0, 256)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := s.Text()
		inp = append(inp, v)
	}
	partOne(inp)
}
