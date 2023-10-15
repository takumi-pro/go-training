package main

import "fmt"

func main() {
	arraySample()
	sliceSample()
	makeSample()
	sliceExpression()
}

func arraySample() {
	fmt.Println("============== array ===============")
	// 配列宣言
	var x1 [3]int
	var x2 = [3]int{10, 20, 30}
	x3 := [4]int{10, 20, 30, 40}
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

	fmt.Printf("y2 --> %d\n", y2)
	fmt.Printf("y3 --> %d\n", y3)

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

	// copyによるメモリを共有しないコピー

	// 文字列のスライス
}
