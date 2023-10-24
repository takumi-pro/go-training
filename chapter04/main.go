package main

import (
	"fmt"
	"math/rand"
)

func main() {
	shadowSample()
	ifSample()
}

func shadowSample() {
	x := 100
	if x > 50 {
		x := 500 // 外側のxの変数は隠されてしまう
		fmt.Printf("x --> %d\n", x)
	}
	fmt.Printf("x --> %d\n", x)
}

func ifSample() {
	// 条件分の前に変数を定義できる
	if n := rand.Intn(10); n > 5 {
		fmt.Println("5以上")
	} else {
		fmt.Println("すくな")
	}
}
