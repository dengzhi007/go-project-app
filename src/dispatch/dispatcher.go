package dispatch

import(
    "common"
    "log"
    "fmt"
    "os"
    "time"
)


type Dispatcher struct {
    logger      *log.Logger
    running     bool
    heartbeat   int64
    alive_limit int64
}

func NewDispatcher() *Dispatcher {
    d := &Dispatcher{
	running: false,
	heartbeat: time.Now().Unix(),
	alive_limit: 60*60,
    }

    logFileName := common.SysPath + "/logs/dispatcher.log"
    logFile, err := os.Create(logFileName)
    if err != nil {
	fmt.Printf("open %s failed.", logFileName)
	return nil
    }
    d.logger = log.New(logFile, "", log.LstdFlags)

    return d
}

func (this *Dispatcher)Start() {
    this.logger.Println("dispatcher start.")
    this.running = true

    for this.running {
	this.heartbeat = time.Now().Unix() 
	sleep(1)

    }

}

func (this *Dispatcher)Stop() {
    this.logger.Println("dispatcher stop.")
    this.running = false
}

func (this *Dispatcher)IsAlive() bool {
    if this.heartbeat + this.alive_limit > time.Now().Unix() {
	return true
    } 
    return false
}
