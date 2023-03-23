package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println("Nilai i = ", i)
	}
	for j := 0; j < 5; j++ {
		fmt.Println("Nilai j = ", j)
	}
	for i, huruf := range "CAMAPBO" {
		fmt.Printf("%#U starts at byte position %d\n", huruf, i)
	}
	for j := 6; j <= 10; j++ {
		fmt.Println("Nilai j  = ", j)
	}
}
