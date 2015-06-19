package task

import (

)

const (

)

type HttpTask struct {
    task_num int
    running  bool

}

func NewHttpTask() *HttpTask {
    h := &HttpTask{
	task_num: 0,
	running: false,
    }
    return h
}

func (this *HttpTask)Start() {
    this.running = true

    for this.running {
    
    }
}

func (this *HttpTask)Stop() {
    this.running = false
}
