package main 

import(
	"testing"
)

var cases = []struct{
		input string 
		expected []string 
	}{
		{
		input : " hello  world  ",
		expected: []string{"hello", "world"},
	},
	}

func TestCleanInput(t *testing.T){

	for _, c := range cases{
		actual := cleanInput(c.input)
		for i:= range actual{
			if actual[i] != c.expected[i]{
				t.Errorf("expected %q at index %d, got %q", c.expected[i], i, actual[i],)
			}
		}
	}

}