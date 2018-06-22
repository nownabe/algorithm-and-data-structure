package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Println(run(os.Stdin))
}

func run(in io.Reader) int {
	sc := bufio.NewScanner(in)

	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	r := make([]int, n)

	for i := 0; i < n; i++ {
		sc.Scan()
		v, err := strconv.Atoi(sc.Text())
		if err != nil {
			panic(err)
		}
		r[i] = v
	}

	return maximumProfit(n, r)
}

func maximumProfit(n int, r []int) int {
	maxProf := math.MinInt32
	minv := r[0]

	for i := 1; i < n; i++ {
		maxProf = max(maxProf, r[i]-minv)
		minv = min(minv, r[i])
	}

	return maxProf
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
