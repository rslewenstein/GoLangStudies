package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("data.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println("Text: ", string(data))
}
