package main

import (
	"flag"
	"fmt"
	"os"
)

func validateString(path *string) bool {
	return *path != ""
}

func main() {

	pathPtr := flag.String("path", "", "path to torrent file")
	flag.Parse()

	if !validateString(pathPtr) {
		fmt.Printf("%v \r\n", "Path to torrent file cannot be emtpy")
		return
	}

	file, err := os.Open(*pathPtr)

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
