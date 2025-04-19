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
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
}

func TestHelloLanguage(t *testing.T){
	t.Run("in spanish",func(t *testing.T){
		got := Hello_lang("Brad", "Spanish")
		want := "Hola Brad"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in french",func(t *testing.T){
		got := Hello_lang("Brad", "French")
		want := "Bonjour Brad"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in english",func(t *testing.T){
		got := Hello_lang("Brad", "English")
		want := "Hello Brad"
		assertCorrectMessage(t, got, want)
	})

	t.Run("if no langauge specified",func(t *testing.T){
		got := Hello_lang("Brad", "")
		want := "Hello Brad"
		assertCorrectMessage(t, got, want)
	})
	t.Run("if no langauge specified",func(t *testing.T){
		got := Hello_lang("Brad", "spAnISH")
		want := "Hola Brad"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}