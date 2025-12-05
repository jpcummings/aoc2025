package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"math"
)

func main() {
	var ranges []string
	var sum int = 0

	scanner := bufio.NewScanner(os.Stdin)

	for (scanner.Scan() ) {
		line:= scanner.Text()
		ranges = strings.Split(line,",")
	}
	for _, r := range ranges {
//		fmt.Println(r)
		limit := strings.Split(r,"-")
		low, _ := strconv.Atoi(limit[0])
		high, _ := strconv.Atoi(limit[1])
		for n := low; n <= high; n++ {
			if !valid(n) {
				fmt.Println(n)
				sum += n
			}
		}
	}
	fmt.Println(sum)
}

func valid(id int) bool {
	valid := true
	ndigits := int(math.Floor(math.Log10(float64(id))))+1
	mid := int(math.Ceil(float64(ndigits/2)))
//	fmt.Println(id)
//	fmt.Println(ndigits)
	for i := 1; i <= mid; i++ {
		div := int(math.Pow(10.0, float64(i)))
		subs :=make([]int,0,10)
		num := id
		for num > 0 {
			s := num % div
//			subs = append(subs, s)
			subs = append([]int{s}, subs...)
			num = num/ div
		}
//		fmt.Println(subs)
		// are all subs the same?
		match := true
		for _, s := range subs {
//			fmt.Println("checking ",s, " against ", subs[0])
			if s != subs[0] {
				match = false
			}
		}
		if match == true {
			fmt.Println("found invalid id ", id)
			fmt.Println(subs)
			valid = false
			break
		}			
	}
	return valid
}

func part1_valid(id int) bool {
	valid := true
	ndigits := int(math.Floor(math.Log10(float64(id))))+1
//	fmt.Println(ndigits)
	if ndigits%2 == 0 {
		left := id/int(math.Pow(10.0,float64(ndigits/2)))
		right := id - left*int(math.Pow(10.0,float64(ndigits/2)))
//		fmt.Println(left, " ", right)
		if left == right {
			valid = false
		}
	}
	return valid
}
