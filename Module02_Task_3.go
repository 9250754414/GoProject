package main

import "fmt"

func main() {

	k := 5
	p := 2
	s := method(&k, &p)

	fmt.Println(s)
}

func method(x *int, y *int) int {

	return (*x) * (*y)
}
