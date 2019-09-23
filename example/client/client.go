package main

import (
	"flag"
	"github.com/golang/protobuf/proto"
	"github.com/saileifeng/go-socket-test/example/pb"
	"github.com/saileifeng/go-socket-test/server"
	"github.com/saileifeng/go-socket-test/utils"
	"log"
	"net"
	"time"
)


var maxConns = 0
var interval int64
var addr = "127.0.0.1:8888"

func main()  {
	flag.IntVar(&maxConns,"maxConns",1000,"--maxConns default 1000")
	flag.Int64Var(&interval,"interval",100,"--interval default 100 ,this is ms")
	flag.StringVar(&addr,"addr","","--addr default 127.0.0.1:8888 ")
	flag.Parse()
	temp := time.Duration(interval)
	log.Println("interval :",time.Millisecond*temp )
	for i := 0;i < maxConns ; i++ {
		time.Sleep(time.Millisecond*temp)
		go client()
	}
	a := make(chan int,1)
	<-a
}

func client() {
	////通过tcp 协议链接 本机 8080端口
	con, err := net.Dial("tcp", addr)
	//如果出现错误 说明链接失败
	if err != nil {
		log.Println("连接服务器端失败",err.Error())
		return
	}
	//记得关闭
	defer con.Close()
	flag := true
	//开始向服务器端发送 hello
	f := func(){
		//b := []byte("hello world!")
		//
		//bLenBytes := utils.IntToBytes(len(b))
		//
		//b2 := []byte("i am golang developer")
		//
		//
		//
		//bLenBytes2 := utils.IntToBytes(len(b2))
		//
		//
		//w := append(bLenBytes,b[:]...)
		//w1 := append(bLenBytes2,b2[:]...)
		//
		//
		//ww := append(w,w1[:]...)

		login := &pb.C_LoginRequest{AccountID:"123456",Password:"654321"}
		loginBytes,err := proto.Marshal(login)
		if err != nil {
			log.Println("proto.Marshal err :",err)
			return
		}

		pbNumBytes := utils.IntToBytes(int(pb.PbNum_Login))

		w := append(pbNumBytes,loginBytes[:]...)
		log.Println("w",w)
		wLen := utils.IntToBytes(len(w))
		//log.Println("wLen",wLen)
		ww := append(wLen,w[:]...)
		//log.Println("ww",ww)
		//log.Println()
		num, write_err := con.Write(ww)
		//如果写入有问题 输出对应的错误信息
		if write_err != nil {
			flag = false
			log.Println(con.LocalAddr(),num,write_err.Error())
		}
		//如果没有问题。显示对应的写入长度
		//fmt.Println(num)
	}

	go connHandle(con,128,4,65535,&flag)

	for flag{
		//log.Println(flag)
		time.Sleep(time.Second*1)
		f()
	}


	//log.Println(utils.IntToBytes(1))
	//a := make(chan int,1)
	//<-a
	//fmt.Println(int(binary.BigEndian.Uint32(IntToBytes(65535))))
}

//connHandle 处理conn接收到的包
func connHandle(conn net.Conn,readLength,packHeader,maxBody int,flag *bool) {
	defer conn.Close()
	readBuff := make([]byte, readLength)
	tempBuff := make([]byte, 0)
	//包头数据大小
	headLength := 0
	dataBuff := make([]byte, 0)
	//session := &TCPSession{handler:}
	for {
		n, err := conn.Read(readBuff)
		//异常关闭
		if err != nil {
			log.Println("read conn byte err :", err)
			flag = new(bool)
			return
		}
		//log.Println("recive msg num :",n)
		tempBuff = append(tempBuff, readBuff[:n]...)
		tempDataBuff := make([]byte, 0)
		tempDataBuff, tempBuff, headLength, err = server.UnPack(tempBuff, packHeader, headLength, maxBody)
		if err != nil {
			log.Println("unPack err :", err)
			flag = new(bool)
			return
		}
		dataBuff = append(dataBuff, tempDataBuff[:]...)
		//log.Println("-----",string(dataBuff))
		if headLength != 0 {
			continue
		}
		//TODO 处理业务
		//log.Println(string(dataBuff))
		doMsg(dataBuff)
		dataBuff = []byte{}
	}
}

func doMsg(msg []byte)  {
	switch pb.PbNum(utils.BytesToInt32(msg[:4])) {
	case pb.PbNum_Login:
		login := &pb.S_LoginResponse{}
		err := proto.Unmarshal(msg[4:],login)
		if err != nil {
			log.Println(err)
		}
		log.Println(login.RequestStatus,login.UID)
	}
}