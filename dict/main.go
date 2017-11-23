package main

import "log"

var q qihu
var input string

func init(){
	q.breakSign = "+++++++++++++++++++++++++++++++"
	q.url = "https://sug.so.360.cn/suggest?callback=suggest_so&encodein=utf-8" +
		"&encodeout=utf-8&format=json&fields=word&word="
}

func main(){
	word := q.clnWord()
	if patternLetters.MatchString(word){
		res := chineseToEngMain(word,q.breakSign)
		if !res{
			queryPinyinOrChinese(word)
		}
	} else {
		chineseToEngMain(word,q.breakSign)
	}
}

func queryPinyinOrChinese(word string){
	url := q.makeLink(word)
	resp := q.clnResp(url)
	disRes := q.display(resp)
	q.translateToEng(disRes,input)
}

func checkError(err error,tag string){
	if err != nil{
		log.Fatalln(tag,err)
	}
}


