package main

import (
	"flag"
	"log"
	"tcp_server/app/internal/server"
)

func main() {
	var port string
	var cfgPath string
	flag.StringVar(&port, "port", "20163", "Specify listen port")
	flag.StringVar(&cfgPath, "config", "./defaultConfig.yaml", "Specify path to config file")
	flag.Parse()

	s, err := server.NewServer(cfgPath, port)
	if err != nil {
		log.Fatalln(err)
	}
	err = s.RunServer()
	if err != nil {
		log.Fatalln(err)
	}
}
