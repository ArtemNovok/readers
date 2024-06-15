package main

import (
	"fmt"

	"github.com/ArtemNovok/readers/readers"
)

func main() {
	mapp, err := readers.ReadFileIncMap("test1.txt", 2)
	if err != nil {
		panic(err)
	}
	for key, val := range mapp {
		fmt.Printf("%s: %v\n", string(key), val)
	}
}
