package main

func main() {
	println(testFunctionParam(1, 2, functionAdd))
	println(testFunctionParam(3, 2, functionSub))

	println(testFunctionParam(2, 2, func(x, y int) int {
		return x * y
	}))
}

func functionAdd(a, b int) int {
	return a + b
}

func functionSub(a, b int) int {
	return a - b
}

func testFunctionParam(c, d int, f func(int, int) int) int {
	return f(c, d)
}
