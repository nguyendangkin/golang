package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s string = "Xin chào các bạn"

	fmt.Println("Chiều dài:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

	fmt.Println()
	fmt.Println("Tính ký tự đặc biệt", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U bắt đầu ở %d\n", runeValue, idx)
	}

	// so sánh ký tự

	fmt.Printf("\nSử dụng decodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U bắt đầu ở %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}

}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'à' {
		fmt.Println("found so sua")
	}
}
