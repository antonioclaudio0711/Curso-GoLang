package main

import "fmt"

func main() {
	var a int     //default == 0
	var b float64 //default == 0
	var c bool    //default == false
	var d string  //default == ""
	var e *int    //default == <nil> --> <null>

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
}
