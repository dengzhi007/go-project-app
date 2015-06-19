package task

import (

)

const (

)

type PingTask struct {
    task_num int
    running  bool 

}

func NewPingTask() *PingTask {
    p := &PingTask{
	task_num: 0,
	running: false,
    }
    return p
}


func (this *PingTask)Start() {
    this.running = true

    for this.running {
	
    }
}

func (this *PingTask)Stop() {
    this.running = false
}
