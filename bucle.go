package main

import "fmt"

func main() {
	fmt.Println(Repeat("aaa"))
}

const repeatCount = 5

func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
