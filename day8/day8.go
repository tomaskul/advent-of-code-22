package day8

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tomaskul/advent-of-code-22/util"
)

const (
	Hidden = iota
	Visible
)

func Solution(sessionCookie, pt1Text, pt2Text string) {
	fmt.Printf(pt1Text)

	heightGrid := getHeightGrid(sessionCookie)

	grid, visibleFromOutside := getVisibilityGrid(heightGrid)
	for _, v := range grid {
		fmt.Printf("%v\n", v)
	}
	fmt.Printf("\n\n")
	//fmt.Printf("%v\n\n", grid)
	fmt.Printf("Trees visible from outside: %d\n", visibleFromOutside)

	//fmt.Printf(pt2Text)
}

func getHeightGrid(sessionCookie string) [][]int {
	rows := util.GetRows("https://adventofcode.com/2022/day/8/input", sessionCookie)
	result := make([][]int, len(rows))
	for i, row := range rows {
		result[i] = getIntRow(row)
	}
	return result
}

func getIntRow(row string) []int {
	items := strings.Split(row, "")
	result := make([]int, len(items))
	for i, v := range items {
		value, _ := strconv.Atoi(v)
		result[i] = value
	}
	return result
}

func getVisibilityGrid(heightGrid [][]int) ([][]int, int) {
	visibilityGrid := make([][]int, len(heightGrid))
	for i := 0; i < len(visibilityGrid); i++ {
		visibilityGrid[i] = getVisibleFromLeftAndRight(heightGrid[i])
	}

	visibilityGrid = getVisibleFromTopAndBottom(heightGrid, visibilityGrid)

	return visibilityGrid, countVisible(visibilityGrid)
}

// Reads heights left-to-right and right-to-left and identifies whether trees
// are visible. Outermost trees are visible by default.
func getVisibleFromLeftAndRight(row []int) []int {
	result := make([]int, len(row))
	result[0], result[len(row)-1] = Visible, Visible

	// Left to right.
	tallestSoFar := row[0]
	for i := 0; i < len(row); i++ {
		if i == 0 {
			continue
		}
		if row[i] > row[i-1] && row[i] > tallestSoFar {
			result[i] = Visible
			tallestSoFar = row[i]
		}
	}

	// Right to left (last tree visible from the right by default).
	tallestSoFar = row[len(row)-1]
	for i := len(row) - 1; i > 0; i-- {
		if i == len(row)-1 {
			continue
		}
		if row[i] > row[i+1] && row[i] > tallestSoFar {
			result[i] = Visible
			tallestSoFar = row[i]
		}
	}
	return result
}

func getVisibleFromTopAndBottom(heightGrid, visibilityGrid [][]int) [][]int {
	maxY := len(heightGrid)
	maxX := len(heightGrid[0])

	for x := 0; x < maxX; x++ {
		// Top to bottom (0th trees picked up by left-to-right scan).
		tallestSoFar := heightGrid[0][x]
		for y := 0; y < maxY; y++ {
			if y == 0 {
				visibilityGrid[y][x] = Visible
				continue
			}

			if heightGrid[y-1][x] < heightGrid[y][x] && heightGrid[y][x] > tallestSoFar {
				visibilityGrid[y][x] = Visible
				tallestSoFar = heightGrid[y][x]
			} else {
				break
			}
		}

		// Bottom to top (0th trees picked up by left-to-right scan).
		tallestSoFar = heightGrid[maxY-1][x]
		for y := maxY - 1; y > 0; y-- {
			if y == maxY-1 {
				visibilityGrid[y][x] = Visible
				continue
			}

			if heightGrid[y][x] > heightGrid[y+1][x] && heightGrid[y][x] > tallestSoFar {
				visibilityGrid[y][x] = Visible
				tallestSoFar = heightGrid[y][x]
			} else {
				break
			}
		}
	}

	return visibilityGrid
}

func countVisible(visibleGrid [][]int) int {
	result := 0
	for i := range visibleGrid {
		for j := range visibleGrid[i] {
			if visibleGrid[i][j] == Visible {
				result++
			}
		}
	}

	return result
}
