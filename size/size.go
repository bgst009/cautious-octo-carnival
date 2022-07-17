package size

import (
	"fmt"
	"unsafe"
)

func intSize() {
	fmt.Printf("unsafe.Sizeof(int(1)): %v\n", unsafe.Sizeof(int(1)))

}

func uintSize() {
	fmt.Printf("unsafe.Sizeof(uint(1)): %v\n", unsafe.Sizeof(uint(1)))
}

func pointerSize() {
	a := 10
	p := &a
	fmt.Printf("unsafe.Sizeof(p): %v\n", unsafe.Sizeof(p))
	fmt.Printf("&a: %p\tp: %v\t&p: %p\n", &a, p, &p)
}

type emptyStruct struct{}

func emptyStructSize() {
	a := emptyStruct{}
	fmt.Printf("unsafe.Sizeof(a): %v\n", unsafe.Sizeof(a))
	fmt.Printf("&a: %p\n", &a)

	b := emptyStruct{}
	fmt.Printf("&b: %p\n", &b)
}

type normalStruct struct {
	a     int
	empty emptyStruct
	b     int
	c     int
}

func normalStructSize() {
	a := normalStruct{
		a:     5,
		empty: emptyStruct{},
		b:     5,
		c:     5,
	}
	fmt.Printf("unsafe.Sizeof(a):\t%v\n", unsafe.Sizeof(a))
	fmt.Printf("unsafe.Sizeof(a.a):\t%v\n", unsafe.Sizeof(a.a))
	fmt.Printf("unsafe.Sizeof(a.empty):\t%v\n", unsafe.Sizeof(a.empty))
	fmt.Printf("unsafe.Sizeof(a.b):\t%v\n", unsafe.Sizeof(a.b))
	fmt.Printf("unsafe.Sizeof(a.c):\t%v\n", unsafe.Sizeof(a.c))
	fmt.Printf("&a: %p\t&a.a: %p\t&a.empty: %p\t&a.b: %p\t&a.c: %p\n", &a, &a.a, &a.empty, &a.b, &a.c)
}
