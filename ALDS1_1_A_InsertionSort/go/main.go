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

	return insertionSort(n, arr)
}

func insertionSort(n int, arr []int) string {
	buf := new(bytes.Buffer)

	print(buf, arr, n != 1)

	for i := 1; i < n; i++ {
		v := arr[i]
		j := i - 1
		for ; j >= 0 && arr[j] > v; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = v
		print(buf, arr, i < n-1)
	}

	return buf.String()
}

func print(b io.Writer, arr []int, nl bool) {
	for i, v := range arr {
		if i > 0 {
			fmt.Fprintf(b, " ")
		}
		fmt.Fprintf(b, "%d", v)
	}
	if nl {
		fmt.Fprintf(b, "\n")
	}
}
