package dispatch

import(
//    "common"
    "net/http"
    "log"
)

type Updater struct {
    url     string
    freq    int 
    running bool
    logger  *log.Logger
}

func NewUpdater() *Updater {
    u := &Updater{
	url: "http://10.16.77.80.8360/index.php",
	freq: 60 * 10,
	running: false,
    }
    return u
}

func (this *Updater)Start() {
    this.running = true

    for this.running {
	resp, err := http.Get(this.url)
	if err != nil {
	    continue
	
	
	}
	if resp.StatusCode != 200 {
	
	}
    }

}

func (this *Updater)Stop() {
    this.running = false
}

func (this *Updater)IsAlive() bool {
    return true
}
