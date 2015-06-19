package task

import (

)

const (

)

type SendTask struct {
    task_num int
    running  bool

}

func NewSendTask() *SendTask {
    s := &SendTask{
	task_num: 0,
	running: false,
    }
    return s
}

func (this *SendTask)Start() {
   this.running = true

   for this.running {
   
   }
}

func (this *SendTask)Stop() {
    this.running = false
}
