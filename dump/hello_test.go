package main 

import (
	"testing"
)



func TestHello(t *testing.T){
	got := Helloworld("Chris")
	want := "Hello, Chris"

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}

}
