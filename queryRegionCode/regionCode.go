package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"sync"
)

var (
	dbssd               *gorm.DB
	dbhelpme            *gorm.DB
	conf                config
	patternLeftBracket  *regexp.Regexp
	patternRightBracket *regexp.Regexp
)

type config struct {
	FilePath  string              `json:"file_path"`
	FileType  string              `json:"file_type"`
	SplitTag  string              `json:"split_tag"`
	Databases map[string]Database `json:"databases"`
}

func init() {
	verifyArgs()
	confPath := os.Args[1]
	//confPath := "/home/gopath/src/myScript/queryRegionCode/queryRegionCodeConf.json"
	conf = loadConf(confPath)
	makeConn(confPath)
	left, _ := regexp.Compile(`（`)
	patternLeftBracket = left
	right, _ := regexp.Compile(`）`)
	patternRightBracket = right
}

func main() {
	provCt := conf.parseFile()
	allCode := conf.getAllCodes()
	value := conf.generateVal(provCt, allCode)
	conf.emptyTable()
	conf.insertVal(value)
	defer dbhelpme.Close()
	defer dbssd.Close()
}

func (c config) insertVal(value [10][]BsdHighDangerArea) {
	var wg sync.WaitGroup
	for _, val := range value {
		if len(val) == 0 {
			log.Fatalln("empty insert value")
		}
		wg.Add(1)
		go func(val []BsdHighDangerArea){
			for _, v := range val {
				log.Println("insert table", v)
				dbhelpme.Create(&v)
			//	note:if there is dbhelpme.commit,the insert will fail.
			}
			defer wg.Add(-1)
		}(val)
	}
	wg.Wait()
}

func (c config) emptyTable() {
	ssd := BsdHighDangerArea{}
	dbhelpme.Delete(&ssd)
	log.Println("empty table")
}

//the thought of switch was learned from Learning C Program
func (c config) generateVal(provCt map[string][]string, allCode []GeoCode) [10][]BsdHighDangerArea {
	var bowl [10][]BsdHighDangerArea
	counter := 0
	for _, geo := range allCode {
		for pro, city := range provCt {
			if len(city) == 0 {
				model := fmt.Sprintf("%s", pro)
				pattern, err := regexp.Compile(model)
				checkErr(err, "addCities province regexp compile failed")
				if pattern.MatchString(geo.ProvinceName) {
					var bsd BsdHighDangerArea
					bsd.ProvinceCode = geo.ProvinceCode
					bsd.ProvinceName = geo.ProvinceName
					bsd.CityFullCode = geo.CityFullCode
					bsd.CityName = geo.CityName
					bsd.DistrictCode = geo.DistrictCode
					bsd.DistrictName = geo.DistrictName
					bsd.Type = "1"
					bowl[counter] = append(bowl[counter], bsd)
				}
			} else {
				for _, c := range city {
					model := fmt.Sprintf("%s", c)
					pattern, err := regexp.Compile(model)
					checkErr(err, "addCities city regexp compile failed")
					if pattern.MatchString(geo.CityName) {
						var bsd BsdHighDangerArea
						bsd.ProvinceCode = geo.ProvinceCode
						bsd.ProvinceName = geo.ProvinceName
						bsd.CityFullCode = geo.CityFullCode
						bsd.CityName = geo.CityName
						bsd.DistrictCode = geo.DistrictCode
						bsd.DistrictName = geo.DistrictName
						bsd.Type = "2"
						bowl[counter] = append(bowl[counter], bsd)
					}
				}
			}
			counter++
			if counter == 10 {
				counter = 0
			}
		}
	}
	log.Println("generate value:", bowl)
	return bowl
}

func (c config) getAllCodes() []GeoCode {
	g := []GeoCode{}
	dbssd.Find(&g)
	return g
}

func verifyArgs() {
	if len(os.Args) < 1 {
		log.Fatalln("please enter config path")
	}
}

func (c config) parseFile() map[string][]string {
	var status bool
	var bowl []string
	var who string
	res := make(map[string][]string)
	readed, err := ioutil.ReadFile(c.FilePath)
	checkErr(err, "getRiskWord")
	splited := strings.Split(string(readed), c.SplitTag)
	for _, s := range splited {
		left := patternLeftBracket.MatchString(s)
		right := patternRightBracket.MatchString(s)
		if left == false && right == false && status == false {
			res[s] = nil
		} else if left == true && right == false && status == false {
			index := patternLeftBracket.FindStringIndex(s)
			sub := index[0]
			sub1 := index[1]
			who = string(s[0:sub])
			bowl = append(bowl, string(s[sub1:]))
			status = true
		} else if status == true && left == false && right == false {
			bowl = append(bowl, s)
		} else if status == true && right == true && left == false {
			index := patternRightBracket.FindStringIndex(s)
			sub := index[0]
			bowl = append(bowl, string(s[0:sub]))
			res[who] = bowl
			status = false
			who = ""
			bowl = []string{}
		} else if left == true && right == true && status == false {
			splited := strings.Split(s, "（")
			provin := splited[0]
			city := strings.Trim(splited[1], "）")
			res[provin] = []string{city}
		} else {
			log.Fatalln("Unexpected situation when parsing file", s)
		}
	}
	log.Println(res)
	return res
}

func loadConf(confPath string) config {
	var conf config
	readed, err := ioutil.ReadFile(confPath)
	checkErr(err, "loadConf")
	err = json.Unmarshal(readed, &conf)
	checkErr(err, "unmarshal conf")
	if len(conf.Databases) == 0 {
		log.Fatalln("unexpected readed conf", conf)
	}
	return conf
}

func checkErr(err error, tag string) {
	if err != nil {
		log.Fatal(tag, err)
	}
}

func makeConn(confPath string) {
	connModel := "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"
	ssdInfo := conf.Databases["ssd"]
	helpmeInfo := conf.Databases["helpme"]
	connSsd := fmt.Sprintf(connModel, ssdInfo.User, ssdInfo.Password, ssdInfo.Host, ssdInfo.Port, ssdInfo.Name)
	connHelpme := fmt.Sprintf(connModel, helpmeInfo.User, helpmeInfo.Password, helpmeInfo.Host, helpmeInfo.Port, helpmeInfo.Name)
	dbSsd, err := gorm.Open("mysql", connSsd)
	checkErr(err, connSsd)
	dbssd = dbSsd
	dbHelpme, err := gorm.Open("mysql", connHelpme)
	checkErr(err, connHelpme)
	dbhelpme = dbHelpme
}
