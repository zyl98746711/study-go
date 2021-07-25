package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var aScope string

func main() {
	aScope = "G"
	println(aScope)
	f1()

	fmt.Println(math.MaxUint16)

	times := int64(time.Now().Nanosecond())
	rand.Seed(times)

	for i := 0; i < 10; i++ {

		n := rand.Intn(10)
		if n == 9 {
			fmt.Println("finished")
			break
		}
		fmt.Println(n)
	}

	reader := strings.NewReader(aScope)
	readRune, size, err := reader.ReadRune()
	fmt.Println(readRune)
	fmt.Println(size)
	fmt.Println(err)

	strInt := "1"
	atoi, _ := strconv.Atoi(strInt)
	fmt.Println(atoi)

	fmt.Println(time.Now())
	fmt.Println(time.Now().Clock())
	fmt.Println(time.Now().Add(1000 * 60 * 1000))
}

func f1() {
	aScope := "O"
	println(aScope)
	f2()
}
func f2() {
	println(aScope)
}
