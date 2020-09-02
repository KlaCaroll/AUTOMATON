package main

import (
	"fmt"
)

func main() {
	var i int32 = 4
	fmt.Println(&i, i)

	var t [2]int32 = [2]int32{1, 2}
	fmt.Printf("%p %v\n", &t, t)
	fmt.Printf("%p\n", &t[0])
	fmt.Printf("%p\n", &t[1])

	var tt = [2][2]int32{
		[2]int32{1, 2},
		[2]int32{3, 4},
	}
	fmt.Println(tt)

	var a int32 = 2
	fmt.Println(a)
	f(&a)
	fmt.Println(a)
	dontChange(a)
	fmt.Println(a)

	var array = []int32{1,2}
	fmt.Println(array)
	changeArray(array)
	fmt.Println(array)
}

func dontChange(i int32) {
	i += 1
}

func f(i *int32) {
	*i += 1
}

func changeArray(a []int32) {
	a[0] = 5
	a[1] = 6
}
