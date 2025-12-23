package main

import (
	"fmt"
	"log"

	"example.com/utility"
)

func main() {
	log.SetPrefix("Phytagoras: ")
	log.SetFlags(0)

	result, err := utility.CalculatePhytagoras(6, 0, 10)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nYour result is:")

	for index, value := range result {
		fmt.Printf("%v -> %d\n", index, value)
	}
}
