package main

import (
	"fmt"
)

func main() {
	arraySample()
	sliceSample()
	makeSample()
	sliceExpression()
	mapSample()
	structSample()
}

func arraySample() {
	fmt.Println("============== array ===============")
	// 配列宣言
	var x1 [3]int                 // varを用いた配列の初期化
	var x2 = [3]int{10, 20, 30}   // varを用いた宣言
	x3 := [4]int{10, 20, 30, 40}  // :=を用いた宣言
	x4 := [...]int{1, 2, 3, 4, 5} // 要素数の省略

	fmt.Printf("x1 --> %d\n", x1)
	fmt.Printf("x2 --> %d\n", x2)
	fmt.Printf("x3 --> %d\n", x3)
	fmt.Printf("x4 --> %d\n", x4)
	fmt.Printf("x4[0] --> %d\n", x4[0])

	// 配列の型
	// 要素数によって型が変わる
	fmt.Printf("x1 type --> %T\n", x1) // x1 type --> [3]int
	fmt.Printf("x2 type --> %T\n", x2) // x2 type --> [3]int
	fmt.Printf("x3 type --> %T\n", x3) // x3 type --> [4]int
	fmt.Printf("x4 type --> %T\n", x4) // x4 type --> [5]int

	// 配列のポインタを確認する
	// 代入するとコピーが作成される
	x5 := [4]int{1, 2, 3, 4}
	x6 := x5

	fmt.Printf("x5 pointer --> %p %p\n", &x5, &x5[0])
	fmt.Printf("x6 pointer --> %p %p\n", &x6, &x6[0])

	fmt.Println()
}

func sliceSample() {
	fmt.Println("============== slice ===============")

	// 要素数を指定しないとスライスになる
	var x1 []int
	var x2 = []int{1, 2, 3, 4, 5}
	var x3 = []int{1, 2, 3, 4, 5}
	var x4 = []int{}

	// スライスのゼロ値はnil
	// nilに型はない
	if x1 == nil {
		fmt.Printf("x1 is nil\n")
	}

	// 空のスライスはnilではない
	if x4 != nil {
		fmt.Printf("x4 is not nil\n")
	}

	fmt.Printf("x1 --> %d\n", x1)
	fmt.Printf("x2 --> %d\n", x2)
	fmt.Printf("x2 type --> %T\n", x2)
	fmt.Printf("x3 type --> %T\n", x3)

	// 多次元スライスの定義
	var y [][]int
	var y1 = [][]int{{1, 2, 3}, {4, 5}, {6, 7, 8, 9}}

	fmt.Printf("y --> %d\n", y)
	fmt.Printf("y1 --> %d\n", y1)

	// スライス同士は比較できない
	// nilは比較可能
	// if y == y1 <-- invalid operation: y == y1 (slice can only be compared to nil)

	// スライスの代入
	// 変数にスライスを代入して新しいスライスを定義するとメモリを共有する
	y2 := []int{1, 2, 3, 4, 5}
	y3 := y2

	fmt.Printf("y2 --> Pointer: %p, Refer: %p, Value: %d\n", &y2, &y2[0], y2)
	fmt.Printf("y3 --> Pointer: %p, Refer: %p, Value: %d\n", &y3, &y3[0], y3)

	// y2の要素を変更するとy3をも変化する
	y2[2] = 100
	y2 = append(y2, 200)

	// 2番目の要素は変更されているがappendによる要素追加は行われない
	fmt.Printf("y2 --> %d\n", y2)
	fmt.Printf("y3 --> %d\n", y3)

	fmt.Println()
}

func appendSample() {
	var x []int
	x1 := []int{10, 20, 30}

	x = append(x, 1, 2, 3)
	x1 = append(x1, x...) // <-- 展開してスライスの全要素を追加

	fmt.Printf("x --> %d\n", x)
	fmt.Printf("x1 --> %d\n", x1)
}

func makeSample() {
	// makeを使うと長さとキャパシティを指定してスライスを作成できる
	x := make([]int, 3)

	fmt.Printf("x --> %d\n", x)
	fmt.Printf("x length --> %d\n", len(x))
	fmt.Printf("x capacity --> %d\n", cap(x))

	x = append(x, 1, 2, 3, 4)

	fmt.Printf("x --> %d\n", x)
	fmt.Printf("x length --> %d\n", len(x))
	fmt.Printf("x capacity --> %d\n", cap(x))

	// キャパシティを指定する
	x1 := make([]int, 5, 10) // <-- 長さ:5, キャパシティ:10

	fmt.Printf("x1 --> %d\n", x1)
	fmt.Printf("x1 length --> %d\n", len(x1))
	fmt.Printf("x1 capacity --> %d\n", cap(x1))
}

// スライス式
func sliceExpression() {
	fmt.Println("================== slice expression =================")
	/** スライス式の宣言
	* x[n:m]
	* x --> スライスor配列
	* n --> 開始オフセット　未指定の場合は0
	* m --> 終了オフセット 未指定の場合は最後の値
	**/
	x := []int{1, 2, 3, 4, 5}
	x1 := x[:1]
	x2 := x[1:]
	x3 := x[1:2]
	x4 := x[2:5]

	fmt.Printf("x --> %d\n", x)
	fmt.Printf("x[:1] --> %d\n", x1)
	fmt.Printf("x[1:] --> %d\n", x2)
	fmt.Printf("x[1:2] --> %d\n", x3)
	fmt.Printf("x[2:5] --> %d\n", x4)

	// slice expressioで作成したスライスはメモリを共有している
	x5 := x[:2]
	x5 = append(x5, 100)

	fmt.Printf("x --> %d\n", x)
	fmt.Printf("x5 --> %d\n", x5)

	// 配列からスライス
	arr := [5]int{1, 2, 3, 4, 5}
	x6 := arr[:]
	x7 := x6

	fmt.Printf("arr --> Pointer: %p, Refer: %p, Value: %d\n", &arr, &arr[0], arr)
	fmt.Printf("x6 --> Pointer: %p, Refer: %p, Value: %d\n", &x6, &x6[0], x6)
	fmt.Printf("x7 --> Pointer: %p, Refer: %p, Value: %d\n", &x7, &x7[0], x7)

	// copyによるメモリを共有しないコピー
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 5)
	copy(slice2, slice1)

	slice1[0] = 100
	slice2[1] = 200

	fmt.Printf("slice1 --> Pointer: %p, Refer: %p, Value: %d\n", &slice1, &slice1[0], slice1)
	fmt.Printf("slice2 --> Pointer: %p, Refer: %p, Value: %d\n", &slice2, &slice2[0], slice2)

	// 文字列のスライス
}

func mapSample() {
	// 宣言
	var x1 map[string]int  // ゼロ値はnil・書き込み不可
	x2 := map[string]int{} // 書き込み可能

	fmt.Printf("x1 --> %v\n", x1)
	fmt.Printf("x2 --> %v\n", x2)
}

func structSample() {
	type human struct {
		name string
		age  int
		hoby string
	}
	var taro human // 構造体のゼロ値

	x1 := human{"taro", 10, "reading a book"}               // 構造体リテラル
	x2 := human{name: "jiro", age: 20, hoby: "programming"} // : で区切る・順番が自由
	var x3 struct {                                         // 無名構造体
		name string
		age  int
		pet  string
	}
	x3.age = 10
	x3.name = "saburo"

	fmt.Printf("human --> %v\n", taro)
	fmt.Printf("x1 --> %v\n", x1)
	fmt.Printf("x1 type --> %T\n", x1)
	fmt.Printf("x2 --> %v\n", x2)
	fmt.Printf("x3 --> %v\n", x3)
}
