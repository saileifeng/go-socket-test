package main

import (
	"flag"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/saileifeng/go-socket-test/example/pb"
	"github.com/saileifeng/go-socket-test/server"
	"github.com/saileifeng/go-socket-test/utils"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
	"time"
)


var port = 8888

var consulAddr = "127.0.0.1:8500"

var serviceName = "test"

func main() {

	flag.IntVar(&port,"port",8888,"--port default 8888")
	flag.StringVar(&consulAddr, "registry_address", "127.0.0.1:8500", "registry address")
	flag.StringVar(&serviceName, "serviceName", "test", "serviceName")
	flag.Parse()


	//启动tcp服务
	ts := server.NewTCPServer("tcp",fmt.Sprintf(":%d",port))
	ts.BindController(func() server.Handler {
		return &MyController{}
	})
	ts.ListenTCP()
	//打印当前在线人数
	var tempCurrentConns int64
	go func() {
		for ; ;  {
			time.Sleep(time.Second)
			if tempCurrentConns != ts.CurrentConns {
				log.Println("current connections : ",ts.CurrentConns)
			}
			tempCurrentConns = ts.CurrentConns
		}
	}()

	//开启pprof
	go func() {
		http.ListenAndServe(":9999", nil)
	}()

	t := make(chan int,0)
	go ShutDownHook(t)
	<-t
}


func ShutDownHook(t chan int) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	a := <-quit
	log.Println("close ",a)
	t<-0
}

var onlines  = sync.Map{}

type MyController struct {
	UID string
}

func (c *MyController) Init()  {
	
}

func (c *MyController)ReceiveMsg(session *server.TCPSession,msg []byte)  {
	switch GetPbNum(msg) {
	case pb.PbNum_Login:
		login := &pb.C_LoginRequest{}
		//err := proto.Unmarshal(msg[4:],login)
		//if err != nil{
		//	log.Println(err)
		//}
		GetMessage(msg,login)
		//储存登陆状态用户
		onlines.LoadOrStore(login.AccountID,c)
		//log.Println(login)
		SendMsg(session,pb.PbNum_Login,&pb.S_LoginResponse{RequestStatus:pb.RequestStatus_SUCCESS,UID:login.AccountID})
	case pb.PbNum_Say:
		say := &pb.C_ChatToSomeOneRequest{}
		GetMessage(msg,say)
		msgID := time.Now().UnixNano()
		//log.Println()
		SendMsg(session,pb.PbNum_Say,&pb.S_ChatToSomeOneResponse{RequestStatus:pb.RequestStatus_SUCCESS,ChatMessageStatus:pb.ChatMessageStatus_Reached,
			TargetUID:say.TargetUID,MessageID:msgID})
		
	}
	log.Println("ReciveMsg",GetPbNum(msg),msg)
}

func (c *MyController)Closed(err error)  {
	log.Println(err)
}

func GetPbNum(msg []byte) pb.PbNum {
	return pb.PbNum(utils.BytesToInt32(msg[:4]))
}

func GetMessage(msgByte []byte,msg proto.Message) error {
	err := proto.Unmarshal(msgByte[4:],msg)
	if err != nil{
		log.Println(err)
		return err
	}
	log.Println(msg)
	return nil
}


func SendMsg(session *server.TCPSession,pbNum pb.PbNum,msg proto.Message)  {
	msgByte,err := proto.Marshal(msg)
	if err != nil {
		log.Println(err)
		return
	}
	pbNumByte := utils.Int32ToBytes(int32(pbNum))
	w := append(pbNumByte,msgByte[:]...)
	session.SendMsg(w)
}