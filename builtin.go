package main

import (
	"fmt"
	"strconv"
	"time"
)

type cmd_struct struct {
	name     string
	executor func([]string) int
}

type item struct {
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

func Add(args []string) int {
	if len(args) != 1 {
		fmt.Println("Usage: todos add <text>")
		return 1
	}
	FetchFromFile()
	list = append(list, &item{Content: args[0], Timestamp: time.Now().Format("2006-01-02 15:04:05")})
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
		fmt.Println(i+1, "|", m.Timestamp, "|", m.Content)
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
		return 1
	}
	if m <= 0 || m >= len(list)+1 {
		fmt.Println("error: invalid index.")
		return 1
	}
	list = append(list[:m-1], list[m:]...)

	WriteToFile()

	return 0
}
