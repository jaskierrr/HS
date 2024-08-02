package main

import "fmt"

func main() {
	testArr := []int{1, 1, 2, 6, 6, 2, 3, 5, 5}
	testArr2 := []int{1, 1, 2, 6}

	fmt.Printf("testArr1: %v\n", testArr)
	fmt.Printf("testArr2: %v\n\n", testArr2)

	fmt.Printf("double1: %v\n", double1(testArr))
	fmt.Printf("double2: %v\n", double2(testArr))
	fmt.Printf("double3: %v\n", double3(testArr))
	fmt.Printf("cross: %v\n", cross(testArr, testArr2))
	fmt.Printf("plus: %v\n", plus(testArr, testArr2))
}

func double1(arr []int) []int {
	numMap := make(map[int]bool, len(arr))

	for _, v := range arr {
		numMap[v] = true
	}

	res := make([]int, 0, len(numMap))

	for i := range numMap {
		res = append(res, i)
	}

	return res
}

func double2(arr []int) []int {
	var newSlice []int

	for _, v := range arr {
		if !contains(newSlice, v) {
			fmt.Println(len(newSlice), cap(newSlice))
			newSlice = append(newSlice, v)
		}
	}
	return newSlice
}

func contains(newSlice []int, v int) bool {
	for i := range newSlice {
		if newSlice[i] == v {
			return true
		}
	}
	return false
}

// !!!!!!!!!!!!!!
// задал длину слайсу и все заработало быстро и без танцев с бубном
func double3(arr []int) []int {
	newSlice := make([]int, 0, len(arr))

	for _, v := range arr {
		newSlice = append(newSlice, v)
		fmt.Println(len(newSlice), cap(newSlice))
	}
	return newSlice
}

func cross(first []int, second []int) []int {
	var newSlice []int
	numMap1 := make(map[int]int)
	numMap2 := make(map[int]int)

	for _, v := range first {
		numMap1[v]++
	}
	for _, v := range second {
		numMap2[v]++
	}

	for number := range numMap1 {
		if numMap2[number] > 0 {
			minCount := min(numMap1[number], numMap2[number])
			for i := 0; i < minCount; i++ {
				newSlice = append(newSlice, number)
		}
		}
	}
	return newSlice
}


func plus(first []int, second []int) []int {
	newSlice := append(first, second...)
	return newSlice
}
