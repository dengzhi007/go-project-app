package network

import (
    "fmt"
    "net"
    "common"
    "time"
    "log"
    "os"
)

type TcpServer struct {
    ip         string  // localhost
    port       int
    maxClients int     // 10000 concurrency connections
    clientNum  int
    connPool   map[int64]*net.Conn
    maxBodyLen int     // 100k bytes
    running    bool
    logger     *log.Logger
    heartbeat  int64
}

func NewTcpServer() *TcpServer {
    s := &TcpServer{
	ip: "",
	port: 8080,
	maxClients: 10000,
	clientNum: 0,
	maxBodyLen: 102400,
	running: false,
	heartbeat: time.Now().Unix(),
    }

    logFileName := common.SysPath + "/logs/tcpserver.log"
    logFile, err := os.Create(logFileName)
    if err != nil {
	fmt.Println("tcp server open /logs/tcpserver.log failed")
	return nil
    }
    s.logger = log.New(logFile, "", log.LstdFlags)

    return s
}

// ip  0:in  1:out 2:invalid
func CheckIp(ip string) int {
    return 0

}

func (this *TcpServer)Start() {
    addr := fmt.Sprintf(":%d", this.port)
    listener, err := net.Listen("tcp", addr)
    if err != nil {
	this.logger.Printf("tcp server listen on port: %d failed:%s\n", this.port, err.Error())
	return
    }

    defer listener.Close()

    this.running = true
    this.logger.Printf("tcp server start to listen on %d\n", this.port)
    
    for this.running {
	this.heartbeat = time.Now().Unix()

	conn, err := listener.Accept()
	if err != nil {
	    this.logger.Printf("tcp server accept failed:%s\n", err.Error())
	    return 
	}

	this.clientNum += 1
	if this.clientNum > this.maxClients {
	    this.logger.Println("client num exceed maxClients")
	    conn.Close()
	    continue
	}

	// add to clientPool
	//...

	go this.process(conn)
    }

}

func (this *TcpServer)process(conn net.Conn) {
    defer conn.Close() 

    //headBuf := make([]byte, 10)
    //bodyBuf := make([]byte, this.maxBodyLen)
    ap := new(common.AlarmPackage)
    var num int = 0
    var bytes []byte = nil
    var err error = nil

    for {
	bytes, err = this.read(10, conn)
	if err != nil {
	    break
	}

	fmt.Sscanf(string(bytes),"%d", &num)
	if num <= 0 || num > 102400 {
	    break
	}

	fmt.Printf("head receive %d bytes\n", num)
	bytes, err = this.read(num, conn)
	if err != nil {
	    break
	}
	fmt.Printf("body bytes:%s\n", string(bytes))

	ap.UpdateFromBytes(bytes)
	
    }

    this.clientNum -= 1
}

func (this *TcpServer)read(length int, conn net.Conn) ([]byte, error) {
    readNum := 0
    buffer := make([]byte, length)
    for readNum < length {
	receivedNum, err := conn.Read(buffer[ readNum : ])
	if err != nil {
	    return nil, err
	}
	readNum += receivedNum
    }
    return buffer, nil
}

func (this *TcpServer)Stop() {
    this.running = false
    //this.clientNum = 0
    //stop all connections in clientPool
    //...
    this.logger.Println("server stop")
}









