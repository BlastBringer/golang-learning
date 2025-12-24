package main 

import("fmt" 
"bufio"
"os"
"strings")

func cleanInput(text string) []string{
	return strings.Fields(text)
}

func commandExit() error{
	fmt.Println("Closing the pokedex...")
	os.Exit(0)
	return nil
}

func commandhelp(allCmds map[string]cliCommand) error{
	fmt.Println("Welcome to the pokedex!")
	fmt.Println("Usage:")
	for _, val := range allCmds{
		fmt.Printf("%s : %s\n", val.name, val.description)
	}
	return nil
}

type cliCommand struct{
	name string 
	description string 
	callback func() error
}




func main(){
	var allCmds map[string]cliCommand
		allCmds = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Help about the pokedex",
			callback: func() error {
				return commandhelp(allCmds)
			},
		},
	}
	scanner := bufio.NewScanner(os.Stdin)
	for{
		fmt.Print("Pokedex > ")
		if !scanner.Scan(){
			break 
		}
			command := cleanInput(scanner.Text())[0]
			cmd, ok := allCmds[command]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}
			cmd.callback()
		
	}
}