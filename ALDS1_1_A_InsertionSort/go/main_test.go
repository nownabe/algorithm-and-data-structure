package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	input := "6\n5 2 4 6 1 3"
	expect := `5 2 4 6 1 3
2 5 4 6 1 3
2 4 5 6 1 3
2 4 5 6 1 3
1 2 4 5 6 3
1 2 3 4 5 6`

	buf := new(bytes.Buffer)
	buf.WriteString(input)

	res := run(buf)

	if res != expect {
		t.Fatalf("Expect %s but actual %s for %s", expect, res, input)
	}
}
