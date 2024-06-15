package main

import (
	"log"

	"github.com/ArtemNovok/readers/readers"
)

func main() {
	mapp, err := readers.ReadFileInc("test1.txt")
	if err != nil {
		panic(err)
	}
	log.Println(mapp)
}
