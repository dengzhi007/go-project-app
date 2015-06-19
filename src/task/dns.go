package task

import (

)

const (

)

type DnsTask struct {
    task_num int
    running  bool

}

func NewDnsTask() *DnsTask {
    d := &DnsTask{
	task_num: 0,
	running: false,
    }
    return d
}

func (this *DnsTask)Start() {
    this.running = true

    for this.running {
	
    }
}

func (this *DnsTask)Stop() {
    this.running = false
}
