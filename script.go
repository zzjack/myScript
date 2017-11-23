package main

import (
	"regexp"
	"fmt"
	"os/exec"
	"log"
	"bytes"
	"time"
	"math/big"
	"strings"
	"os"
	"encoding/csv"
	"bufio"
	"io"
)

func main(){
	//regexp_bracket()
	//unequal_bool()
	//split_left()
	//counter()
	//aa := call_s_val()
	//fmt.Println(aa)
	//measure_execution_time()
	//split_no_sep()
	//
	//var s splitedSlice
	//res := s.extractFromSplited(`"id`)
	//fmt.Println(res)
	//
	//split_comma()
	//parse()
	//chinese_index()
	index_sub()
}

func index_sub(){
	pattern,_ := regexp.Compile(`\d+`)
	res := pattern.FindAllStringSubmatch("3[0:23]",-1)
	fmt.Println(res)
}

func chinese_index(){
	a := `"你哈珀`
	w := []rune(a)
	fmt.Println(w,len(w))
}

func split_comma(){
	a := `531,c35f5d5d7d6241c1a47809ae7b0757ff,"330523198704150045",陈思,600,2,ZM201709023000000683900694155519,"{""zhima_credit_score_brief_get_response"":{""is_admittance"":""Y"",""code"":""10000"",""biz_no"":""ZM201709023000000683900694155519"",""msg"":""Success""},""sign"":""V9Vd+Q97RvlmNOZewI6hLzqQqFzwWHEWMYcQ7l+HZ/P1Dk0E4YYfoX6EoRkBisUi4tE2LDQteXlo01hmzx5GIqyqLKZ7CdDKq+t3O4lya4HIBxHDyPyTflU2JdCmkhWhu94Hk1huuqVjRUYUSx7KGW6BRORiWC3MrGU5horRANv8jtnvNWWHLEwY4+TUCX/a4Hvj9cXblk6ESxjkRAwgdo0/xNsqwjD/KKyP/+1ncLy6nXkTcVbxemKgGHBi7XgZYQ+YnNTj9OZ7sDP2tc4ZTSZrG4Z5XXkud3x6jUvccIZf1x2ImTbZQ2QGPYvpwZEk7CIIXBS9xIIs9w9299vnjw==""}",2017-09-02 03:08:29,2017-09-02 03:08:29`
	splited := strings.Split(a,",")
	fmt.Println("--->",len(splited))
	for w,i := range splited{
		fmt.Println(w,i)
	}
}

type splitedSlice struct {
	num int
	item string
	length int
}

func (s splitedSlice)extractFromSplited(f string)[]splitedSlice{
	var res []splitedSlice
	splited := strings.Split(f,`"`)
	for num,item := range splited{
		s.num = num
		s.item = item
		s.length = len(item)
		res = append(res,s)
	}
	return res
}


func split_no_sep(){
	s := `id"`
	w := strings.Split(s,`"`)
	fmt.Println(w,len(w))
}

func measure_execution_time(){
	start := time.Now()
	r := new(big.Int)
	fmt.Println(r.Binomial(1000,10))
	elapsed := time.Since(start)
	log.Printf("Binomial took %s\n",elapsed)
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

func parse(){
	f,_ := os.Open("/home/zzjack/Downloads/bsd_zhima_credit_score芝麻信用.csv")
	r := csv.NewReader(bufio.NewReader(f))
	var counter int
	for{
		counter++
		record,err := r.Read()
		if err == io.EOF{
			log.Println(err)
		}
		if counter == 10{
			break
		}
		fmt.Println(len(record),record)
	}
}

//func split_left(){
//	s := "新疆（乌鲁木齐）"
//	pro := strings.Split(s,"（")
//	fmt.Println(res1[0])
//	city := strings.Trim(res1[1],"）")
//	fmt.Println(res)
//}

func test_commnad(){
	cmd := exec.Command("tr","a-z","A-Z")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil{
		log.Fatalln(err)
	}
	fmt.Printf("in all caps:%q\n",out.String())
}

func regexp_scripts(){
	path := "bsd_zhima_credit_score芝麻信用.csv"
	pattern,err := regexp.Compile(`\w+`)
	if err != nil{
		fmt.Println(err)
	}
	res := pattern.FindAllString(path,-1)
	fmt.Println(res)
}

type name struct{
	sub int
	item string
}


func init_struct(){
	var nl []name
	for i := 0;i< 5;i++{
		n := name{}
		n.sub = i
		n.item = "hello"
		nl = append(nl,n)
	}
	fmt.Println(nl)
}

func modify_struct(){
	a := []name{name{1,"a"},name{2,"b"}}
	for s,i := range a{
		if i.sub == 1{
			//i.item = "change"
			a[s].item = "change"
		}
	}
	fmt.Println(a)
}

func test_regexp(){
	patternAddPath,_ := regexp.Compile(`[\.|/]*([^/]*)[/]*`)
	res := patternAddPath.FindAllStringSubmatch("./sdaDuanYan.csv",-1)
	fmt.Println(res)
}

func test_lookPath(){
	path,err := exec.LookPath("python")
	if err != nil{
		log.Fatal("installing fortune is in your future")
	}
	fmt.Printf("fortune is available at %s\n",path)
}