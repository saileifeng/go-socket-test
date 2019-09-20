package main

import (
	"flag"
	"fmt"
	"github.com/saileifeng/go-socket-test/server"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"
)

var port = 8888

func main() {

	flag.IntVar(&port,"port",8888,"--port default 8888")
	flag.Parse()
	ts := server.NewTCPServer("tcp",fmt.Sprintf(":%d",port))
	ts.BindController(func() server.Handler {
		return &MyController{}
	})
	ts.ListenTCP()
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

type MyController struct {

}

func (c *MyController)ReciveMsg(session *server.TCPSession,msg []byte)  {
	//log.Println(string(msg))
	session.SendMsg(msg)
}

func (c *MyController)Closed(err error)  {
	log.Println(err)
}

