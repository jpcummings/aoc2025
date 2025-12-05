package main

import (
	"fmt"
	"bufio"
	"os"
)

type Point struct {
	i, j int
}

func main() {

	floor := ReadMap()

	// Part 1
//	PrintMap(floor)
	na := CountAccessible(floor)
	fmt.Println(na)

	// Part 2
	removed := 0
	for pts := ListAccessible(floor); len(pts) > 0; {
	//	fmt.Println(pts)
		floor = RemoveRolls(floor, pts)
		removed += len(pts)
//		PrintMap(floor)
		pts = ListAccessible(floor)
	}
	fmt.Println(removed)

}

func RemoveRolls(floor [][]string, points []Point) [][]string {

	for _, pt := range points {
		floor[pt.i][pt.j] = "."
	}

	return floor
}

func ListAccessible(floor [][]string) []Point {

	naccessible := 0

	var points []Point

	for i, row := range floor {
		for j, _ := range row {
//			fmt.Print(floor[i][j])
//			fmt.Println("cell = ",i,",",j)
			nrolls := 0
			for deli := -1; deli < 2; deli += 1 {
				for delj := -1; delj <2; delj += 1 {
					if deli == 0 && delj == 0 {
						continue
					}
					if (i+deli < 0) || (i+deli > len(floor[j])-1 ) {
						continue
					}
					if (j+delj < 0) || (j+delj > len(floor)-1 ) {
						continue
					}
//					fmt.Print(i+deli, ",", j+delj, " ")
					// check for a roll
					if floor[i+deli][j+delj]=="@" {
						nrolls+=1
					}
				}
			}
//			fmt.Println("nrolls = ", nrolls)
			if (floor[i][j] == "@" && nrolls < 4) {
				points = append(points, Point{i, j})
				naccessible += 1
			}
		}
//		fmt.Println()
	}


	return points
}


func CountAccessible(floor [][]string) int {

	naccessible := 0

	for i, row := range floor {
		for j, _ := range row {
//			fmt.Print(floor[i][j])
//			fmt.Println("cell = ",i,",",j)
			nrolls := 0
			for deli := -1; deli < 2; deli += 1 {
				for delj := -1; delj <2; delj += 1 {
					if deli == 0 && delj == 0 {
						continue
					}
					if (i+deli < 0) || (i+deli > len(floor[j])-1 ) {
						continue
					}
					if (j+delj < 0) || (j+delj > len(floor)-1 ) {
						continue
					}
//					fmt.Print(i+deli, ",", j+delj, " ")
					// check for a roll
					if floor[i+deli][j+delj]=="@" {
						nrolls+=1
					}
				}
			}
//			fmt.Println("nrolls = ", nrolls)
			if (floor[i][j] == "@" && nrolls < 4) {
				naccessible += 1
			}
		}
//		fmt.Println()
	}


	return naccessible
}


func ReadMap() [][]string {
	scanner := bufio.NewScanner(os.Stdin)

	var floor [][]string

	for scanner.Scan() {
		line := scanner.Text()
//		fmt.Println(line)
		var row []string
		for _, c := range line {
			row = append(row,string(c))
		}
		floor = append(floor,row)
	}
	return floor
}

func PrintMap(floor [][]string) {
	for i, row := range floor {
		for j, _ := range row {
			fmt.Print(floor[i][j])
		}
		fmt.Println()
	}

}