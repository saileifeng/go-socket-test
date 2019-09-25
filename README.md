# go-socket-test
基于tcp,protobuffer3的简单聊天demo

**example** 是一个简单的单服聊天demo

**example3** 是一个基于consul实现服务发现的聊天服务集群,目前还未加入tcp

### 运行方式
    
- **example**
    - **server**
        >go run example/server/server.go --port 8888
    - **client**
        >go run example/client/client.go --addr 127.0.0.1:8888 --maxConns 1000 --interval 1
- **example3**
    - **consul**
        >consul agent -dev
    - **server**
        >go run example3/test/remoteTest.go
    - **server2**
        >go run example3/test2/remoteTest.go


### 代码优化

- 创建1w+客户端连接的时候服务端报错

  ```shell
  runtime: program exceeds 10000-thread limit
  fatal error: thread exhaustion
  ```

  排查之后发现是自己引用了[tcpkeepalive][tcpkeepalive]，服务端在每一次创建客户端连接的调用该代码都会去创建一个线程，注释掉该代码
  ```shell
	kaConn, _ := tcpkeepalive.EnableKeepAlive(session.conn)
	kaConn.SetKeepAliveIdle(30*time.Second)
	kaConn.SetKeepAliveCount(4)
	kaConn.SetKeepAliveInterval(5*time.Second)
  ```


### 客户端机器优化

- 修改操作系统端口范围，让客户端机器可以开启更多的tcp连接（cannot assign requested address）

    ```shell
    vim /etc/sysctl.conf
    net.ipv4.ip_local_port_range = 1024 65535
    sysctl -p #刷新生效
    ```

### 服务端机器优化

- 修改文件句柄数

    ```shell
    echo 10240000 > /proc/sys/fs/nr_open
    ```
    
    ```shell
    vim /etc/security/system.conf
    DefaultLimitCORE=infinity
    DefaultLimitNOFILE=10240000
    DefaultLimitNPROC=10240000
    ```
    
    ```shell
    vim /etc/systemd/limits.conf
    root            soft    nofile          10240000
    ```
    
	```shell
	vim /etc/systemd/limits.conf
	root            soft    nofile          10240000
	root            hard    nofile          10240000
	root            soft    nproc           10240000
	root            hard    nproc           10240000
	```
	
    ```shell
    vim /etc/security/limits.d/20-nproc.conf
    *          soft    nproc     10240000
    ```
    
    ```shell
    reboot
    ```
- 待续

[tcpkeepalive]: https://github.com/felixge/tcpkeepalive "tcpkeepalive"
[runtime]: https://godoc.org/runtime "runtime"


