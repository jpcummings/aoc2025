package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"slices"
)

type Range struct {
	start, end int
}

func main() {

	rawranges, ids := ReadDB()
//	fmt.Println(rawranges)
//	fmt.Println(ids)
	sortedranges := SortRanges(rawranges)
//	fmt.Println(sortedranges)
	flatranges := FlattenRanges(sortedranges)
//	fmt.Println(flatranges)

	fresh := 0
	for _, id := range ids {
		for _, r := range flatranges {
			if inrange(id,r) {
				fresh += 1
			}
		}
	}
	fmt.Println(fresh)
	nfresh := CountFresh(flatranges)
	fmt.Println(nfresh)
}

func CountFresh(ranges []Range) int {
	nfresh := 0

	for _, r := range ranges {
		nfresh += r.end-r.start+1
	}

	return nfresh
}


func FlattenRanges(rawranges []Range) []Range {

	var flatranges []Range

	sortedranges := SortRanges(rawranges)

	flatranges = append(flatranges, sortedranges[0])
	sortedranges = sortedranges[1:]
	for _, r := range sortedranges {
		for i, f := range flatranges {
			if inrange(r.start, f) && r.end > f.end {
				flatranges[i].end = r.end
			} else if i == len(flatranges)-1 && r.start > f.end {
				flatranges = append(flatranges, r)
			} 
		}
	}
	return flatranges
}

func inrange(i int, r Range) bool {
	return i >= r.start && i <= r.end
}

func SortRanges( rawranges []Range) []Range {
	slices.SortFunc(rawranges, func(a, b Range) int {
		if a.start < b.start {
			return -1
		} else if a.start > b.start {
			return 1
		}
		return 0
	})
	return rawranges
}

func ReadDB() ([]Range, []int) {

	var rawrange []Range
	var ID []int
	rangesect := true

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			rangesect = false
//			fmt.Println("blank")
		} else if rangesect {
			srange := strings.Split(line,"-")
			start, _ := strconv.Atoi(srange[0])
			end, _ := strconv.Atoi(srange[1])
			rawrange = append(rawrange, Range{start, end})
		} else {
			i, _ := strconv.Atoi(line)
			ID = append(ID, i)
		}	
	}
	return rawrange, ID
}

