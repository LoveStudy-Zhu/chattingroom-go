package main

import (
	common "ChatSystem/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func readPkg(conn net.Conn)(mes common.Message,err error) {
	//这里将读取的数据包直接封装成一个函数readPkg(),
	buf:= make([]byte,8096)

	//conn.Read在coon没有被关闭的情况下，才会阻塞
	_,err = conn.Read(buf[:4])
	if err != nil{
		fmt.Println("conn.Read err=",err)
		return
	}
	//根据buf[0:4]转成uint32类型
	var pkgLen uint32
	pkgLen=binary.BigEndian.Uint32(buf[0:4])

	//根据pkgLen读取消息内容
	n,err :=conn.Read(buf[:pkgLen])
	if  n != int(pkgLen) || err != nil{
		fmt.Printf("conn.Read fail err=%v,n=%v,pkgLen=%v",err,n,pkgLen)
		return
	}

	//把pkgLen反序列化为Message
	err = json.Unmarshal(buf[:pkgLen],&mes)
	if err != nil{
		fmt.Println("json.Unmarsha err=",err)
		return
	}
	return

}

func writePkg(conn net.Conn,data []byte) (err error){
	//先发送一个长度给对方，
	var pkgLen uint32
	pkgLen = uint32((len(data)))
	var bytes [4]byte
	binary.BigEndian.PutUint32(bytes[0:4],pkgLen)
	//发送长度
	n,err := conn.Write(bytes[:4])
	if n != 4||err !=nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	//发送data本身
	n,err = conn.Write(bytes[:4])
	if n != int(pkgLen)||err !=nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	return
}
