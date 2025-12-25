package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args

	if len(args) < 2 {
		fmt.Printf("%v \r\n", "Please enter correct number of arguments")
		return
	}

	filePath := os.Args[1]

	if filePath == "" {
		fmt.Printf("%v \r\n", "Path to torrent file cannot be emtpy")
		return
	}

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("%v \r\n", err)
		return
	}
	defer file.Close()

	reader := newResp(file)
	val, err := reader.Decode()

	if err != nil {
		return
	}

	fmt.Printf("%v \r\n", val)
}
