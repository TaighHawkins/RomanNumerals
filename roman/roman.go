package roman

import (
	"fmt"
	"strings"
)

func ToInt(input string) (int, error) {
	numeralToInt := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	characters := strings.Split(input, "")
	mainCount := 0
	currentValue := 0
	prevValue := 0

	for _, v := range characters {
		currentValue = numeralToInt[v]
		mainCount += currentValue

		if currentValue > prevValue {
			mainCount -= (2 * prevValue)
		}

		prevValue = currentValue
	}

	fmt.Printf("Converted %v to %d\n", input, mainCount)
	return mainCount, nil
}

type NumeralInfo struct {
	Key   string
	Value int
}

func ToNumerals(input int) (out string, err error) {
	numeralToInt := []NumeralInfo{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}
	working := input
	outString := ""
	for working > 0 {
		for _, v := range numeralToInt {
			for working >= v.Value {
				outString += v.Key
				working -= v.Value
			}
		}
	}
	fmt.Printf("Converted %d to %v\n", input, outString)
	return outString, nil
}
