package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(run(os.Stdin))
}

func run(in io.Reader) string {
	sc := bufio.NewScanner(in)

	sc.Scan()
	n, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}

	sc.Scan()
	arrs := strings.Split(sc.Text(), " ")
	arr := make([]int, n)
	for i, v := range arrs {
		if arr[i], err = strconv.Atoi(v); err != nil {
			panic(err)
		}
	}

	return SelectionSort(n, arr)
}

func SelectionSort(n int, arr []int) string {
	buf := new(bytes.Buffer)

	cnt := 0

	for i := 0; i < n; i++ {
		minIdx := i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		if i != minIdx {
			swap(arr, i, minIdx)
			cnt++
		}
	}

	for i, v := range arr {
		if i > 0 {
			fmt.Fprintf(buf, " ")
		}
		fmt.Fprintf(buf, "%d", v)
	}
	fmt.Fprintf(buf, "\n%d", cnt)

	return buf.String()
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
