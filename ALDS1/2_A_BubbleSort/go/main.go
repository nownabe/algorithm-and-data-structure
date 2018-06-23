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

	return bubbleSort(n, arr)
}

func bubbleSort(n int, arr []int) string {
	buf := new(bytes.Buffer)

	cnt := 0

	flag := true

	for flag {
		flag = false
		for i := n - 1; i > 0; i-- {
			if arr[i] < arr[i-1] {
				swap(arr, i, i-1)
				cnt++
				flag = true
			}
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
