package main

import "fmt"

func main() {

	var x int = 1
	var y int = 15

	var p *int
	var k *int

	p = &x
	k = &y

	fmt.Println(p)
	fmt.Println(k)
	fmt.Println(x)
	fmt.Println(y)

	p, k = k, p

	fmt.Println(p)
	fmt.Println(k)
	fmt.Println(x)
	fmt.Println(y)

	*p, *k = *k, *p

	fmt.Println(p)
	fmt.Println(k)
	fmt.Println(x)
	fmt.Println(y)
}
