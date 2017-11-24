package main

import (
	"encoding/json"
	"fmt"
	"github.com/ThomasRooney/gexpect"
	"io/ioutil"
	"log"
	"os"
)

var (
	configPath = "/home/zzjack/tools/config.json"
	commandSet = "set timeout 10"
	commandSpawn = "ssh -p %s %s@%s"
	commandExpect = `%s@%s`
	commandSend = `%s`
	commandInteract = "interact"
)

type command struct {
	spawnSettedPort string
	spawn string
	expect string
	send string
	interact string
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


func main(){
	verifyArgs(2)
	machineName := getMachineName(os.Args)
	log.Println("machine name :",machineName)
	relatedConf := getConfBy(machineName)
	log.Println("related config",relatedConf)
	completed := makeCompletedCommand(relatedConf)
	execCommand(completed)
}

func execCommand(completed command){
	_,err := gexpect.Command(commandSet)
	child,err := gexpect.Spawn(completed.spawn)
	checkErr(err,"gexpect.Spawn")
	_,err = child.ExpectRegex(completed.expect)
	checkErr(err,"child.ExpectRegex")
	err = child.SendLine(completed.send)
	checkErr(err,"chile.SendLine")
	child.Interact()
}

func makeCompletedCommand(relatedConf UserConfig)command{
	var c command
	c.spawn = fmt.Sprintf(commandSpawn,relatedConf.Port,relatedConf.UserName,relatedConf.Host)
	c.expect =  fmt.Sprintf(commandExpect,relatedConf.UserName,relatedConf.Host)
	c.send = fmt.Sprintf(commandSend,relatedConf.Passwd)
	c.interact = commandInteract
	return c
}

func getConfBy(machineName string)UserConfig{
	path := getConfPath()
	allConf := readConf(path)
	return allConf.UserConf[machineName]
}

func readConf(path string)conf{
	var c conf
	readed,err := ioutil.ReadFile(path)
	checkErr(err,"read conf")
	err = json.Unmarshal(readed,&c)
	checkErr(err,"unmarshal conf")
	return c
}

func getConfPath()string{
	var path string
	if len(os.Args) == 3{
		path = os.Args[2]
	} else {
		path = configPath
	}
	return path
}

func getMachineName(args []string)string{
	return args[1]
}

func verifyArgs(num int){
	if len(os.Args) < num{
		log.Fatalln("arguments not legal")
	}
}

func checkErr(err error,tag string){
	t := fmt.Sprintf("Unexpected when %s",tag)
	if err != nil{
		log.Fatalln(t,err)
	}
}