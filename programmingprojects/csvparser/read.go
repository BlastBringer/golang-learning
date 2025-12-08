package main

import (
"fmt"
"os"
"encoding/csv"
"strconv"
)
type csvInfo struct{
	name string 
	age int 
	city string
}


func main(){
	file , err := os.Open("hello.csv")
	if err != nil{
		fmt.Println("Error occured while reading a file:", file)
	}

	defer file.Close()
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()
	if err != nil{
		panic(err)
	}


	infoCsv := []csvInfo{}
	for _, line := range lines{
		age, err:= strconv.Atoi(line[1])
		if err!=nil{panic(err)}

		newinfo := csvInfo{
			name : line[0],
			age : age,
			city : line[2],
		}

		infoCsv = append(infoCsv, newinfo)

	}

	fmt.Println(infoCsv)
}