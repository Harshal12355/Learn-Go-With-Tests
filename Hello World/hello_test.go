package main

import "testing"

func TestHello(t *testing.T){
	got := Hello_World()
	want := "Hello World"

	if got != want{
		t.Errorf("got: %q and want: %q", got, want)
	}
} 