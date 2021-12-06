package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func getInput() []string {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal("Failed to open!")
	}
	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}

func solutionOne(text *[]string) int64 {
	length := len((*text)[0])
	lines := len(*text)
	sums := make([]int, length)
	var gamma, epsilon string

	for _, line := range *text {
		for i, bit := range line {
			sums[i] += int(bit - 48)
		}
	}

	for _, sum := range sums {
		if sum > lines/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaDec, _ := strconv.ParseInt(gamma, 2, 64)
	epsilonDec, _ := strconv.ParseInt(epsilon, 2, 64)
	return gammaDec * epsilonDec
}

func solutionTwo(text *[]string) int64 {
	oxyRate, _ := strconv.ParseInt(solTwoHelper(*text, 0, true), 2, 64)
	co2Rate, _ := strconv.ParseInt(solTwoHelper(*text, 0, false), 2, 64)
	return oxyRate * co2Rate
}

func solTwoHelper(values []string, position int, common bool) string {
	if len(values) == 1 {
		return values[0]
	}

	var zero, one []string
	for _, value := range values {
		if rune(value[position]) == '1' {
			one = append(one, value)
		} else {
			zero = append(zero, value)
		}
	}

	if len(one) >= len(zero) == common {
		return solTwoHelper(one, position+1, common)
	}
	return solTwoHelper(zero, position+1, common)
}

func main() {
	text := getInput()
	fmt.Println("Answer 1:", solutionOne(&text))
	fmt.Println("Answer 2:", solutionTwo(&text))
}
