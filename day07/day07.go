package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {

	manifold := ReadManifold()

//	PrintManifold(manifold)
//	manifold, nsplit := RayTrace(manifold)
//	PrintManifold(manifold)
//	fmt.Println(nsplit)

//	PrintManifold(manifold)
	fmt.Println(QRayTrace(manifold))
}

func ReadManifold() [][]rune {

	manifold := [][]rune {}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		row := []rune{}
		line := scanner.Text()
		for _, c := range line {
			row = append(row, c)
		}
		manifold = append(manifold, row)
	}
	return manifold
}

func PrintManifold(m [][]rune) {
	for _, row := range m {
		for _, c := range row {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
}

func RayTrace(manifold [][]rune) ([][]rune, int) {

	// find start
	for icol, c := range manifold[0] {
		if c == 'S' {
			manifold[0][icol] = '|'
		}
	}

	nsplit := 0
	for irow, row := range manifold {
		for icol, _ := range row {
			 if manifold[irow][icol] == '|' {
				if irow < len(manifold)-1 && manifold[irow+1][icol] == '.' {
					manifold[irow+1][icol]  = '|'
				} else if irow < len(manifold)-1 && manifold[irow+1][icol] == '^' {
					nsplit += 1
					manifold[irow+1][icol-1]  = '|'
					manifold[irow+1][icol+1]  = '|'
				}
			}
		}
	}
	return manifold, nsplit
}


func QRayTrace(manifold [][]rune) int {

	npaths := make([][]int, len(manifold))
	for i := range npaths {
		npaths[i] = make([]int, len(manifold[0]))
	}

	// find start
	for icol, c := range manifold[0] {
		if c == 'S' {
			manifold[0][icol] = '|'
			npaths[0][icol] = 1
		}
	}

	nsplit := 0
	for irow, row := range manifold {
		for icol, _ := range row {
			 if manifold[irow][icol] == '|' {
		//		fmt.Println("(",irow,", ",icol,")")
				if irow < len(manifold)-1 && manifold[irow+1][icol] == '^' {
					nsplit += 1
					manifold[irow+1][icol-1]  = '|'
					npaths[irow+1][icol-1]  += npaths[irow][icol]
					manifold[irow+1][icol+1]  = '|'
					npaths[irow+1][icol+1]  += npaths[irow][icol]
		//			fmt.Println("splitting ", npaths[irow][icol])
				}
				if irow < len(manifold)-1 && (manifold[irow+1][icol] == '.' || manifold[irow+1][icol] == '|') {
					manifold[irow+1][icol]  = '|'
					npaths[irow+1][icol]  += npaths[irow][icol]
		//			fmt.Println("carrying ", npaths[irow][icol], " down")
				}
			}
		}
	}
//	PrintPaths(npaths)
	bottomrow := npaths[len(npaths)-1]
	sum := 0
	for _, n := range bottomrow {
		sum += n
	}
	return sum
}

func PrintPaths(paths [][]int) {

	for irow, row := range paths {
		for icol, _ := range row {
			fmt.Print(paths[irow][icol], " ")
		}
		fmt.Println()
	}
	
}