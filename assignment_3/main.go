package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := os.Args[1]
	file, er := os.Open(filename)
	if er != nil {
		fmt.Println(er)
		os.Exit(1)
	}
	io.Copy(os.Stdout, file)
}
