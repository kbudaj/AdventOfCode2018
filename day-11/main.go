package main

import "fmt"

type square struct {
	x, y      int
	totalFuel int
}

func calcFuel(x, y, serialNr int) int {
	rackID := x + 10
	return ((((rackID*y + serialNr) * rackID) / 100) % 10) - 5
}

func partOne(serialNr int) string {
	topX, topY, topSquareFuel := 0, 0, 0
	size := 3
	for x := 1; x <= 300-size; x++ {
		for y := 1; y <= 300-size; y++ {
			fuel := 0
			for i := 0; i < size; i++ {
				for j := 0; j < size; j++ {
					fuel += calcFuel(x+i, y+j, serialNr)
				}
			}
			if fuel > topSquareFuel {
				topX, topY = x, y
				topSquareFuel = fuel
			}
		}
	}
	return fmt.Sprintf("%d,%d", topX, topY)
}

func partTwo(serialNr int) string {
	// https://www.codeproject.com/Articles/441226/Haar-feature-Object-Detection-in-Csharp#integral
	squareFuel := [301][301]int{}
	topX, topY, topSquareSize, topSquareFuel := 0, 0, 0, 0
	for s := 1; s <= 300; s++ {
		for x := 1; x <= 300-s; x++ {
			for y := 1; y <= 300-s; y++ {
				squareFuel[x][y] += calcFuel(x+s-1, y+s-1, serialNr)
				for i := 0; i < s-1; i++ {
					squareFuel[x][y] += calcFuel(x+i, y+s-1, serialNr)
					squareFuel[x][y] += calcFuel(x+s-1, y+i, serialNr)
				}
				if squareFuel[x][y] > topSquareFuel {
					topX, topY = x, y
					topSquareFuel = squareFuel[x][y]
					topSquareSize = s
				}
			}
		}
	}
	return fmt.Sprintf("%d,%d,%d", topX, topY, topSquareSize)
}

func main() {
	fmt.Println("1: ", partOne(7400))
	fmt.Println("2: ", partTwo(7400))
}
