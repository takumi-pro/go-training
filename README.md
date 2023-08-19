# Go Training

## packages
- 外部に公開する場合は関数を大文字から始める必要がある。
- 小文字で始まる関数はexportされない

## for

```Go
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
```

## if
```Go
// if with a short statement
	if n := 101; n < 100 {
		fmt.Println("n smaller than 100")
	} else {
		fmt.Println("n larger than 100")
	}
```

## pointer
```Go
	// pointer
	// ポインタ変数の宣言
	i := 100
	var p *int = &i
	fmt.Printf("i pointer: %p\n", p)
	fmt.Printf("i: %d\n", *p)
```

## struct

```Go
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
```

## ATour of Go解答
### Exercise: Slices
```Go
package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	u := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		s := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			s[x] = uint8(x*y)
		}
		u[y] = s
	}
	return u
}

func main() {
	pic.Show(Pic)
}
```

### Exercise: Maps

```Go
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var m = make(map[string]int)
	splitStrings := strings.Split(s, " ")
	for _, v := range splitStrings {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}

```