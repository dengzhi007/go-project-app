package main

import (
    //basic
    "fmt"
    "time"
    "flag"
    "os"
    "os/signal"
    "syscall"
    "runtime"

    //third-party lib
    "lib/go-config/config"
    
    //project package
    "common"
    "network"
    "dispatch"
)

const (
    VERSION       string = "v1.0"
    MAX_PROCESSOR int = 5
)

func initialize() error {
    fmt.Println("initilize ...")

    confFile := flag.String("c", "main.conf", "config file name")
    flag.Parse()
    if confFile == nil {
	return fmt.Errorf("no config file")
    }

    fmt.Printf("load configuration:%s\n", *confFile)
    conf, err := config.ReadDefault(*confFile)
    if err != nil {
	return fmt.Errorf("load config file failed:%s\n", err.Error())
    }
    common.Conf = conf

    common.SysPath, _ = os.Getwd()
    fmt.Printf("get working directory:%s\n", common.SysPath)

    max_processor, err := common.Conf.Int("global", "max_processor")
    if err != nil {
	max_processor = MAX_PROCESSOR
    }
    runtime.GOMAXPROCS(max_processor)
    fmt.Printf("set runtime max processors:%d\n", max_processor)

    fmt.Printf("Program %s start success in %s at: %s\n", VERSION, common.SysPath, time.Now())

    return nil
}


func main() {
    err := initialize()
    if err != nil {
	panic(fmt.Sprintf("initialize failed:%s\n", err.Error()))
    }

    //start to work...
    server := network.NewTcpServer()
    if server == nil {
	panic("new tcp server failed.")
    }
    go server.Start()

    dispatcher := dispatch.NewDispatcher()
    if dispatcher == nil {
	panic("new dispatcher failed.")
    }
    go dispatcher.Start() 

    updater := dispatch.NewUpdater()
    if updater == nil {
	panic("new updater failed.")
    }
    go updater.Start()

    //wait for signal
    sig_chan := make(chan os.Signal)
    signal.Notify(sig_chan, os.Interrupt)
    signal.Notify(sig_chan, syscall.SIGTERM)
    <-sig_chan

    //clear works...
    server.Stop()
    dispatcher.Stop()
    updater.Stop()

    time.Sleep(1 * time.Second)
    fmt.Printf("Program %s quit success at: %s\n", VERSION, time.Now())
}

