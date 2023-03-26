package main

import (
	"fmt"
	"os"

	"github.com/baimu7/homework/linklist"
)

const(
	CMD_MAX_LEN=128
	DESC_LEN=1024
	CMD_NUM=10
)
func main(){
	for {
		cmd:=make([]byte,CMD_MAX_LEN);
		fmt.Print(">>> Input a command: ")
		fmt.Scanln(&cmd)
		p:=linklist.Findcmd(head,string(cmd))
		if p==nil{
			fmt.Println("this is a wrong cmd")
			continue
		}
		fmt.Printf("%s - %s\n",p.Cmd,p.Desc)
		if p.Cmd=="help"{
			help()
		}
		if p.Handler!=nil{
			p.Handler()
		}
	}
}

var head *linklist.DataNode=&linklist.DataNode{
	Cmd: "help",
	Desc: "this is help command",
	Handler: nil,
	Next: &linklist.DataNode{
		Cmd: "version",
		Desc: "menu program v1.0",
		Handler: nil,
		Next: &linklist.DataNode{
			Cmd: "quit",
			Desc: "exit the program",
			Handler: quit,
			Next: nil,
		},
	},
}

var help=func() int{
	linklist.ShowAllcmd(head)
	return 0
}

var quit=func() int{
	fmt.Println("bye.")
	os.Exit(0)
	return 0
}