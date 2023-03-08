package main

import "fmt"

func main() {
	var i int = 21
	var j bool = true
	var a = '\u042F'
	var c int = 15
	var k float64 = 123.456

	fmt.Println(i)
	fmt.Printf("%T \n", i)
	fmt.Printf("%% \n")
	fmt.Println(j)
	fmt.Println(!j)

	fmt.Printf("%c \n", a)
	fmt.Printf("%d \n", i)
	fmt.Printf("%o \n", i)
	fmt.Printf("%x \n", c)
	fmt.Printf("%X \n", c)
	fmt.Printf("%U \n", a)

	fmt.Printf("%g \n", k)
	fmt.Printf("%f \n", k)
	fmt.Printf("%E \n", k)

}
