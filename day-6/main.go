package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type point struct {
	x, y int
}

type field struct {
	value    int
	distance int
}

type area struct {
	infinite bool
	size     int
}

func inputToPoints(inp []string) []point {
	points := make([]point, 0, len(inp))
	re := regexp.MustCompile("[0-9]+")
	for _, v := range inp {
		inputArr := re.FindAllString(v, 2)
		x, _ := strconv.Atoi(inputArr[0])
		y, _ := strconv.Atoi(inputArr[1])
		points = append(points, point{x, y})
	}
	return points
}

func buildCanvas(points []point) [][]field {
	bX, bY := 0, 0
	for _, p := range points {
		if p.x > bX {
			bX = p.x
		}
		if p.y > bY {
			bY = p.y
		}
	}
	canvas := make([][]field, bX+2, bX+2)
	for i := range canvas {
		canvas[i] = make([]field, bY+2, bY+2)
	}
	return canvas
}

func printCanvas(canvas *[][]field) {
	for _, x := range *canvas {
		for _, y := range x {
			if y.distance == 0 {
				fmt.Printf(string(y.value + 65))
			} else {
				fmt.Printf(string(y.value + 97))
			}
		}
		fmt.Printf("\n")
	}
}

func isInfinite(x int, y int, canvas *[][]field, areaMap *map[int]*area) bool {
	pIdx := (*canvas)[x][y].value
	if (*areaMap)[pIdx].infinite == true {
		return true
	}
	xMax := len(*canvas)
	yMax := len((*canvas)[0])
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if x+i >= xMax || x+i < 0 || y+j >= yMax || y+j < 0 {
				return true
			}
		}
	}
	return false
}

func partOne(points []point) int {
	canvas := buildCanvas(points)
	for x, row := range canvas {
		for y := range row {
			var tmpField *field

			for pIdx, p := range points {
				dist := int(math.Abs(float64(x-p.x)) + math.Abs(float64(y-p.y)))
				if tmpField == nil || dist < tmpField.distance {
					tmpField = &field{pIdx, dist}
				} else if dist == tmpField.distance {
					tmpField = &field{-1, dist}
				} else {
					continue
				}
				canvas[x][y] = *tmpField
			}
		}
	}

	areaMap := make(map[int]*area)
	for x, row := range canvas {
		for y := range row {
			pIdx := canvas[x][y].value
			if areaMap[pIdx] == nil {
				areaMap[pIdx] = &area{false, 0}
			}
			if isInfinite(x, y, &canvas, &areaMap) {
				areaMap[pIdx].infinite = true
			} else {
				areaMap[pIdx].size++
			}
		}
	}
	biggestArea := 0
	for k := range areaMap {
		if !areaMap[k].infinite && areaMap[k].size > biggestArea && k > 0 {
			biggestArea = areaMap[k].size
		}
	}
	return biggestArea
}

func partTwo(points []point) int {
	canvas := buildCanvas(points)
	areaSize := 0
	for x, row := range canvas {
		for y := range row {
			totalDistance := 0
			for _, p := range points {
				if totalDistance >= 10000 {
					break
				}
				totalDistance += int(math.Abs(float64(x-p.x)) + math.Abs(float64(y-p.y)))
			}
			if totalDistance < 10000 {
				areaSize++
			}
		}
	}
	return areaSize
}

func main() {
	inp := make([]string, 0, 256)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		v := s.Text()
		inp = append(inp, v)
	}
	points := inputToPoints(inp)
	fmt.Println("1: ", partOne(points))
	fmt.Println("2: ", partTwo(points))
}
