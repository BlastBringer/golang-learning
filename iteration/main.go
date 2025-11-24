package iteration 


import "strings"


func Repeat(character string, repeatcount int) string{
	var repeated strings.Builder

	for i := 0; i < repeatcount; i++{
		repeated.WriteString(character)
	}

	return repeated.String()
}