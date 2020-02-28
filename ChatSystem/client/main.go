package main

import (
	"fmt"
	"os"
)

var userId int
var passWord string

func main()  {
	var key int
	var loop =true
	for loop {
		fmt.Println("-------------欢迎登录聊天系统-----------------")
		fmt.Println("\t\t1.登录聊天室")
		fmt.Println("\t\t2.注册用户")
		fmt.Println("\t\t3.退出系统")
		fmt.Println("请选择（1-3）")
		fmt.Println("--------------------------------------------")
		fmt.Scanf("%d\n",&key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop =false
		case 3:
			fmt.Println("退出系统")
			//loop =false
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

	if key==1{
		fmt.Println("请输入用户id:")
		fmt.Scanf("%d",&userId)
		fmt.Println("请输入用户密码:")
		fmt.Scanf("%s",&passWord)
		err:=login(userId,passWord)
		if err !=nil {
		}else {
		}
	}
}
