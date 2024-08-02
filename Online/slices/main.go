package main

import (
	"fmt"
)

func main() {
	// блок copy
	mainSl := []string{"Hello", "Dasha", "How are you?"}

	var sl []string
	copy(sl, mainSl)
	fmt.Println(sl) // output: "Hello", "Dasha", "How are you?"

	sl1 := make([]string, 0, 1)
	copy(sl1, mainSl)
	fmt.Println(sl1) // output: "Hello"

	sl3 := make([]string, 2)
	copy(sl3, mainSl)
	fmt.Println(sl3) // output: "Hello", "Dasha",

	sl2 := make([]string, 3)
	copy(sl2, mainSl)
	fmt.Println(sl2) // output: "Hello", "Dasha", "How are you?"

	sl4 := make([]string, 2)
	sl4[0] = "EEE"

	copy(sl4, mainSl)
	fmt.Println(sl4) // output: "EEE", "Hello"

	//блок append
	var (
		one = make([]int, 4)
		two = one[1:3]
	)

	one[2] = 11
	fmt.Println(one, two) // output: {0, 11}, {0, 11}

	one = append(one, 1)

	one[2] = 22
	fmt.Println(one, two) // output: {0, 22}, {0, 0}

	newStr := Reverse("привет")

	fmt.Println(newStr)

	newStr1 := Reverse1("привет")

	fmt.Println(newStr1)
}

func Reverse(in string) string {
	str := []rune(in)
	newStr := make([]rune, 0, len(str))

	for i := len(str) - 1; i != -1; i-- {
		newStr = append(newStr, str[i])
	}

	return string(newStr)
}

func Reverse1(in string) string {
	str := in
	newStr := ""

	for i := len(str) - 1; i != -1; i-- {
		newStr += string(str[i])
	}

	return string(newStr)
}

func Reverse2(in string) (out string) {
	for _, r := range in {
		out = string(r) + out
	}
	return
}
