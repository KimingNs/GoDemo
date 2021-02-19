package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/Project/project04/redis/common/message"
	"net"
)

//这里将这些方法关联到结构体中
type Transfer struct {
	//分析它应该有哪些字段
	Conn   net.Conn   `json:"conn"`
	Buffer [8096]byte `json:"buffer"`
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buffer[0:4], pkgLen)
	//发送长度
	n, err := this.Conn.Write(this.Buffer[:4])
	if err != nil {
		fmt.Println("conn.Write() fail", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buf := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据")
	_, err = this.Conn.Read(this.Buffer[:4])
	if err != nil {
		fmt.Println("utils.go.conn.Read err=", err)
		return
	}
	//根据buf[:4] 转成一个 uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buffer[0:4])
	//根据pkgLen 读取消息内容
	n, err := this.Conn.Read(this.Buffer[:pkgLen])
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Read fail err=", err)
		return
	}
	//把pkgLen 反序列化成-》message.Message
	err = json.Unmarshal(this.Buffer[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}
