package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
//		fmt.Println(line)
//		fmt.Println(string(rune(line[0])))
//		fmt.Println(string(rune(line[len(line)-1])))
//		sum += maxjolt(line)
		sum += maxjolt2(line)
	}
	fmt.Println(sum)
}

func maxjolt2(bank string) int {

//	fmt.Println(bank)

	num := ""
	remaining := 11
	
	for remaining >= 0 && len(bank) - remaining > 1 {
		max := rune(bank[0])
		imax := 0
	//	fmt.Println("len(bank) - remaining: ",len(bank) - remaining)
		for i, c := range bank[0:len(bank) - remaining] {
			if c > max {
				max = c
				imax = i
			}
		}
	//	fmt.Println("string(max): ", string(max))
		num += string(max)
	//	fmt.Println("imax: ", imax)
		bank = bank[imax+1:]
	//	fmt.Println("bank: ", bank)
		remaining--
	}
//	fmt.Println("num: ", num)
//	fmt.Println("bank: ", bank)
	if len(num) < 12 {
		num += bank
	}
	inum, _ := strconv.Atoi(num)
//	fmt.Println("inum: ", inum)
	return inum
}


func maxjolt(bank string) int {
	tens := bank[0:len(bank)-1]
	ones := bank
	//fmt.Println(bank)
	maxten := rune(tens[0])
	maxi := 0
	for i, c := range tens {
		if c > maxten {
			maxten = c
			maxi = i
		}
	}
	// fmt.Println(string(maxten)," ", maxi)
	// remove this from the ones
	// fmt.Println(ones)
	ones = ones[maxi+1:]
	// fmt.Println(ones)
	maxone := rune(ones[0])
	maxi = 0
	for i, c := range ones {
		if c > maxone {
			maxone = c
			maxi = i
		}
	}
	// fmt.Println(string(maxten), string(maxone))
	
	max := string(maxten)+string(maxone)
	// fmt.Println(max)
	maxjolt, _ :=  strconv.Atoi(max)
	return maxjolt
}