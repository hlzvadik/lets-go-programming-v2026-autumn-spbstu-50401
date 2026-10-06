package main

import (
	"fmt"
)

func main() {
	var n int

	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Invalid input n")
		return
	}

	for i := 0; i < n; i++ {
		min := 15
		max := 30

		var k int
		_, err = fmt.Scan(&k)
		if err != nil || k <= 0 {
			fmt.Println("Invalid input k")
			return
		}

		for j := 0; j < k; j++ {
			var operation string
			_, err = fmt.Scan(&operation)
			if err != nil {
				fmt.Println("Invalid input operation")
				return
			}
			var number int
			_, err = fmt.Scan(&number)
			if err != nil || number < 15 || number > 30 {
				fmt.Println("Invalid input number")
				return
			}
			switch operation {
			case ">=":
				if number > min {
					min = number
				}
			case "<=":
				if number < max {
					max = number
				}
			default:
				fmt.Println("Invalid operation")
				return
			}
			if min > max {
				fmt.Println("-1")
			} else {
				fmt.Println(min)
			}
		}
	}
}
