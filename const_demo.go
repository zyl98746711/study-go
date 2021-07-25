package main

import fm "fmt"

const str string = "abc"

const STR string = "ABC"

const one, two, three = 1, 2, "3"

func main() {
	fm.Println(str)
	fm.Println(STR)
	fm.Println(one)
	fm.Println(len(three) + one)
	fm.Println(MALE)
	fm.Println(isMale(1))
	fm.Println(ABLE)
}

const (
	MALE   = 1
	FEMALE = 2
)

func isMale(sex int) bool {
	return MALE == sex
}

const (
	ENABLE = iota + 1
	ABLE
)
