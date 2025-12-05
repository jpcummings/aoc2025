package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)


func main() {
	var dial int = 50
	var rot int
	var zeros, crossings int = 0, 0
	scanner := bufio.NewScanner(os.Stdin)

	for ( scanner.Scan() ) {
		line := scanner.Text()
		fmt.Println(line)
		rot,_ = strconv.Atoi(line[1:])
		//fmt.Println(i)
		fmt.Println("dial: ", dial)
		if line[0] == 'R' {
			//fmt.Println("right")
			crossings += (dial+rot)/100
			dial = mod( (dial+rot), 100)
			if dial == 0 {
				crossings -= 1
				zeros += 1
			}
		} else if line[0] == 'L' {
			//fmt.Println("left")
			fmt.Println(mod(-1*dial, 100) + rot)
			crossings += (mod(-1*dial, 100) + rot)/100
			dial = mod( (dial-rot), 100)
			if dial == 0 {
				crossings -= 1
				zeros += 1
			}
		} else {
			fmt.Println("huh?")
		}
		//fmt.Println(line)
		fmt.Println("crossings: ", crossings)
		fmt.Println("zeros: ", zeros)
		fmt.Println("dial: ", dial)
		fmt.Println()
	}
	fmt.Println(zeros)
	fmt.Println(crossings)
	fmt.Println(zeros+crossings)
}

func mod(a, b int) int {
	return (a % b + b) % b
}