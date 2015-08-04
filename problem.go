package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	FastSolve(os.Stdin, func(k *int64) {
		fmt.Println(*k)
	})
}

// SlowSolve - slow solution for 'DifferentProblem' kattis problem, uses buffio.NewScanner
func SlowSolve(r io.Reader, displayer func(*int64)) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, " ")

		if len(values) == 2 {
			n, _ := strconv.ParseInt(values[0], 10, 64)
			k, _ := strconv.ParseInt(values[1], 10, 64)

			result := n - k
			if result < 0 {
				result *= -1
			}

			displayer(&result)
		} else {
			panic("something went wrong - do data to process")
		}
	}
}

// FastSolve - fast solution for 'DifferentProblem' kattis problem, uses binary ioutil.ReadAll()
func FastSolve(r io.Reader, displayer func(*int64)) {
	//assumes POSITIVE INTEGERS. Check v for '-' if you have negative.
	var buf []byte
	buf, _ = ioutil.ReadAll(r)
	intOne := int64(-1)

	num := int64(0)
	found := false
	for _, v := range buf {
		if '0' <= v && v <= '9' {
			num = 10*num + int64(v-'0') //could use bitshifting here.
			found = true
		} else if found {
			if intOne == -1 {
				// first number from the pair
				intOne = num
			} else {
				// second number from the pair - solve and display
				result := solveMe(&intOne, &num)
				displayer(result)
				intOne = -1
			}
			found = false
			num = 0

		}
	}
	if found {
		if intOne == -1 {
			intOne = num
		} else {
			// second number from the pair - solve and display
			result := solveMe(&intOne, &num)
			displayer(result)
		}
		found = false
		num = 0
	}
}

func solveMe(a *int64, b *int64) *int64 {
	result := *a - *b
	if result < 0 {
		result *= -1
	}
	return &result
}
