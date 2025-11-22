package main 


import(
	"fmt"
)

const englishHelloPrefix = "Hello, "



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



func Helloworld(name string) string{
	return englishHelloPrefix + name
}



func main(){
	fmt.Println(hello("rahul", "en"),
	Hello("rahul", "fr"))


}