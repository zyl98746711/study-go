// 学习go的函数使用
package main

import fm "fmt"

var A = 5

func functionName() {
	fm.Println("It's a function")
}

func add(a int, b int) (int, int) {
	return b, a
}

func init() {
	fm.Println("init function")
}

func main() {
	functionName()
	fm.Println(A)
	fm.Println(add(1, 2))
	s := float64(A)
	fm.Println(s)
	fm.Println("Finished")
}
