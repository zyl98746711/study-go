package main

import (
	"fmt"
	"os"
	"runtime"
)

var name string = "tom"

func func2(a, b int) (int, int) {
	return a + b, a - b
}

func main() {
	fmt.Println("name is : " + name)
	fmt.Println(os.Getenv("HOME"))
	// 局部变量使用":="方式初始化  局部变量 var a string = "abc" 	等同于 a := "abc"
	path := os.Getenv("PATH")
	fmt.Println(path)

	goos := runtime.GOOS
	fmt.Printf("The Operating system is: %s\n", goos)
	fmt.Println(&goos)

	sprintf := fmt.Sprintf("The Operating system is: %s,sprintf", goos)
	fmt.Println(sprintf)

	path = "mac"
	fmt.Print(path)

	var path2 string
	fmt.Println(path2)

	path2 = "string2"
	fmt.Println(path2)

	_, aa := func2(2, 3)
	fmt.Println(aa)
}
