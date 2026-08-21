package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func FetchFromFile() {
	list = list[:0]
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(data, &list)
}

func WriteToFile() {
	data, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		panic(err)
	}
	_ = os.WriteFile("data.json", data, 0644)
}

func ClearFile() {
	file, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func PrintHelpMsg() {
	fmt.Println("error: invalid argument.\nsupported commands:")
	fmt.Println("todos add <text>")
	fmt.Println("todos drop <index>")
	fmt.Println("todos clear")
	fmt.Println("todos show")
	return
}
