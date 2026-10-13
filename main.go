package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("****__Calculator__****")
	operation := selectOperation()
	userInput := getUserInput()
	slice := convertToSlice(userInput)
	result := calculate(operation, slice)
	fmt.Println("The numbers provided:", slice)
	fmt.Println("Operation:", operation)
	fmt.Println("The result is:", result)
}

func selectOperation() string {
	fmt.Println("Select operation: (AVG/SUM/MED)")
	var operation string
	for {
		fmt.Scan(&operation)
		operation = strings.ToUpper(operation)
		switch operation {
		case "AVG", "SUM", "MED":
			return operation
		default:
			fmt.Println("Invalid input: enter (AVG/SUM/MED)")
		}
	}
}

func getUserInput() string {
	fmt.Println("Enter numbers using comma separator")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func convertToSlice(text string) []float64 {
	splitText := strings.Split(text, ",")
	floatSlice := make([]float64, 0, len(splitText))
	for _, value := range splitText {
		valueTrimmed := strings.TrimSpace(value)
		elem, err := strconv.ParseFloat(valueTrimmed, 64)
		if err != nil {
			fmt.Println("Invalid input detected - will be skipped")
			continue
		}
		floatSlice = append(floatSlice, elem)
	}
	slices.Sort(floatSlice)
	return floatSlice
}

func AVG(numbers []float64) float64 {
	var sum float64
	for _, value := range numbers {
		sum += value
	}
	result := sum / float64(len(numbers))
	return result
}

func SUM(numbers []float64) float64 {
	var result float64
	for _, value := range numbers {
		result += value
	}
	return result
}

func MED(numbers []float64) float64 {
	switch evenOddCheck(numbers) {
	case true:
		index1 := len(numbers) / 2
		index2 := len(numbers)/2 - 1
		return (numbers[index1] + numbers[index2]) / 2
	case false:
		index := (len(numbers) - 1) / 2
		return numbers[index]
	default:
		return 0
	}
}

func evenOddCheck(numbers []float64) bool {
	if len(numbers)%2 == 0 {
		return true
	} else {
		return false
	}
}

func calculate(operation string, numbers []float64) float64 {
	switch operation {
	case "AVG":
		return AVG(numbers)
	case "SUM":
		return SUM(numbers)
	case "MED":
		return MED(numbers)
	default:
		return 0
	}
}
