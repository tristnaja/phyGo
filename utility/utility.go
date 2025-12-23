package utility

import (
	"errors"
	"math"
	"slices"
)

func CalculatePhytagoras(sideA int, sideB int, sideC int) (int, error) {
	sides := []int{sideA, sideB, sideC}
	slices.Sort(sides)

	if sides[0] != 0 {
		return -1, errors.New("Please input one of the 3 sides to be the value of 0.")
	}

	targetSide := sides[0]

	targetSide = int(math.Sqrt(float64(intPow(sides[1], 2)) + float64(intPow(sides[2], 2))))

	return targetSide, nil
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
