package main

import (
	"bytes"
	"fmt"
	"testing"
)

type testCase struct {
	expect int
	input  []int
}

func TestRun(t *testing.T) {
	tests := []testCase{
		testCase{expect: 3, input: []int{5, 4, 1, 3, 4, 3}},
		testCase{expect: -1, input: []int{4, 3, 2}},
	}

	for _, tc := range tests {
		n := len(tc.input)

		buf := new(bytes.Buffer)

		fmt.Fprintf(buf, "%d\n", n)
		for _, v := range tc.input {
			fmt.Fprintf(buf, "%d\n", v)
		}

		a := run(buf)

		if a != tc.expect {
			t.Fatalf("Expect %d but actual %d for %v", tc.expect, a, tc.input)
		}
	}
}
