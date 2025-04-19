package main

import "testing"

func TestHello(t *testing.T){
	got := Hello_World()
	want := "Hello World"

	if got != want{
		t.Errorf("got: %q and want: %q", got, want)
	}
} 

func TestHelloYou(t *testing.T){
	t.Run("Wanting to say hello to people", func(t *testing.T){
		got := Hello_You("Chris")
		want := "Hello Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("When name is an empty string then just return Hello World", func(t *testing.T){
		got := Hello_You("")
		want := "Hello"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}