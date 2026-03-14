package hellopackage

import (
	"io"
	"os"
	"testing"
)

func TestHello(t *testing.T) {
	r, w, _ := os.Pipe()
	old := os.Stdout
	os.Stdout = w

	Hello()

	w.Close()
	os.Stdout = old

	out, _ := io.ReadAll(r)

	got := string(out)
	want := "hello, from the package\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
