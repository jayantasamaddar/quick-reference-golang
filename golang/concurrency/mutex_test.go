package concurrency

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_MutexDemo(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	MutexDemo()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "$34320.00") {
		t.Error("Wrong balance returned")
	}

}
