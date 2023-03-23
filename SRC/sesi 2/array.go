package main

import "fmt"

func main() {

	var numbers = [4]int{1, 2, 3, 4}
	var length = len(numbers)
	for i := 0; i < length; i++ {
		fmt.Println(numbers[i])
	}

	a := []string{"Foo", "Bar", "cek"}
	for i, s := range a {
		fmt.Println(i, s)
	}

}
