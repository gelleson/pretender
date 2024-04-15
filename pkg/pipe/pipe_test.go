package pipe

import (
	"bytes"
	"os"
	"testing"
)

func TestRead(t *testing.T) {
	// Simulate stdin with a buffer containing the test data
	content := "test data"
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // clean up

	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Seek(0, 0); err != nil {
		t.Fatal(err)
	}

	// Temporarily replace os.Stdin with our test data file
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }() // Restore original Stdin
	os.Stdin = tmpfile

	// Call the Read function and check for expected output
	data, err := Read()
	if err != nil {
		t.Errorf("Read() error = %v, wantErr %v", err, false)
	}
	if !bytes.Equal(data, []byte(content)) {
		t.Errorf("Read() = %v, want %v", data, []byte(content))
	}

	// Close the tmpfile and restore the original Stdin
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
}
