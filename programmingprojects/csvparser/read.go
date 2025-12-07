package main

import (
	"bufio"
	"fmt"
)

func parser(str string) (map[string]string){
	result := map[string]string{}
	reader := bufio.NewReader(os.Stdin)

	strvalues := []string{}
	currentstring := ""
	for _, val := range str{
		if val == "\n"{

		}
		currentstring += val 
		
	}





	return result


}


func main(){
	input := `
	name=Rahul
	age=21
	city=bangalore
	`

	out := parser(input)
	fmt.Println(out)
}