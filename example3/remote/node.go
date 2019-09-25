package remote

import (
	"encoding/json"
	"github.com/saileifeng/go-socket-test/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
)

//NodeInfoMap 节点信息map
type NodeInfoMap struct {
	nodes map[string]*grpc.ClientConn
	mux sync.RWMutex
}

//Put ....
func (node *NodeInfoMap)Put(key string,value *grpc.ClientConn)  {
	node.mux.Lock()
	defer node.mux.Unlock()
	node.nodes[key] = value
}


//Get ...
func (node *NodeInfoMap)Get(key string) (*grpc.ClientConn,bool) {
	node.mux.RLock()
	defer node.mux.RUnlock()
	value := node.nodes[key]
	if value == nil{
		return nil,false
	}
	return value,true
}

//Del ...
func (node *NodeInfoMap)Del(key string)  {
	node.mux.Lock()
	defer node.mux.Unlock()
	delete(node.nodes,key)
}

//ForEach ...
func (node *NodeInfoMap)ForEach(fun func(key string,value *grpc.ClientConn))  {
	node.mux.Lock()
	defer node.mux.Unlock()
	for k, v := range node.nodes {
		fun(k,v)
	}
}

//MyCallBack ...
type MyCallBack struct {

}

var nodeInfo = &NodeInfoMap{nodes: map[string]*grpc.ClientConn{}}


//Do 处理当前节点信息，将无效的节点去除掉
func (c *MyCallBack) Do(state resolver.State)  {
	//log.Println("------------------",state)
	temp := make(map[string]bool,1)
	//更新节点信息
	for _, v := range state.Addresses {
		//log.Println("CallBack",v.Addr)
		_,ok := nodeInfo.Get(v.Addr)
		if ok{
			//temp.(*grpc.ClientConn).Close()
		}else {
			cc,err := grpc.Dial(v.Addr,grpc.WithInsecure())
			if err == nil {
				nodeInfo.Put(v.Addr,cc)
			}
		}
		temp[v.Addr] = true
	}
	//删除失效节点信息
	tempDel := make([]string,0)
	nodeInfo.ForEach(func(key string, value *grpc.ClientConn) {
		k := temp[key]
		//log.Println(" -------- ",key,k)
		if !k {
			//_,ok := nodeInfo.Get(key)
			//if ok {
			//err := v.Close()
			//log.Println(err)
			//nodeInfo.Del(key)
			tempDel = append(tempDel,key)
			//}
		}
	})

	for _, value := range tempDel {
		v,ok := nodeInfo.Get(value)
		if ok {
			err := v.Close()
			log.Println("close disconnect",v.Target(),err)
			nodeInfo.Del(value)
		}
	}
	//log.Println(nodeInfo.nodes)
}

//GetClientConn 获得远程节点的grpc连接
func GetClientConn(addr string) (*grpc.ClientConn,bool) {
	return nodeInfo.Get(addr)
}

//Remote 远程节点信息
type Remote struct {
	ConsulRegistry *consul.Registry
}

//Run 启动远程节点服务
func (remote *Remote) Run()  {
	remote.ConsulRegistry.Run()
}

//RegisteRpcMessage 注册rpc服务消息
func (remote *Remote) RegistryRpcMessage(sd *grpc.ServiceDesc, ss interface{})  {
	remote.ConsulRegistry.Server.RegisterService(sd, ss)
}

//PutPidWithName 注册rpc全局pid
func (remote *Remote) PutPidWithName(name string,pid *Pid) error {
	pidBytes,err := json.Marshal(pid)
	if err != nil{
		return err
	}
	return remote.ConsulRegistry.Register.PutKV(name,pidBytes)
}

//GetPidWithName 获取rpc全局pid
func (remote *Remote) GetPidWithName(name string) (*Pid,error) {
	pidBytes,err :=remote.ConsulRegistry.Register.GetValue(name)
	if err != nil{
		return nil,err
	}
	pid := &Pid{}
	err =json.Unmarshal(pidBytes,pid)
	if err != nil{
		return nil,err
	}
	return pid,nil
}


//CreateRemote 启动远程节点监听
func CreateCluster(consulAddr,service,ip string,port int) *Remote {
	consul.NewClietnConnWithUpdateCallBack(consulAddr, service,&MyCallBack{})
	reg := consul.NewRegister(consulAddr, service, ip,port)
	return &Remote{ConsulRegistry:reg}
}

//var localIP  = utils.LocalIP()

var pidID int64

type Pid struct {
	Addr string `protobuf:"bytes,1,opt,name=Addr,proto3" json:"Addr,omitempty"`
	Id string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (pid *Pid)String() string {
	return pid.Addr+"/"+pid.Id
}

func NewPid(serverAddr string) *Pid {
	return &Pid{Addr:serverAddr,Id:strconv.FormatInt(atomic.AddInt64(&pidID,1),10)}
}