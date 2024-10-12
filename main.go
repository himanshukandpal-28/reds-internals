package main

import (
	"flag"
	"log"
	"reds-internals/config"
	"reds-internals/server"
)

func setupFlags() {

	flag.StringVar(&config.Host, "host", "0.0.0.0", "host for the dice db server")
	flag.IntVar(&config.Port, "port", 7379, "port for the dice db server")
	flag.Parse()
}

func main() {
	setupFlags()
	log.Println("rolling dice server :) ")
	server.RunSyncTcpServer()
}
