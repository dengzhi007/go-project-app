package common

import (
    "lib/go-config/config"
)

var SysPath string

var Test int = 1
var Conf *config.Config

var TaskMap map[int]*Task
var PQ *PQueue

var TcpTaskChan  chan *Task
var PingTaskChan chan *Task
var HttpTaskChan chan *Task
var DnsTaskChan  chan *Task
var UdpTaskChan  chan *Task


