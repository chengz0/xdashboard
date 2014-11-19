package main

import (
	"github.com/chengz0/xdashboard/global"
	"github.com/chengz0/xdashboard/http"
	"log"
)

func main() {
	log.Println("starting...")
	//config
	global.InitGlobal("cfg.ini")
	// mamrtini
	http.StartMartini()
}
