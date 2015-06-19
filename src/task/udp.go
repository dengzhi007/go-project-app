package task

import (

)

const (

)

type UdpTask struct {
    task_num int
    running  bool
}

func NewUdpTask() *UdpTask {
    u := &UdpTask{
	task_num: 0,
	running: false,
    }
    return u
}

func (this *UdpTask)Start() {
    this.running = true

    for this.running {
    
    }
}

func (this *UdpTask)Stop() {
    this.running = false
}
