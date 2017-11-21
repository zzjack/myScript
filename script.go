package main

import (
	"regexp"
	"fmt"
)

func main(){
	//regexp_bracket()
	//unequal_bool()
	//split_left()
	//counter()
	aa := call_s_val()
	fmt.Println(aa)
}

type aa struct {
	s string
	b int
}

func call_s_val()[10][]aa{
	var a [10][]aa
	struct_val(a)
	return a
}

func struct_val(a [10][]aa){
	counter := 1
	for counter < 9{
		var c aa
		c.s = "s"
		c.b = counter
		a[counter] = append(a[counter],c)
		counter++
	}
}

func regexp_bracket(){
	pattern,_ := regexp.Compile(`\(`)
	res := pattern.FindAllStringSubmatch("(new Book)",-1)
	fmt.Println(res)
}

func unequal_bool(){
	var a bool
	fmt.Println(!a)
}

func counter(){
	var c int
	if c < 10{
		c++
		fmt.Println(c)
	}
}

//func split_left(){
//	s := "新疆（乌鲁木齐）"
//	pro := strings.Split(s,"（")
//	fmt.Println(res1[0])
//	city := strings.Trim(res1[1],"）")
//	fmt.Println(res)
//}