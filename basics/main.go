package main

import (
	"fmt"
	"strconv"
)

func main(){
	var theHTTP string = "https://google.com"
	var i int //declare a variable with type integer
	i = 43
	fmt.Println(i)
	fmt.Println(theHTTP)

	//Type conversion
	var num int = 45
	fmt.Printf("%v, %T\n", num, num)

	var j = float32(num)  // Easily convert Integer to float and vice-versa
	fmt.Printf("%v, %T\n", j,j)

	var numr int = 23
	fmt.Printf("%v, %T\n", numr, numr)

	var tsg = strconv.Itoa(numr)
	fmt.Printf("%v, %T\n", tsg,tsg)
}
