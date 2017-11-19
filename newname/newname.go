package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"os/exec"
)

var minuim int
var patternStick *regexp.Regexp
var patternFile *regexp.Regexp

type nametuple struct {
	sub  int
	item string
}

func init() {
	patternStick, _ = regexp.Compile(`\W`)
	patternFile, _ = regexp.Compile(`[A-Z|a-z|\s|\d]`)
	minuim = 1
}

func main() {
	fileName := extractFrom(os.Args)
	illegaled, fileChar := pickFrom(fileName)
	joined := joinillegaled(illegaled)
	spell := turnToSpell(joined)
	splited := splitSpell(spell)
	checked := checksplitedRes(splited)
	replaced := replaceIllegaled(illegaled, checked)
	legalFileName := makelegalFilename(replaced, fileChar)
	path,newFileName := makeNew(strings.Join(legalFileName, ""))
	renameFile(newFileName,path,fileName)
}

func renameFile(newFileName,path,filName string){
	newTarget := path + newFileName
	log.Println("fileName: ",filName)
	log.Println("newTarget: ",newTarget)
	c := exec.Command("mv",filName,newTarget)
	c.Run()
	if c.Stderr != nil{
		log.Fatalln(c.Stderr)
	}
	}

//debug:the data is not compelely right that third interface returned.As shown:
//[./ duan yan]
func checksplitedRes(splited []string) []string {
	var loopRes [][]string
	for _, s := range splited {
		loopRes = append(loopRes, isSticked(s))
	}
	return _flattenLoopRes(loopRes)
}

func _flattenLoopRes(lr [][]string) []string {
	var bowl []string
	for _, slice := range lr {
		for _, rice := range slice {
			bowl = append(bowl, rice)
		}
	}
	return bowl
}

func isSticked(s string) []string {
	var res []string
	if patternStick.MatchString(s) && len(s) > 1 {
		for _, i := range s {
			res = append(res, string(i))
		}
	} else {
		res = append(res, s)
	}
	return res
}

func makelegalFilename(replaced []nametuple, filechar []*nametuple) []string {
	var bowl []string
	for _, r := range replaced {
		for _, f := range filechar {
			if r.sub == f.sub {
				f.item = r.item
				break
			}
		}
	}
	for _, c := range filechar {
		bowl = append(bowl, c.item)
	}
	return bowl
}

func replaceIllegaled(illegaled []nametuple, splited []string) []nametuple {
	if len(illegaled) != len(splited) {
		log.Fatalln("interface data returned illegal\n", "illegaled:", illegaled, "\nsplited", splited)
	}
	for s, v := range splited {
		illegaled[s].item = v
	}
	return illegaled
}

func splitSpell(spell string) []string {
	return strings.Split(spell, " ")
}

func joinillegaled(illegaled []nametuple) string {
	var bowl []string
	for _, v := range illegaled {
		bowl = append(bowl, v.item)
	}
	return strings.Join(bowl, "")
}

func extractFrom(input []string) string {
	var s string
	if len(input) == minuim {
		log.Fatal("please add the file name that you want to transform")
	} else {
		for _,i := range input[1:]{
			s += i
		}
	}
	return s
}

func makeNew(legal string)(string,string){
	var (
		path string
		fileName string
	)
	splited := strings.Split(legal,"/")
	if len(splited) == 0{
		path = ""
		fileName = legal
	} else {
		lastSub := len(splited) - 1
		for _,p := range splited[0:lastSub]{
			path += p + "/"
		}
		fileName = splited[lastSub]
	}
	return path,fileName
}

func pickFrom(fileName string) ([]nametuple, []*nametuple) {
	var legal []*nametuple
	var illegal []nametuple
	for i, c := range fileName {
		if !patternFile.MatchString(string(c)) {
			n := nametuple{}
			n.sub = i
			n.item = string(c)
			illegal = append(illegal, n)
		}
		if c == ' ' {
			c = '_'
		}
		l := nametuple{}
		l.sub = i
		l.item = string(c)
		legal = append(legal, &l)
	}
	return illegal, legal
}

func turnToSpell(s string) string {
	url := "https://www.dute.me/tools/pinyin/?action=do"
	data := map[string][]string{
		"hanzi":                  []string{s},
		"with_separate":          []string{"1"},
		"first_letter_uppercase": []string{"1"},
	}
	resp, err := http.PostForm(url, data)
	if err != nil {
		log.Fatal(err)
	}
	readed, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(readed)
}
