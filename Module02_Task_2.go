package main

import "fmt"

func main() {

	count, sum := sumint([]int{1, 2, 3, 4, 5})

	fmt.Println(count, sum)
}

func sumint(param []int) (int, int) {

	countParam := len(param)
	sumParam := sum(countParam, param)
	return countParam, sumParam
}

func sum(count int, param []int) int {
	result := 0

	for i := 0; i < count; i++ {
		if i != count {
			result = result + param[i]
		}
	}

	return result
}
