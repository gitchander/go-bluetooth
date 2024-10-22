package main

import (
	"fmt"
	"log"

	"github.com/gitchander/go-bluetooth/hw/linux/hciconfig"
)

func main() {
	ais, err := hciconfig.GetAdapters()
	checkError(err)
	for _, ai := range ais {
		fmt.Println(ai)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
