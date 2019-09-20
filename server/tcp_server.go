package server

import (
	"errors"
	"github.com/saileifeng/go-socket-test/utils"
	"log"
	"net"
	"sync"
	"sync/atomic"
)

type Handler interface {
	ReciveMsg(session *TCPSession,msg []byte)
	Closed(err error)
}

type TCPSession struct {
	conn    net.Conn
	handler Handler
	s *TCPServer
}

type TCPServer struct {
	network                         string
	address                         string
	packHeader, maxReadLength, maxBody int
	handlerFunc func() Handler
	Conns sync.Map
	CurrentConns int64
}

func NewTCPServer(network, address string) *TCPServer {
	return &TCPServer{packHeader: 4,maxReadLength:128,maxBody:65535,network:network,address:address,Conns:sync.Map{}}
}

func (s *TCPServer)BindController(f func() Handler) *TCPServer {
	s.handlerFunc = f
	return s
}

func (t *TCPSession)SendMsg(msg []byte) {
	//log.Println("to client msg :",string(msg))
	t.conn.Write(Pack(msg))
}

//启动tcp监听
func (s *TCPServer)ListenTCP() error {
	listen, err := net.Listen(s.network, s.address)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("server listen :",listen.Addr())
	go func() {
		defer listen.Close()
		for {
			conn, err := listen.Accept()
			if err != nil {
				log.Println("err conn :", err)
			}
			go connHandle(&TCPSession{conn:conn,handler:s.handlerFunc(),s:s})
		}
	}()
	return nil
}

//connHandle 处理conn接收到的包
func connHandle(session *TCPSession) {
	defer session.conn.Close()
	//key := session.conn.RemoteAddr().String()
	//defer session.s.Conns.Delete(key)
	//session.s.Conns.LoadOrStore(key  ,session)
	defer atomic.AddInt64(&session.s.CurrentConns,-1)
	atomic.AddInt64(&session.s.CurrentConns,1)
	//设置心跳 这个包有问题，会导致线程无限创建
	//kaConn, _ := tcpkeepalive.EnableKeepAlive(session.conn)
	//kaConn.SetKeepAliveIdle(30*time.Second)
	//kaConn.SetKeepAliveCount(4)
	//kaConn.SetKeepAliveInterval(5*time.Second)

	readBuff := make([]byte, session.s.maxReadLength)
	tempBuff := make([]byte, 0)
	//包头数据大小
	headLength := 0
	dataBuff := make([]byte, 0)
	//session := &TCPSession{handler:}
	for {
		n, err := session.conn.Read(readBuff)
		//异常关闭
		if err != nil {
			session.handler.Closed(err)
			log.Println("read conn byte err :", err)
			return
		}
		tempBuff = append(tempBuff, readBuff[:n]...)
		tempDataBuff := make([]byte, 0)
		tempDataBuff, tempBuff, headLength, err = UnPack(tempBuff, session.s.packHeader, headLength, session.s.maxBody)
		if err != nil {
			log.Println("unPack err :", err)
			session.handler.Closed(err)
			return
		}
		dataBuff = append(dataBuff, tempDataBuff[:]...)
		if headLength != 0 {
			continue
		}
		//TODO 处理业务
		session.handler.ReciveMsg(session,dataBuff)
		dataBuff = []byte{}
	}
}

//TODO 需要对打包的包头大小进行判断并生成包头
func Pack(msg []byte) []byte {
	headBytes := utils.IntToBytes(len(msg))
	//fmt.Println("------ ",headBytes)
	return append(headBytes,msg[:]...)
}

//unPack 解析数据包，处理黏包，断包
//args：readBuff::解析的字节切片，packHeader::包头长度，headLength::需要解析的包大小
//return 当前解析出来的字节，剩余未解析字节，当前还差多少个字节需要解析
//TODO 后续需要重构packHeader
func UnPack(readBuff []byte, packHeader, headLength, maxBody int) (data []byte, tempBuff []byte, headCount int, err error) {
	header := headLength
	if headLength == 0 {
		//如果当前解析的字节切片长度小于解析的协议头长度，说明断包了
		if packHeader > len(readBuff) {
			return []byte{}, readBuff, 0, nil
		}
		headerBytes := readBuff[:packHeader]
		header = utils.BytesToInt(headerBytes)
		readBuff = readBuff[packHeader:]
	}
	//非法数据解析
	if header == 0 || header > maxBody {
		return nil, nil, 0, errors.New("header unNormal")
	}
	readBuffLen := len(readBuff)
	//如果解析的数据比包头中标示的数据大，那么有可能黏包或者刚刚好
	if len(readBuff) >= header {
		//temp := header - len(readBuff)
		//if temp < 0 {
		//	temp = 0
		//}
		return readBuff[:header], readBuff[header:], 0, nil
	}
	//如果解析的数据比包头中标示的数据小，那么说明断包了，需要再次解析拼接
	return readBuff, []byte{}, header - readBuffLen, nil
}
