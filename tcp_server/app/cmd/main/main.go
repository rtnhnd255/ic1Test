package main

import (
	"flag"
	"github.com/rtnhnd255/ic1Test/app/internal/server"
	"log"

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
