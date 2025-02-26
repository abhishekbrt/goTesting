package main

import(
	"testing"
)

func TestHello(t *testing.T){
	t.Run("saying hello to the people ",func (t *testing.T)  {
		got:=Hello("abhishek")
		want:="Hello abhishek"
		assertCorrectMessage(t,got,want)				
	})

	t.Run("say Hello world when an empty string is provided",func (t *testing.T)  {
		got:=Hello("")
		want:="Hello world"
		assertCorrectMessage(t,got,want)
		
	})
}

func assertCorrectMessage(t testing.TB,got,want string){
	t.Helper()
	if got!=want {
		t.Errorf("got %q  want %q \n",got, want)
	}
}

