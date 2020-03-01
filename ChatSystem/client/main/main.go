package main

import (
	"ChatSystem/client/process"
	"fmt"
	"os"
)

var userId int
var passWord string
var useName string

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
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d",&userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s",&passWord)
			//创建一个userProcess的实例
			up := &process.UserProcess{}
			up.Login(userId,passWord)

			//loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id:")
			fmt.Scanf("%d\n",&userId)
			fmt.Println("请输入用户密码:")
			fmt.Scanf("%s\n",&passWord)
			fmt.Println("请输入用户名称:")
			fmt.Scanf("%s\n",&useName)

			up := &process.UserProcess{}
			up.Register(userId,passWord,useName)

			//loop =false
		case 3:
			fmt.Println("退出系统")
			//loop =false
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

}
