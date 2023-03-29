package main

import (
	"fmt"
	"strconv"
)

var (g int = 33)

func main(){
	i := 999999999999999999
	var s string = "JUSE";
	fmt.Println(i)
	fmt.Println(g)
	fmt.Println(s)
	fmt.Println(string(i)) // that's bad
	fmt.Println(strconv.Itoa(i)) // this is good
	{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{{}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}}} // nice
}