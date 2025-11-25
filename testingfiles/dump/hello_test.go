package main 

import (
	"testing"
)


//subtests
func TestHello(t *testing.T){
	t.Run("saying hello to people", func(t *testing.T){
		got := Helloworld("Rahul", "")
		want := "Hello, Rahul"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T){
		got := Helloworld("", "")
		want := "Hello, World"	
		assertCorrectMessage(t, got, want)
	})


	t.Run("in Spanish", func(t *testing.T){
		got := Helloworld("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in french", func(t *testing.T){
		got := Helloworld("Rahul", "French")
		want := "Bonjour, Rahul"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got ,want string){
	t.Helper() 
	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
}


/* cycle repeats for testing
write a test
make the compiler pass 
run the test, see that it fails and check the error message is meaningful
write enough code to make the test pass
refactor

*/