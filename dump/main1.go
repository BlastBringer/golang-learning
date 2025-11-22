package main 


import(
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola, "
const french = "French"
const frenchHelloPrefix = "Bonjour, "





func Helloworld(name string, language string) string{
	if name == ""{
		name = "World"
	}

	if language == french{
		return frenchHelloPrefix + name
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	return englishHelloPrefix + name
}




func main(){
	fmt.Println(Helloworld("Rahul", ""))


}