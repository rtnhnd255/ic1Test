package server

import (
	"log"
	"net"
	"tcp_server/app/internal/config"
	"tcp_server/app/internal/db"
	"tcp_server/app/internal/parser"
)

type Server struct {
	Cfg    config.Config
	DB     *db.Db
	Parser *parser.Parser
}

func NewServer(cfgPath string) (*Server, error) {
	cfg, err := config.ParseConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	return &Server{
		Cfg:    *cfg,
		DB:     db.NewDBConn(cfg),
		Parser: parser.NewParser(),
	}, nil
}

func (s *Server) RunServer(port string) error {
	log.Println("Starting server")
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}
	conn, err := ln.Accept()
	if err != nil {
		return err
	}
}
