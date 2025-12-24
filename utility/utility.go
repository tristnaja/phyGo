package utility

import (
	"errors"
	"fmt"
	"math"
)

func CalculatePhytagoras(legA int, legB int, hyp int) (map[string]int, error) {
	sides := make(map[string]int, 3)
	sides["LegA"] = legA
	sides["LegB"] = legB
	sides["Hypotenuse"] = hyp
	var targetSide string

	for index, value := range sides {
		if value == 0 {
			targetSide = index
		}
	}

	if targetSide == "" {
		return nil, errors.New("Please put 0 on one of the values, it will indicates the value that needs to be targeted for the calculations")
	}

	fmt.Printf("Your targeted side is: %v\n", targetSide)

	sidesTmp := sides

	delete(sidesTmp, targetSide)

	if targetSide == "Hypotenuse" {
		result, err := hypotenusePhyCalc(sidesTmp)

		if err != nil {
			return nil, err
		}

		sides[targetSide] = result
	} else {
		result, err := notHypotenusePhyCalc(sidesTmp)

		if err != nil {
			return nil, err
		}

		sides[targetSide] = result
	}

	return sides, nil
}

func DisplayKnownValue(sides map[string]int) {
	fmt.Println("\nHere is the value you have:")
	for index, value := range sides {
		fmt.Printf("%v = %d\n", index, value)
	}
}

func hypotenusePhyCalc(sides map[string]int) (int, error) {
	sideSlices := make([]int, 0, 2)
	var legA int
	var legB int

	for _, value := range sides {
		sideSlices = append(sideSlices, value)
	}

	legA = sideSlices[0]
	legB = sideSlices[1]

	result := int(math.Sqrt(float64(intPow(legA, 2)) + float64(intPow(legB, 2))))

	if result < legA || result < legB {
		return 0, errors.New("Calculation faults. The value of Hypotenuse is smaller than legs, which, hypothetically speaking is impossible.")
	}

	return result, nil
}

func notHypotenusePhyCalc(sides map[string]int) (int, error) {
	sideSlices := make([]int, 0, 2)
	var hyp int
	var leg int

	for _, value := range sides {
		sideSlices = append(sideSlices, value)
	}

	if sideSlices[0] == sideSlices[1] {
		return 0, errors.New("The value of the hypotenuse must not equal to the value of a known leg. It will result with 0, meaning no triangle at all.")
	}

	hyp = sides["Hypotenuse"]
	delete(sides, "Hypotenuse")

	if len(sides) != 1 {
		return 0, errors.New("Process Terminated")
	}

	for _, value := range sides {
		leg = value
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
