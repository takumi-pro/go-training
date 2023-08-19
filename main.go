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

	// if with a short statement
	if n := 101; n < 100 {
		fmt.Println("n smaller than 100")
	} else {
		fmt.Println("n larger than 100")
	}

	// pointer
	// ポインタ変数の宣言
	i := 100
	var p *int = &i
	fmt.Printf("i pointer: %p\n", p)
	fmt.Printf("i: %d\n", *p)

	// struct
	type User struct {
		name  string
		age   int
		email string
	}

	// structの宣言
	u := User{"taro", 23, "test@test.com"}
	u.age = 100

	// structは参照渡しではない
	// uをcuにコピーしてcu.emailを変更してもuへの影響はない
	cu := u
	cu.email = "dummy"

	// structのポインタを渡す
	pu := &u
	pu.email = "pointer email"

	fmt.Printf("User struct: %+v\n", u)
	fmt.Printf("User struct: %+v\n", cu)
	fmt.Printf("User struct: %+v\n", pu)
}
