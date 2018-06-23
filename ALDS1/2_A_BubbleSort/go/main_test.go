package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	input := "5\n5 3 2 4 1"
	expect := "1 2 3 4 5\n8"

	buf := new(bytes.Buffer)
	buf.WriteString(input)

	res := run(buf)

	if res != expect {
		t.Fatalf("Expect %s but actual %s for %s", expect, res, input)
	}
}
