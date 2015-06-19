package network

import (
    "fmt"
    "net"
    "time"
)

type TcpClient struct {
    client_id   int64
    remote_ip   string
    remote_port int
    local_ip    string
    local_port  int 
    conn        net.Conn
}

func NewTcpClient() *TcpClient {
    c := &TcpClient{
	client_id: time.Now().UnixNano(),
	remote_ip: "10.16.77.80",
	remote_port: 8360,
    }

    return c
}

func (this *TcpClient)connect() {
    addr := fmt.Sprintf("%s:%d", this.remote_ip, this.remote_port)
    conn, err := net.Dial("tcp", addr)
    if err != nil {
	fmt.Printf("connect error: %s", err.Error())
	return
    }
    this.conn = conn

}

func (this *TcpClient)send(bytes []byte) error {
    sendingNum := len(bytes)
    sentNum := 0
    for sentNum < sendingNum {
	num, err := this.conn.Write(bytes[sentNum:])
	if err != nil {
	    fmt.Printf("send data failed: %s", string(bytes))
	    return err
	}
	sentNum += num
    }

    return nil
}
