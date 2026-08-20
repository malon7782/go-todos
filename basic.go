package main

import (
	"fmt"
	"os"
	"strings"
)

func FetchFromFile() {
	list = list[:0]
	f, err := os.ReadFile("data.txt")
	if err != nil {
		panic(err)
	}
	for _, m := range strings.Split(string(f), "\n") {
		if m == "" {
			continue
		}
		list = append(list, &item{content: m})
	}
}

func WriteToFile() {
	f := ""
	for _, m := range list {
		f += m.content
		f += "\n"
	}
	_ = os.WriteFile("data.txt", []byte(f), 0644)
}

func ClearFile() {
	file, err := os.Create("data.txt")
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
