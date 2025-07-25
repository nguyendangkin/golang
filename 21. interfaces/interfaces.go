package main

import "fmt"

// interface động vật biết kiêu
type Animal interface {
	Speak() string
}

// struct gấu
type Bear struct {
	name string
}

// struct mèo
type Cat struct {
	name string
}

// bear thực hiện phương thức speak
func (b Bear) Speak() string {
	return b.name + " gầm gừ!"
}

// cat cũng thực hiện speak
func (c Cat) Speak() string {
	return c.name + " meo meo!"
}

// hàm nhận vào bất kỳ con vật nào biết Speak
func makeSound(a Animal) {
	fmt.Println(a.Speak())
}

// dùng type assertion để kiểm tra có phải gấu không
func detectBear(a Animal) {
	if b, oke := a.(Bear); oke { // check xem a(b) có phải là bear không
		fmt.Println("Đây là con gấu tên: ", b.name)
	}
}

func main() {
	b := Bear{name: "Po"}
	c := Cat{name: "Kitty"}

	makeSound(b)
	makeSound(c)

	detectBear(b)
	detectBear(c)

}
