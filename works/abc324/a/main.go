package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var a int
	fmt.Scan(&a)

	result := "Yes"
	for i := 0; i < n-1; i++ {
		var _a int
		fmt.Scan(&_a)
		if a != _a {
			result = "No"
		}
	}

	fmt.Println(result)
}
