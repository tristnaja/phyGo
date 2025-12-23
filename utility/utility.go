package utility

import (
	"errors"
	"fmt"
	"math"
)

func CalculatePhytagoras(legA int, legB int, hyp int) (map[string]int, error) {
	sides := make(map[string]int, 3)
	sides["legA"] = legA
	sides["legB"] = legB
	sides["Hypotenuse"] = hyp
	var targetSide string

	for index, value := range sides {
		if value == 0 {
			targetSide = index
		}
	}

	fmt.Printf("Your targeted side is: %v\n", targetSide)

	sidesTmp := sides

	delete(sidesTmp, targetSide)

	fmt.Println("\nHere is the value you have:")
	for index, value := range sidesTmp {
		fmt.Printf("%v = %d\n", index, value)
	}

	if targetSide == "Hypotenuse" {
		result := hypotenusePhyCalc(sidesTmp)
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

func hypotenusePhyCalc(sides map[string]int) int {
	sideSlices := make([]int, 0, 2)
	var legA int
	var legB int

	for _, value := range sides {
		sideSlices = append(sideSlices, value)
	}

	legA = sideSlices[0]
	legB = sideSlices[1]

	result := int(math.Sqrt(float64(intPow(legA, 2)) + float64(intPow(legB, 2))))

	return result
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

	if sideSlices[0] > sideSlices[1] {
		hyp = sideSlices[0]
		leg = sideSlices[1]
	} else {
		hyp = sideSlices[1]
		leg = sideSlices[0]
	}

	result := int(math.Sqrt(float64(intPow(hyp, 2)) - float64(intPow(leg, 2))))

	return result, nil
}

func intPow(num int, pow int) int {
	if pow == 0 {
		return 1
	}

	if pow == 1 {
		return num
	}

	result := num

	for i := 2; i <= pow; i++ {
		result *= num
	}

	return result
}
