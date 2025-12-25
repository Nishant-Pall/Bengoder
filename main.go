package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("./fixture/.torrent")

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
