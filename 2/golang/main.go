package main

import (
	"fmt"
)

func main() {
	var a int = 1
	var b int = 1
	var seq []int = []int{1, 1}
	for {
		i := a + b
		a = b
		b = i
		if i >= 4000000 {
			var combined int = 0
			for _, v := range seq {
				if v%2 == 0 {
					combined = combined + v
				}
			}
			fmt.Printf("%d\n", combined)
			break
		}
		seq = append(seq, i)
	}
}
