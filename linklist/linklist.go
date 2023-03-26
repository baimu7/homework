package linklist

import "fmt"

func Findcmd(head *DataNode,cmd string) *DataNode{
	if head==nil || cmd==""{
		return nil
	}
	var p *DataNode=head
	for p!=nil{
		if p.Cmd==cmd{
			return p
		}
		p=p.Next
	}
	return nil
}

func ShowAllcmd(head *DataNode) int {
	fmt.Println("Menu List:")
	var p *DataNode =head;
	for p!=nil{
		fmt.Printf("%-7s - %s\n",p.Cmd,p.Desc)
		p=p.Next
	}
	return 0
}