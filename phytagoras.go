package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	log.SetPrefix("Phytagoras Log: ")
	log.SetFlags(0)
	reader := bufio.NewReader(os.Stdin)
	triangle := Triangle{}

	fmt.Print("Hi! Welcome to Phytagoras Calculator!\n\n")

	fmt.Println("Guides:\n1. You will be prompted to input 3 values\n2. Put the value of 0 to the sides that you want to calculate\n3. Enjoy")

	fmt.Print("\nPlease input your value for Leg A: ")
	input, err := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	triangle.LegA, err = strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Please input your value for Leg B: ")
	input, err = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	triangle.LegB, err = strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Please input your value for Hypotenuse: ")
	input, err = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	triangle.Hypotenuse, err = strconv.Atoi(input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nHere is your registered values:")

	DisplayValues(&triangle)

	fmt.Print("\nProcessing...\n")

	result, err := CalculatePhytagoras(&triangle)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("\nYour result is:")

	DisplayValues(&result)
}
