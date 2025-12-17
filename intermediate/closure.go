package main

import "fmt"

func main() {

	result := adder()

	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())

	subtracker := func() func(int) int {
		cnt := 99
		return func(x int) int {
			cnt -= x
			return cnt
		}
	}()

	res := subtracker

	fmt.Println(res(1))
	fmt.Println(res(1))
	fmt.Println(res(1))
	fmt.Println(res(1))

}

func adder() func() int {
	i := 0
	fmt.Println("The value of i is ", i)
	return func() int {
		i++
		fmt.Println("added to i")
		return i
	}
}
