package dependencyInjection

import (
	"bytes"
	//"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "chris")

	got := buffer.String()
	want := "Hello chris"
	if got != want {
		t.Errorf("got: %v, want: %v", got, want)
	}
}