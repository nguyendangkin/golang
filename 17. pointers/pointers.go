package main

import "fmt"

func noPoint(i int) {
	i = 0
}

// đối số i (chứa địa chỉ) có kiểu là một con trỏ có kiểu int
func havePoint(i *int) {
	// trỏ đến địa chỉ i, đổi giá trị ở đó thành 999
	*i = 999
}

func main() {
	i := 1
	fmt.Println("i ban dau:", i)

	noPoint(i)
	fmt.Println("i khi dung ham thuong:", i)

	havePoint(&i) // truyền địa chỉ biến
	fmt.Println("i khi dung con tro:", &i)
}
