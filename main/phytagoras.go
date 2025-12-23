package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"example.com/utility"
)

func main() {
	log.SetPrefix("Phytagoras: ")
	log.SetFlags(0)
	reader := bufio.NewReader(os.Stdin)
	var legA int
	var legB int
	var hyp int

	fmt.Println("Hi! Welcome to Phytagoras Calculator!\n")
	time.Sleep(300 * time.Millisecond)

	fmt.Println("Guides:\n1. You will be prompted to input 3 values\n2. Put the value of 0 to the sides that you want to calculate\n3. Enjoy")
	time.Sleep(300 * time.Millisecond)

	fmt.Print("\nPlease input your value for Leg A: ")
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	legA, err = strconv.Atoi(input)
	time.Sleep(300 * time.Millisecond)

	fmt.Print("Please input your value for Leg B: ")
	input, err = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	legB, err = strconv.Atoi(input)
	time.Sleep(300 * time.Millisecond)

	fmt.Print("Please input your value for Hypotenuse: ")
	input, err = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	hyp, err = strconv.Atoi(input)
	time.Sleep(300 * time.Millisecond)

	fmt.Println("\nProcessing...")
	time.Sleep(300 * time.Millisecond)

	result, err := utility.CalculatePhytagoras(legA, legB, hyp)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nYour result is:")

	for index, value := range result {
		fmt.Printf("%v -> %d\n", index, value)
	}
}
