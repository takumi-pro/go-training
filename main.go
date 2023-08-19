package main

import (
	"fmt"
	"training/mypackage"
)

func main() {
	mypackage.Greeting()

	// for文
	// ()はつけない
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	fmt.Println()

	// 初期化と後処理は任意
	// whileを表現する際はforを使う
	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Printf("sum :%d\n", sum)
	}
}
