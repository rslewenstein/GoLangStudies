package main

import (
	"fmt"
	"os"
)

func main() {
	fl, err := os.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	tx, err := fl.WriteString("Deus abençoou \n Oh, vitória chegou \n Oh, pai, só gratidão \n oh, mãe, só gratidão")
	if err != nil {
		fmt.Println(err)
		fl.Close()
		return
	}
	fmt.Println(tx, "bytes written")
	err = fl.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}
