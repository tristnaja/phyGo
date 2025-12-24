package main

import (
	"errors"
	"fmt"
	"math"
)

type Triangle struct {
	LegA       int
	LegB       int
	Hypotenuse int
}

func CalculatePhytagoras(sides *Triangle) (Triangle, error) {
	if sides.Hypotenuse != 0 {
		result, err := notHypotenusePhyCalc(sides)

		if err != nil {
			return *sides, err
		}

		if sides.LegA == 0 {
			sides.LegA = result

			return *sides, nil
		}

		sides.LegB = result

		return *sides, nil
	}

	result, err := hypotenusePhyCalc(sides)

	if err != nil {
		return *sides, nil
	}

	sides.Hypotenuse = result

	return *sides, nil
}

func DisplayValues(sides *Triangle) {
	fmt.Printf("- Leg A: %d\n", sides.LegA)
	fmt.Printf("- Leg B: %d\n", sides.LegB)
	fmt.Printf("- Hypotenuse: %d\n", sides.Hypotenuse)
}

func hypotenusePhyCalc(sides *Triangle) (int, error) {
	legA := sides.LegA
	legB := sides.LegB

	hyp := int(math.Sqrt(float64(intPow(legA, 2)) + float64(intPow(legB, 2))))

	if hyp < legA || hyp < legB {
		return 0, errors.New("Calculation faults. The value of Hypotenuse is smaller than legs, which, hypothetically speaking is impossible.")
	}

	return hyp, nil
}

func notHypotenusePhyCalc(sides *Triangle) (int, error) {
	hyp := sides.Hypotenuse
	var leg int

	if sides.LegA == 0 {
		leg = sides.LegB
	} else if sides.LegB == 0 {
		leg = sides.LegA
	}

	if hyp == leg {
		return 0, errors.New("Invalid input. In the case of phytagoras, the value of hypotenuse cannot be equal with a value of a leg, because it means that there is no Triangle that exists")
	}

	if hyp < leg {
		return 0, errors.New("In the case of searching the unknown value of 1 leg, the value of the known leg must NOT be bigger than the value of hypotenuse.")
	}

	result := int(math.Sqrt(float64(intPow(hyp, 2)) - float64(intPow(leg, 2))))

	return result, nil
}

func intPow(base int, exponent int) int {
	if exponent == 0 {
		return 1
	}

	if exponent == 1 {
		return base
	}

	result := base

	for i := 2; i <= exponent; i++ {
		result *= base
	}

	return result
}
