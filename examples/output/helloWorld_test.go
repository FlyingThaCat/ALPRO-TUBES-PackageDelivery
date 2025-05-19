package output

import (
	"bytes"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call the function
	HelloWorld()

	// Restore stdout and read output
	w.Close()
	os.Stdout = old
	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	got := buf.String()
	want := "Hello, World!\n"

	if got != want {
		t.Errorf("HelloWorld() = %q; want %q", got, want)
	}
}
