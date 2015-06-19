package task

import (

)

const (

)

type TcpTask struct {
    task_num int
    running  bool
}

func NewTcpTask() *TcpTask {
    t := &TcpTask{
	task_num: 0,
	running: false,
    }
    return t
}

func (this *TcpTask)Start() {
    this.running = true

    for this.running {
    
    }
}

func (this *TcpTask)Stop() {
    this.running = false
}
