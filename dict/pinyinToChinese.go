package main

import (
	"fmt"
	"os"
	"net/http"
	"log"
	"io/ioutil"
	"encoding/json"
	"regexp"
	"strconv"
)

//version2.0:turn pinyin to english directly
//version3.0:use one command to call
//2017.11.23.
//there is judging by sequence numbers are equal
//add new feature: support enter sequence number/sequence number slice/character

var patternLetters *regexp.Regexp
var patternDigit *regexp.Regexp
var patternDigitSlice *regexp.Regexp
var patternDigitAll *regexp.Regexp

type qihu struct {
	url string
	breakSign string
	qihuResp
}

type qihuResp struct {
	Query string `json:"query"`
	Result []ResultArr
}

type ResultArr struct {
	Word string `json:"word"`
}

type displayRes [][2]string

func init(){
	p,err := regexp.Compile(`[A-Z|a-z]+`)
	checkError(err,"compile letter error")
	patternLetters = p
	digit,err := regexp.Compile(`^\d+$`)
	checkError(err,"compile digit error")
	patternDigit = digit
	dSlice,err := regexp.Compile(`^\d+\[\d*?:\d\]*?`)
	checkError(err,"compile digit slice error")
	patternDigitSlice = dSlice
	all,err := regexp.Compile(`\d+`)
	checkError(err,"all regexp")
	patternDigitAll = all
}



func (q qihu) translateToEng(disRes displayRes,input string){
	fmt.Printf("\nPlease enter the num of word wanted to look up: ")
	fmt.Scan(&input)
	judged := q.judgeInputType(disRes,input)
	fmt.Println("Be trying To translate ",judged," into English...\n")
	chineseToEngMain(judged,q.breakSign)
	}

func (q qihu)judgeInputType(disRes displayRes,input string)string{
	var judged string
	if patternDigit.MatchString(input){
		for _,item := range disRes{
			if input == item[0]{
				judged = item[1]
			}
		}
	} else if patternDigitSlice.MatchString(input){
		index := patternDigitAll.FindAllStringSubmatch(input,-1)
		if len(index) != 3{
			log.Fatalln("can't match 3 digits","index:",index,"; input:",input)
		}
		parsed := parseThree(index)
		for _,item := range disRes{
			if index[0][0] == item[0]{
				judged = func()string{
					var s string
					r := []rune(item[1])[parsed[1]:parsed[2]]
					for _,i := range r{
						s += string(i)
					}
					return s}()
			}
		}
		} else {
			judged = input
	}
	return judged
}


func parseThree(index [][]string)[3]int{
	var parsed [3]int
	for seq,i := range index{
		p,err := strconv.ParseInt(i[0],10,64)
		checkError(err,"parseThreee")
		parsed[seq] = int(p)
	}
	return parsed
}

func (q qihu)display(resp qihuResp)displayRes{
	var resArr displayRes
	fmt.Println("\nWhich one do you need:")
	fmt.Println(q.breakSign)
	for seq,res := range resp.Result{
		fmt.Println(seq,res.Word)
		seqStr := strconv.FormatInt(int64(seq),10)
		resArr = append(resArr,[2]string{seqStr,res.Word})
		if seq+1 == len(resp.Result){
			fmt.Println(q.breakSign)
		}
	}
	return resArr
}

func (q qihu)clnResp(url string)qihuResp{
	var qi qihuResp
	resp,err := http.Get(url)
	q.fatalErr(err,url)
	readRes,err := ioutil.ReadAll(resp.Body)
	q.fatalErr(err,"ioutil.ReadAll")
	pattern,err := regexp.Compile("{.*}")
	q.fatalErr(err,"compile pattern")
	regexpRes := pattern.FindAllString(string(readRes),-1)
	var unmarshalObj []byte
	if len(regexpRes) > 0{
		unmarshalObj = []byte(regexpRes[0])
	}
	err = json.Unmarshal(unmarshalObj,&qi)
	q.fatalErr(err,string(readRes))
	return qi
}

func (q qihu)fatalErr(err error,msg interface{}){
	if err != nil{
		log.Println("ERROR MSG",msg)
		log.Fatal(err)
	}
}

func (q qihu)makeLink(word string)string{
	return fmt.Sprintf("%s%s",q.url,word)
}

func (q qihu)clnWord()string{
	args := os.Args
	if len(args) == 1{
		log.Fatalln("please enter the word wanted to look up")
	}
	return args[1]
}











