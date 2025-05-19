package output

import (
	"os"
	"testing"
)

func TestAskName(t *testing.T) {
	// Write simulated input to a temp file
	tmpFile, err := os.CreateTemp("", "input")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name()) // clean up
	_, _ = tmpFile.WriteString("Alice\n")
	tmpFile.Seek(0, 0) // rewind to beginning

	// Save and replace os.Stdin
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = tmpFile

	got := askName()
	want := "Alice"

	if got != want {
		t.Errorf("askName() = %q; want %q", got, want)
	}
}
