package main

import "craftinginterpreters/scanning"

func main() {
	err := scanning.OpenFile("example_input.txt")
	if err != nil {
		return
	}
}
