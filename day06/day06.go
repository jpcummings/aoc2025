package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

type Prob struct {
	operand []int
	operator string
}


func main() {
	//probs := ReadProblems()
	//fmt.Println(probs)
	//answer := EvalProblems(probs)
	//fmt.Println(answer)

	probs := ReadProblems2()
	answer := EvalProblems(probs)
	fmt.Println(answer)
}

func EvalProblems(probs []Prob) int {
	var sum int = 0
	var panswer int = 0

	for _, p := range probs {
		if p.operator == "*" {
			panswer = 1
			for _, n := range p.operand {
				panswer *= n
			}
		} else if p.operator == "+" {
			panswer = 0
			for _, n := range p.operand {
				panswer += n
			}
		} else {
			fmt.Println("unknown operator ", p.operator)
		}
		sum += panswer
	}

	return sum
}


func ReadProblems2() []Prob {

	lines := []string{}
	var probs []Prob

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	operands := []int{}
	operator := ""

	for icol := len(lines[0])-1; icol >=0; icol-=1 {
		snum := ""
		for irow:=0; irow < len(lines); irow+=1 {
			snum += string(lines[irow][icol])
		}

		snum = strings.Trim(snum," ")
		if len(snum) != 0 && (snum[len(snum)-1] == '*' || snum[len(snum)-1] == '+') {
			operator = string(snum[len(snum)-1])
			snum = snum[:len(snum)-1]
		}
		snum = strings.Trim(snum," ")
		num, _ := strconv.Atoi(snum)
		if num != 0 {
			operands = append(operands, num)
		}
		if operator == "*" || operator == "+" {
			probs = append(probs,Prob{operands, operator})
			operands = []int{}
			operator = ""
		}
	}
	return probs
}


func ReadProblems() []Prob {

	scanner := bufio.NewScanner(os.Stdin)

	var probs []Prob
	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Fields(line)
		for iprob, word := range words {
			if len(probs) <= iprob {
				probs = append(probs, Prob{[]int{}, " "})
			}
			if word == "*" || word == "+" {
				probs[iprob].operator = word
			} else {
				nword, _ := strconv.Atoi(word)
				probs[iprob].operand = append(probs[iprob].operand, nword)
			}
		}
	}
	return probs
}