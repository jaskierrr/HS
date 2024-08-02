package main

import (
	"fmt"
	"strings"
)

func main() {
	in := "hello"
	fmt.Println(reverse(in))

	fmt.Println(find1(in, "ol"))
	fmt.Println(find2(in, "ol"))

	arr := []rune{1055, 1088, 1080, 1074, 1077, 1090}
	fmt.Println(join(arr))
}

func reverse(in string) string {
	str := []rune(in)
	newStr := strings.Builder{}

	for i := len(str) - 1; i != -1; i-- {
		newStr.WriteRune(str[i])
	}

	return newStr.String()
}

func find1(mainStr string, subStr string) (int, int, bool)  {
	firstIndex := strings.Index(mainStr, subStr)
	lastIndex := firstIndex+len(subStr)-1
	ok := false
	if firstIndex != -1 {ok = true}

	return firstIndex, lastIndex, ok
}

func find2(mainStr string, subStr string) (int, int, bool)  {
	ok := false
	firstIndex := -1

	for i := range mainStr {
		if mainStr[i] == subStr[0] {
			if i + len(subStr) > len(mainStr){
				break
			}
			firstIndex = i
			for j := range subStr {
				if mainStr[i+j] == subStr[j] {
					ok = true
				} else {
					ok = false
					firstIndex = -1
				}
			}
			if ok {break}
		}
	}

	lastIndex := firstIndex+len(subStr)-1

	return firstIndex, lastIndex, ok
}

func join(arr []rune) string {
	str := strings.Builder{}
	for i := range arr {
		str.WriteRune(arr[i])
	}

	return str.String()
}
