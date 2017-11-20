package main

import (
	"regexp"
	"fmt"
)

func main(){
	regexp_bracket()
}

func regexp_bracket(){
	pattern,_ := regexp.Compile(`\(`)
	res := pattern.FindAllStringSubmatch("(new Book)",-1)
	fmt.Println(res)
}