package consul

import (
	"fmt"
	"github.com/saileifeng/go-socket-test/registry/consul/register"
	"github.com/saileifeng/go-socket-test/registry/consul/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/health/grpc_health_v1"
	resolver2 "google.golang.org/grpc/resolver"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type CallBack struct {
	
}

func (c *CallBack) Do(state resolver2.State)  {
	for _, v := range state.Addresses {
		log.Println("CallBack",v.Addr)
	}
}

func NewClietnConnWithUpdateCallBack(consulAddr,serviceName string,callBack resolver.UpdateStateCallBack) *grpc.ClientConn {
	schema, err := resolver.StartConsulResolverWithState(consulAddr, serviceName,callBack)
	log.Println("NewClietnConn schema :",schema)
	//consul集群在未完成选举的时候会创建失败,需要再次重试创建,这种问题似乎出现在低版本的consul
	for i := 0;i<10;i++ {
		if err!=nil {
			time.Sleep(time.Second)
			log.Println("retry NewClietnConn")
			schema, err = resolver.StartConsulResolver(consulAddr, serviceName)
		}else {
			break
		}
	}
	if err != nil {
		log.Fatal("init consul resovler err", err.Error())
	}

	// Set up a connection to the server.
	conn, err := grpc.Dial(fmt.Sprintf("%s:///%s", schema,serviceName), grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

//NewClietnConn 创建客户端
func NewClietnConn(consulAddr,serviceName string) *grpc.ClientConn {
	return NewClietnConnWithUpdateCallBack(consulAddr,serviceName,nil)
}

//Registry 服务注册自定义结构体
type Registry struct {
	consulAddr,serviceName string
	port int
	Listener net.Listener
	Server *grpc.Server
	Register *register.ConsulRegister
	ServerAddr string
}


//NewRegister 创建新的服务注册
func NewRegister(consulAddr,service ,ip string,port int) *Registry {
	listener, err := net.Listen("tcp", fmt.Sprintf("%v:%v",ip,port))
	if err != nil {
		log.Fatalln(err)
	}
	addrs := strings.Split(listener.Addr().String(),":")
	port,err = strconv.Atoi(addrs[len(addrs)-1])
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("start server port :",addrs[len(addrs)-1])
	//consul service register
	nr := register.NewConsulRegister(consulAddr,service,ip,port)
	nr.Register()
	//start grpc server
	serv :=  grpc.NewServer()
	//registe health check
	grpc_health_v1.RegisterHealthServer(serv, &register.HealthImpl{})

	return &Registry{consulAddr:consulAddr,serviceName:service,port:port,Listener:listener,Server:serv,Register:nr,ServerAddr:listener.Addr().String()}
}
//Run 启动
func (r *Registry)Run()  {
	//server hook
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
		<-quit
		log.Println("do run hook")
		err := r.Register.Deregister()
		log.Println("Deregister",err)
		r.Server.Stop()
	}()

	if err := r.Server.Serve(r.Listener); err != nil {
		panic(err)
	}
}
