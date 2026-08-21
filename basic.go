package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

/*
 * Fetches json text from data.json and stores them in the global 'list'.
 */

func FetchFromFile() {
	list = list[:0]
	data, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}
	_ = json.Unmarshal(data, &list)
}

/*
 * Converts and writes the data in 'list' into data.json.
 */

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

/*
 * Checks if a path exists. Note that (false, err) output stands for
 * Whether the path exists or now is unknown.
 */

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

/*
 * Receives a string, checks if it is numeric and the corresponding index is in
 * the range of the to-do list (os.Exit(1) if not). Returns the numeric string
 * converted to int.
 */

func HandleStringIndex(s string) int {
	m, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("error: invalid index.\nindex must be an integer number.")
		return 1
		os.Exit(1)
	}
	if m <= 0 || m >= len(list)+1 {
		fmt.Println("error: invalid index.")
		return 1
		os.Exit(1)
	}
	return m
}

func PrintHelpMsg() {
	fmt.Println("error: invalid argument.\nsupported commands:")
	fmt.Println("todos add <text>   # add an item to the todo list")
	fmt.Println("todos drop <index> # drop a certain item")
	fmt.Println("todos clear        # clear all items")
	fmt.Println("todos show         # show all existing items")
	fmt.Println("todos pr <index>   # 'pr' for 'prioritize'. place a certain item on the top")
	return
}
