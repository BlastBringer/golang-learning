package main 

import (
	"fmt"
	"encoding/json"
	"os"
	"path/filepath"
)

func checkValid(data []byte) bool{
	var js interface{}
	err := json.Unmarshal(data, &js)
	return err == nil
}


func checkValidFolder(folderPath string) []string{
	results := []string {}
	files, err := os.ReadDir(folderPath)
	if err != nil{
		fmt.Println("Failed to read folder: ", err)
		return []string{}
	}

	for _, file := range files{
		fullpath := filepath.Join(folderPath, file.Name())
		data, err := os.ReadFile(fullpath)
		if err!= nil{
			fmt.Println("ERROR READING FILE: ", file)
		}

		if checkValid(data){
			results = append(results, fmt.Sprintf("%s is a valid json", fullpath))
		} else{
			results = append(results, fmt.Sprintf( "%s is invalid JSON", fullpath))
		}
		
	}
	return results
}

type Values struct{
	key1 *bool `json:"key1"`
	key2 *bool `json:"key2"`
	key3 *int `json:"key3"`
	key4 *string `json:"key4"`
	key5 *int `json:"key5"`
}

func main(){
	filePath := "tests1/step3/valid.json"
	data, err := os.ReadFile(filePath)
	if err != nil{
		fmt.Println("Error reading file", err)
		return
	}

	var v Values
	err = json.Unmarshal(data, &v)
	if err != nil{
		fmt.Println("Error while reading file: ", err)
	} 
	fmt.Println(v)




}