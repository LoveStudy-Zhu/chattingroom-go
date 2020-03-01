package utils

import (
	common "ChatSystem/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)
//将这些方法关联到结构体中
type  Transfer struct {
	//分析它有哪些字段
	Conn net.Conn
	Buf [8096]byte
}


func (this *Transfer)ReadPkg()(mes common.Message,err error) {
	//这里将读取的数据包直接封装成一个函数readPkg(),


	//conn.Read在coon没有被关闭的情况下，才会阻塞
	_,err = this.Conn.Read(this.Buf[:4])
	if err != nil{
		fmt.Println("conn.Read err=",err)
		return
	}
	//根据buf[0:4]转成uint32类型
	var pkgLen uint32
	pkgLen=binary.BigEndian.Uint32(this.Buf[0:4])

	//根据pkgLen读取消息内容
	n,err :=this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil{
		fmt.Printf("conn.Read fail err=%v\n",err)
		return
	}

	//把pkgLen反序列化为Message
	err = json.Unmarshal(this.Buf[:pkgLen],&mes)
	if err != nil{
		fmt.Println("json.Unmarsha err=",err)
		return
	}
	return

}

func (this *Transfer)WritePkg(data []byte) (err error){
	//先发送一个长度给对方，
	var pkgLen uint32
	pkgLen = uint32((len(data)))
	binary.BigEndian.PutUint32(this.Buf[0:4],pkgLen)
	//发送长度
	n,err := this.Conn.Write(this.Buf[:4])
	if n != 4||err !=nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	//发送data本身
	n,err = this.Conn.Write(data)
	if n != int(pkgLen)||err !=nil{
		fmt.Printf("conn.Write(bytes) fail=%v,n=%v,pkgLen=%v\n",err,n,pkgLen)
		return
	}
	return
}