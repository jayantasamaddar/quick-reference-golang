package concurrency

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("PASSED")

	wg.Wait()

	if msg != "PASSED" {
		t.Error("not found: expected to find 'PASSED'")
	}
}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "PASSED"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, msg) {
		t.Error("not found: expected to find 'PASSED'")
	}
}

func Test_WaitGroups(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	WaitGroups()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello, universe!") {
		t.Error("not found: expected to find 'Hello, universe!'")
	}

	if !strings.Contains(output, "Hello, cosmos!") {
		t.Error("not found: expected to find 'Hello, cosmos!'")
	}

	if !strings.Contains(output, "Hello, world!") {
		t.Error("not found: expected to find 'Hello, world!'")
	}
}
