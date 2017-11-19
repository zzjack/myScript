package main

import (
	"encoding/json"
	"fmt"
	"github.com/ThomasRooney/gexpect"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"reflect"
)

//It went through in Ubuntu 17.04

//2017/11/17
//add new feature,support different port;


type targetArgu struct {
	UserName string
	Passwd   string
	Host     string
    Port string 
    SendPath string
}

type conf struct {
	UserConf map[string]UserConfig `json:"user_conf"`
}

type UserConfig struct {
	UserName string `json:"user_name"`
	Passwd   string `json:"passwd"`
	Host     string `json:"host"`
	SendPath string `json:"send_path"`
    Port string `json:"port"`
}

var configPath = "/home/zzjack/tools/config.json"

func main() {
	//confirm an argument existing
	if len(os.Args) < 3 {
		log.Fatal("please enter packaged file name")
	}
	input := os.Args[1]
	which := os.Args[2]
	//read config from config.json
	conf := readConfig(configPath, which)
	var t targetArgu
	t.UserName = conf.UserName
	t.Passwd = conf.Passwd
	t.Host = conf.Host
	t.SendPath = conf.SendPath
    t.Port = conf.Port
	log.Println("read config of send path ==>", t.SendPath)
	filePath, fileName := t.distinctPathFile(input)
	tarName := fmt.Sprintf("%s.tar.gz", fileName)
	t.compressFile(tarName, input)
	tarFilePath := fmt.Sprintf("%s%s", filePath, tarName)
	t.sendFile(tarFilePath)
}

func readConfig(file, which string) UserConfig {
	var c conf
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(content, &c)
	if err != nil {
		log.Fatal(err)
	}
	w := c.UserConf[which]
	//insure config not empty string
	insureNotEmpty(w)
	return w
}

func insureNotEmpty(conf UserConfig){
	numField := reflect.ValueOf(conf).NumField()
	for i := 0;i<numField;i++{
		lenOfVal := reflect.ValueOf(conf).Field(i).Len()
		if lenOfVal == 0{
			name := reflect.TypeOf(conf).FieldByIndex([]int{i}).Name
			log.Panic("In config",name," is empty")
		}
	}
}

func (t targetArgu) compressFile(tarName, input string) {
	cmd := exec.Command("tar", "cvf", tarName, input)
	_, err2 := cmd.CombinedOutput()
	if err2 != nil {
		log.Fatal(err2)
	}
	printCommand(cmd)
}

func (t targetArgu) sendFile(tarFilePath string) {
	//send to another machine
	spawnCommand := fmt.Sprintf("scp -P %s %s %s@%s:%s",t.Port,tarFilePath, t.UserName, t.Host, t.SendPath)
	log.Println("Executing Command ==>", spawnCommand)
	child, err := gexpect.Spawn(spawnCommand)
	if err != nil {
		log.Fatal(err)
	}
	searchString := fmt.Sprintf("%s@%s's password:", t.UserName, t.Host)
	log.Println("expected string ==>", searchString)
	child.Expect(searchString)
	passwdEnter := fmt.Sprintf("%s\n", t.Passwd)
	child.SendLine(passwdEnter)
	child.Interact()
	child.Close()
}

func (t targetArgu) distinctPathFile(input string) (string, string) {
	var filePath string
	var fileName string
	var pathFileCln string
	//debug: the pathFile ends with "/" or not,the result is deffence
	if string(input[len(input)-1]) == "/" {
		pathFileCln = string(input[0 : len(input)-1])
	} else {
		pathFileCln = input
	}
	splitBySlant := strings.Split(pathFileCln, "/")
	switch len(splitBySlant) {
	case 1:
		filePath = ""
		fileName = splitBySlant[0]
	default:
		lastIndex := len(splitBySlant) - 1
		filePath = fmt.Sprintf("%s/", strings.Join(splitBySlant[0:lastIndex], "/"))
		fileName = splitBySlant[lastIndex]
	}
	return filePath, fileName
}

func printCommand(cmd *exec.Cmd) {
	log.Printf("Executing Command ==> %s\n", strings.Join(cmd.Args, " "))
}
