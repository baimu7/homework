package main

import ("fmt")

func main(){
	for{
		var cmd string
		fmt.Println("请输入命令:")
		fmt.Scanln(&cmd)
		if cmd=="help"{
			fmt.Println("命令列表:")
			fmt.Println("--quit")
			fmt.Println("--hellp")
		}else if cmd=="quit"{
			break
		}else if cmd=="hellp"{
			fmt.Println("This is help cmd")
		}else {
			fmt.Println("Wrong cmd!")
		}
	}
}
