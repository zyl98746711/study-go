package main

import (
	"fmt"
)

func main() {
	maxNum := 9
	if num := getNum(); num > maxNum {
		fmt.Printf("大于%d", maxNum)
	}

	str := "abc"
	fmt.Println(len(str))

	for i := 0; i < len(str); i++ {
		fmt.Printf("%c", str[i])
	}

	for i := 0; i < 25; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}

	for index, char := range str {
		fmt.Printf("位置：%d，值为：%c\n", index, char)
	}
	ss := "abc"
	AA(ss)
	BB(&ss)
	fmt.Println(&ss)
}

func getNum() int {
	return 10
}

func AA(strValue string) {
	aa := &strValue
	fmt.Println(aa)
}

func BB(strValue *string) {
	fmt.Println(strValue)
}
