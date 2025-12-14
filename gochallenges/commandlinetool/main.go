package main

import (
	"bufio"
	"fmt"
	"os"
)

func countLines(file *os.File) int {
	scanner := bufio.NewScanner(file)
	lines := 0
	for scanner.Scan(){
		lines++
	}
	return lines
}

func countChars(file *os.File) int{
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	chars := 0
	for scanner.Scan(){
		chars++
	}

	return chars
}

func countWords(file *os.File) int{
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	words := 0
	for scanner.Scan(){
		words++
	}
	return words
}
func lenfile(file *os.File) int{
	buf := make([]byte, 1024)
	bytes := 0
	for {
		n, err := file.Read(buf)
		bytes += n
		if err != nil{
			break 
		}
	}

	return bytes
}



func main(){
	if len(os.Args) < 2{
		fmt.Println("Usage : ccwc [-l|-w|-c|-m] file")
		return 
	}
	command := os.Args[1]
	var file *os.File
	var filename string
	if len(os.Args) > 2{
		f, err := os.Open(os.Args[2])
		if err != nil{
			fmt.Println("Error while getting file path!")
		}
		defer f.Close()
		file = f
		filename = os.Args[2]
	} else {
		file = os.Stdin
	}

	
	
	switch command{
	case "-c":
		fmt.Println(lenfile(file), filename)
	case "-l":
		fmt.Println(countLines(file), filename)
	case "-w":
		fmt.Println(countWords(file), filename)
	case "-m":
		fmt.Println(countChars(file), filename)
	default:
		fmt.Println("No command mentioned")
	}
	
}