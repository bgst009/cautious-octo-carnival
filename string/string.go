package string

import (
	"fmt"
	"unsafe"
)

func stringSize() {
	s1 := "好好学习"
	s2 := "好好学习天天向上"
	fmt.Printf("unsafe.Sizeof(s1): %v\n", unsafe.Sizeof(s1))
	fmt.Printf("unsafe.Sizeof(s2): %v\n", unsafe.Sizeof(s2))
}

func traverseString() {
	s1 := "好好学习天天向上"
	for i := 0; i < len(s1); i++ {
		fmt.Printf("// s1[%d]: %v\n", i, s1[i])
	}

	for i, v := range s1 {
		fmt.Printf("// v%d: %c\n", i, v)
	}
}

func splitString() {
	s1 := "好好学习天天向上"
	fmt.Printf("// string([]rune(s1)[:3]): %v\n", string([]rune(s1)[:3]))
}

func createSlice() {
	//1. 使用数组
	array1 := [5]int{1, 2, 3, 4, 5}
	slice1 := array1[:3]
	fmt.Printf("// slice1: %v\tlen: %d\tcap: %d\n", slice1, len(slice1), cap(slice1))
	//2. 使用 make
	slice2 := make([]int, 5, 10)
	fmt.Printf("// slice2: %v\tlen: %d\tcap: %d\n", slice2, len(slice2), cap(slice2))
	//3. 使用字面量
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf("// slice3: %v\tlen: %d\tcap: %d\n", slice3, len(slice3), cap(slice3))
}

func appendSlice() {
	//1. 使用数组
	array1 := [5]int{1, 2, 3, 4, 5}
	slice1 := array1[:3]
	fmt.Printf("// slice1: %v\tlen: %d\tcap: %d\n", slice1, len(slice1), cap(slice1))
	//2. 使用 make
	slice2 := make([]int, 5, 10)
	fmt.Printf("// slice2: %v\tlen: %d\tcap: %d\n", slice2, len(slice2), cap(slice2))
	//3. 使用字面量
	slice3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Printf("// slice3: %v\tlen: %d\tcap: %d\n", slice3, len(slice3), cap(slice3))

	fmt.Printf("// \"after append\": %v\n", "after append")
	// cap 足够
	slice1 = append(slice1, 6)
	fmt.Printf("// slice1: %v\tlen: %d\tcap: %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("// array1: %v\n", array1)
	slice2 = append(slice2, 1)
	fmt.Printf("// slice2: %v\tlen: %d\tcap: %d\n", slice2, len(slice2), cap(slice2))
	// cap 不够
	slice3 = append(slice3, 12)
	fmt.Printf("// slice3: %v\tlen: %d\tcap: %d\n", slice3, len(slice3), cap(slice3))
	slice1 = append(slice1, 1, 2, 3, 4, 5)
	fmt.Printf("// slice1: %v\tlen: %d\tcap: %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("// array1: %v\n", array1)
}
