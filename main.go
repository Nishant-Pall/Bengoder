package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(".torrent")

	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer file.Close()

	reader := newResp(file)
	val, err := reader.Decode()

	if err != nil {
		return
	}

	fmt.Println(val)
}
