package main 


import "fmt"

func hello(name, language string) string{
	if language == "es"{
		return "Hola, " + name
	}

	if language == "fr"{
		return "Bonjour, "+ name
	}

	return "Hello, " + name
}


//refactored
func Hello(name, language string) string{
	return fmt.Sprintf("%s , %s", greeting[language], name)
}
var greeting = map[string]string{
	"es" : "Hola",
	"fr" : "Bonjour",
	"en" : "Hello",
}



func main(){
	fmt.Println("hello world!")
	fmt.Println(hello("rahul", "en"),
	Hello("rahul", "fr"))
}