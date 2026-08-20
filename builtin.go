package main

import (
	"fmt"
	"strconv"
)

type cmd_struct struct {
	name     string
	executor func([]string) int
}

type item struct {
	content string
}

func Add(args []string) int {
	if len(args) != 1 {
		fmt.Println("Usage: todos add <text>")
		return 1
	}
	FetchFromFile()
	list = append(list, &item{content: args[0]})
	WriteToFile()
	return 0
}

func Show(args []string) int {
	if len(args) != 0 {
		fmt.Println("Usage: todos show")
		return 1
	}
	FetchFromFile()
	for i, m := range list {
		fmt.Println(i+1, m.content)
	}
	return 0
}

func Clear(args []string) int {
	if len(args) != 0 {
		fmt.Println("Usage: todos clear")
		return 1
	}
	ClearFile()
	return 0
}

func Drop(args []string) int {
	if len(args) != 1 {
		fmt.Println("Usage: todos drop <item index>.")
		return 1
	}

	FetchFromFile()

	m, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("error: invalid index.\nindex must be a number.")
	}
	if m <= 0 || m >= len(list)+1 {
		fmt.Println("error: invalid index.")
	}
	list = append(list[:m-1], list[m:]...)

	WriteToFile()

	return 0
}
