package main

import (
	"fmt"
	"strings"
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

	// array
	var a [2]int
	a[0] = 1
	a[1] = 2
	var e [3][1]uint8
	fmt.Println(e)

	n := [5]int{1, 2, 3, 4, 5}
	fmt.Println(n)

	slice()
	maps()
}

func slice() {
	// slice
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	// make関数を使用することで動的配列を作成できる
	a := make([]int, 5)
	b := make([]int, 5, 10)

	// append関数でsliceに要素を追加
	s = append(s, 100)

	// range
	for i, v := range s {
		fmt.Printf("s slice %d:  %d\n", i, v)
	}
	fmt.Printf("len=%d cap=%d %v\n", len(a), cap(a), a)
	fmt.Printf("len=%d cap=%d %v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}

// map
func maps() {

	type User struct {
		name string
		age  int
	}

	// mapの宣言
	var m = make(map[string]int)
	m["key1"] = 10
	m["key2"] = 20

	// 型にstructを用いる
	var sm = make(map[string]User)
	sm["key1"] = User{"taro", 23}
	sm["key2"] = User{"jiro", 12}

	s := "i am takumi"
	sp := strings.Split(s, " ")
	for _, v := range sp {
		fmt.Println(v)
	}

	fmt.Println(m)
	fmt.Println(sm)
}
