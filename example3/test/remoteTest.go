package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/saileifeng/go-socket-test/example3/innerPb"
	"github.com/saileifeng/go-socket-test/example3/remote"
	"github.com/saileifeng/go-socket-test/utils"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)



type Actor struct {
	message chan interface{}
	Pid *remote.Pid
}

func NewActor(pid *remote.Pid) *Actor{
	actor := &Actor{message:make(chan interface{},10),Pid:pid}
	go actor.Do()
	return actor
}

func (actor *Actor)Do()  {
	for  msg := range actor.message {
		//time.Sleep(time.Microsecond*100)
		//select {
		//case msg := <-actor.message:
			switch msg.(type) {
			case *innerPb.RPC_ChatToSomeOneMessageRequest:
				log.Println("actor do",msg)
			default:


			}
		//default:
		//}
	}
}

func (actor *Actor)SendMsg(msg interface{})  {
	actor.message <- msg
}


var localRouterMaps  = &sync.Map{}

type ChatService struct {

}

func (cs *ChatService)ChatToSomeOne(ctx context.Context, req *innerPb.RPC_ChatToSomeOneMessageRequest) (*innerPb.RPC_ChatToSomeOneMessageResponse, error)  {
	//查找出节点上对应的处理对象，将消息路由到该对象的消息队列里面
	pid := &remote.Pid{Addr:req.TargetPID.Addr,Id:req.TargetPID.Id}

	actor,ok := localRouterMaps.Load(pid.String())
	if !ok {
		return nil,errors.New("no proc")
	}
	actor.(*Actor).SendMsg(req)
	//pid := req.SenderPID
	//router.RouterMessage(&router.PID{Address:pid.Addr,Id:pid.Id},req)
	//根据pid找出对应的对象

	return &innerPb.RPC_ChatToSomeOneMessageResponse{MessageID:req.MessageID,ChatMessageStatus:innerPb.ChatMessageStatus_Reached},nil
}


func RPCSendMessageToSomeOne(targetName ,senderName,msg string,senderPid *remote.Pid) (*innerPb.RPC_ChatToSomeOneMessageResponse,error) {
	//根据targetUID找出对应的pid
	pid,err := cluster.GetPidWithName(targetName)
	log.Printf("RPCSendMessageToSomeOne : %v <%v> send \"%v\" to %v <%v>",senderName,senderPid,msg,targetName,pid)
	if err != nil {
		return nil,err
	}
	cc,ok := remote.GetClientConn(pid.Addr)
	if ok {
		request := &innerPb.RPC_ChatToSomeOneMessageRequest{
			SenderPID:&innerPb.Pid{Addr:senderPid.Addr,Id:senderPid.Id},
			TargetPID:&innerPb.Pid{Addr:pid.Addr,Id:pid.Id},
			Message:msg,
			MessageID:time.Now().UnixNano(),
			SenderUID:senderName,
			TargetUID:targetName,
		}
		return innerPb.NewChatServiceClient(cc).ChatToSomeOne(context.Background(),request)
		//if err != nil {
		//	return resp,nil
		//}
		//return nil,err
	}
	return nil,errors.New("no grpc client")
}

var consulAddr = "127.0.0.1:8500"
var serviceName = "mytest"
var port = 0

var cluster *remote.Remote

func main() {
	flag.IntVar(&port,"port",8888,"--port default 8888")
	cluster = remote.CreateCluster(consulAddr,serviceName,utils.LocalIP(),port)
	innerPb.RegisterChatServiceServer(cluster.ConsulRegistry.Server,&ChatService{})

	for i:= 0; i < 100;i++  {
		go func(i int) {
			pid :=remote.NewPid(cluster.ConsulRegistry.ServerAddr)
			actor := NewActor(pid)

			localRouterMaps.Store(pid.String(),actor)

			//log.Println(actor,pid)
			name := fmt.Sprintf("%v%v","number",i)
			err := cluster.PutPidWithName(name+""+pid.Addr,pid)
			log.Println("PutPidWithName",name,pid,err)
			//
			//p,err := cluster.GetPidWithName(name)
			//log.Println(p,err)
		}(i)

	}
	log.Println("put finished")

	cluster.Run()
}




func ShutDownHook(t chan int) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	a := <-quit
	log.Println("close ",a)
	t<-0
}