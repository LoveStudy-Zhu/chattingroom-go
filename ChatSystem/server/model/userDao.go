package model

import (
	common "ChatSystem/common/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var(
	MyUserDao *UserDao
)

//在服务器启动后，就初始化一个userDao
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式创建一个UserDao实例
func  NewUserDao(pool *redis.Pool) (userDao *UserDao){
	userDao =&UserDao{
		pool:pool,
	}
	return
}


func (this *UserDao) getUserById(conn redis.Conn,id int) (user *User,err error){
	//通过给定id去redis里查询用户
	res,err :=redis.String(conn.Do("hget","users",id))
	if err !=nil{
		if err == redis.ErrNil{//在users哈希中，没有找到对应id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//这里需要把res反序列化为User实例
	user =&User{}
	err = json.Unmarshal([]byte(res),&user)
	if err != nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return

}

//完成登录校验Login
func (this *UserDao) Login(userId int,userPwd string) (user *User,err error){
	conn := this.pool.Get()
	defer conn.Close()
	user ,err = this.getUserById(conn,userId)
	if err !=nil{
		return
	}

	if user.UserPwd != userPwd{
		err = ERROR_USER_PWD
		return
	}
	return
}


func (this *UserDao) Register(user *common.User) (err error){
	conn := this.pool.Get()
	defer conn.Close()
	_,err = this.getUserById(conn,user.UserId)
	if err ==nil{
		err = ERROR_USER_EXISTS
		return
	}
	//这时说明id在redis还没有，则可以完成注册
	data,err:=json.Marshal(user)//序列化
	if err != nil{
		return 
	}
	//入库
	_,err=conn.Do("HSet","users",user.UserId,string(data))
	if err !=nil{
		fmt.Println("保存用户错误 err =",err)
		return 
	}
	return 
}
