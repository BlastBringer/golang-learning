package main 


import(
	"fmt"
)


const(
	englishHelloPrefix = "Hello, "
	spanish = "Spanish"
	spanishHelloPrefix = "Hola, "
	french = "French"
	frenchHelloPrefix = "Bonjour, "
)



func Helloworld(name string, language string) string{
	if name == ""{
		name = "World"
	}


	prefix := greetingPrefix(language)

	return prefix + name
}

func greetingPrefix(language string) (prefix string){
	switch language{
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix

	}

	return 
}




func main(){
	fmt.Println(Helloworld("Rahul", ""))


}