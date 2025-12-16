package main

import "fmt"

func main() {
	// normal sytax
	// var arrayname [size]elementtype

	var nums [5]int
	fmt.Println(nums)

	nums[4] = 20
	fmt.Println(nums)
	nums[0] = 9
	fmt.Println(nums)

	fruits := [4]string{"Apple", "BANANA", "KIVI", "Avagardo"}
	fmt.Println(fruits)

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// important concepts

	ori := [3]int{1, 2, 3}
	copy := ori
	copy[2] = 20
	fmt.Println(ori)
	fmt.Println(copy)

	// underscore are blank values used to store the unused values
	b := 2
	_ = b

	// comparing the arrays
	array1 := [3]int{1, 2, 3}
	array2 := [3]int{1, 2, 3}

	fmt.Println(array1 == array2)

	// multidimentional array in golang
	var matrix [3][3]int = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{6, 7, 8},
	}

	fmt.Println(matrix)
}
