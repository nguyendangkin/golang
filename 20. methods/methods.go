package main

import "fmt"

type Bear struct {
	shirt string
}

func (b Bear) changeShirtWrong() {
	b.shirt = "áo đỏ"
	fmt.Println(b)
}

func (b *Bear) changeShirtRight() {
	b.shirt = "áo đỏ"
}

func main() {
	b := Bear{shirt: "áo xanh"}

	// không có con trỏ
	b.changeShirtWrong()

	// con trỏ b (đỉa chỉ)
	// cũng có thể như vầy bb := &b
	(&b).changeShirtRight()
	fmt.Println("có con trỏ", b.shirt)

}
