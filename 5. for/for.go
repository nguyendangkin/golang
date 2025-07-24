package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	for {
		fmt.Println("Loop")
		// lặp vô tận cho đến khi có break thì thoát
		break
	}

	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
