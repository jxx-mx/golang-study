package main

import (
	"firstGo/apkg"
	"fmt"
)

func swap(x, y string) (string, string) {
	return y, x
}

type computeAcion func(int, int) int

func compute(x, y int, action computeAcion) int {
	return action(x, y)
}

func add(a, b int) int {
	return a + b
}

func div(a, b int) int {
	return a / b
}

func mul(a, b int) int {
	return a * b
}

var i, j int = 1, 2
// var c, pyon, java bool

func foo() {
	m := 1
	
	p := &m
	
	*p = 2
	fmt.Println("fdads", m)
}

type Vertex struct {
	X int
	Y int
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func bar() {
	pos, neg := adder(), adder()
	for i := 0; i< 10; i++ {
		fmt.Println(pos(i),neg(-2*1))
	}
}

func main() {
	fmt.Println(Vertex{X: 1})
	fmt.Println("hello world")
	apkg.Foo()
	fmt.Println(apkg.GetNumber())
	fmt.Println(add(5, 2))
	fmt.Println(swap("as", "df"))
	fmt.Println(compute(3, 2 ,add))
	fmt.Println(compute(3, 2 ,div))
	fmt.Println(compute(3, 2 ,mul))
	foo()
	bar()
	
	var c, pyon, java = true, true, "no!"
	k := 3
	fmt.Println(k, i, j)
	fmt.Println(c, pyon, java)
}