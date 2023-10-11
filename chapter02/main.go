package main

import (
	"fmt"
	"math"
)

func main() {
	integerSample()
	floatSample()
	constSample()
}

// 整数型
func integerSample() {
	fmt.Printf("================= integer ================\n")
	// =========== 型チェック ============
	n := 100
	var i int = 100
	var i8 int8 = 100
	var i16 int16 = 100
	var i32 int32 = 100
	var i64 int64 = 100
	var ui8 uint8 = 100
	var ui16 uint16 = 100
	var ui32 uint32 = 100
	var ui64 uint64 = 100
	var wert = i64

	fmt.Printf("n ---- type: %T\n", n)
	fmt.Printf("i ---- type: %T\n", i)
	fmt.Printf("i8 --- type: %T\n", i8)
	fmt.Printf("type: %T\n", i16)
	fmt.Printf("type: %T\n", i32)
	fmt.Printf("type: %T\n", i64)
	fmt.Printf("type: %T\n", ui8)
	fmt.Printf("type: %T\n", ui16)
	fmt.Printf("type: %T\n", ui32)
	fmt.Printf("type: %T\n", ui64)
	fmt.Printf("type: %T\n", wert)

	// =========== ゼロ値チェック ============
	var n_z int
	var s_z string
	var f32_z float32
	var f64_z float64
	var b_z bool
	fmt.Printf("int_zero     --> %d\n", n_z)
	fmt.Printf("string_zero  --> %s\n", s_z)
	fmt.Printf("float32_zero --> %f\n", f32_z)
	fmt.Printf("float64_zero --> %f\n", f64_z)
	fmt.Printf("bool_zero    --> %t\n", b_z)

	// =========== 演算 ============
	var x int = 10
	var y byte = 10
	// 型が異なるため演算できない
	// sum := x + y
	// sum := x - y
	// sum := x / y
	// sum := x % y

	// キャストして同じ方にすれば演算できる
	sum := byte(x) + y
	fmt.Printf("cast: %d\n", sum)
	fmt.Println()
}

// 浮動小数点型
func floatSample() {
	fmt.Printf("================= float ================\n")
	// ゼロ値
	var f32_z float32
	var f64_z float64
	fmt.Printf("float32_zero: %f\n", f32_z)
	fmt.Printf("float64_zero: %f\n", f64_z)

	// 範囲
	fmt.Printf("float32 max: %f\n", math.MaxFloat32)
	fmt.Printf("float32 min: %f\n", math.SmallestNonzeroFloat32)
	fmt.Printf("flaot64 max: %f\n", math.MaxFloat64)
	fmt.Printf("float64 min: %f\n", math.SmallestNonzeroFloat64)

	// 型の異なる演算
	var f32 float32 = 10
	var f64 float64 = 10
	// sum := f32 / f64 <- 型が異なるためエラー
	sum := f32 / float32(f64) // <- キャストすればok
	fmt.Printf("sum: %f\n", sum)

	// リテラルと明示された型との演算
	// n := 10
	fizz := 10 / float64(3)
	fmt.Printf("fizz: %f\n", math.Ceil(fizz)) // <- 切り上げ
	fmt.Printf("fizz: %f\n", fizz)            // <- 切り捨て

	fmt.Println()
}

// 定数
func constSample() {
	fmt.Printf("================= const ================\n")
	const x = 10
	const x_i int = 10
	const y = 10.1
	const y_f float64 = 10.1

	sum_i := x + x_i
	sum_f := y + y_f
	fmt.Printf("sum_i: %d\n", sum_i)
	fmt.Printf("sum_f: %f\n", sum_f)
}
