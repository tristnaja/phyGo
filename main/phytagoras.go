package main

import (
	"fmt"
	"log"

	"example.com/utility"
)

func main() {
	log.SetPrefix("Phytagoras: ")
	log.SetFlags(0)

	result, err := utility.CalculatePhytagoras(6, 8, 0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
