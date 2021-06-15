package arrays

import "fmt"

// Sum sums the passed array in a for range loop
func Sum(array [6]int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

// SumAlt sums the passed slice in a simple for loop
func SumAlt(slice []int) int {
	result := 0
	for i := 0; i < len(slice); i++ {
		result += slice[i]
	}
	fmt.Printf("Printf slice: %v", slice)
	fmt.Println()
	fmt.Println("Println slice:", slice)
	fmt.Printf("Type of var slice: %T", slice)
	fmt.Println()
	return result
}