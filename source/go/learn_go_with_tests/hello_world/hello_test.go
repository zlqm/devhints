package main

import "testing"

func TestHello(t *testing.T){
	assertCorrectMessage := func(t *testing.T, got, want string){
		t.Helper()  // when error occured, will print caller
		if got != want {
			t.Errorf("got '%q' want '%q'", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T){
		got := Hello("Chirs", "English")
		want := "Hello, Chirs"
		assertCorrectMessage(t, got, want)
	})
	t.Run("say hello world when an empty string is supplied", func(t *testing.T){
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("in Chinese", func(t *testing.T){
		got := Hello("世界", "Chinese")
		want := "你好, 世界"
		assertCorrectMessage(t, got, want)
	})
}
