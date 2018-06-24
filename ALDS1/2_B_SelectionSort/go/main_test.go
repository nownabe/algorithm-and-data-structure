package main

import (
	"bytes"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	input := "6\n5 6 4 2 1 3"
	expect := "1 2 3 4 5 6\n4"

	buf := new(bytes.Buffer)
	buf.WriteString(input)

	res := run(buf)

	if res != expect {
		t.Fatalf("Expect %s but actual %s for %s", expect, res, input)
	}
}
