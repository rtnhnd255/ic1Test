package server

import (
	"log"
	"net"
	"tcp_server/app/internal/config"
	"tcp_server/app/internal/db"
	"tcp_server/app/internal/parser"
)

type Server struct {
	config config.Config
	db     *db.Db
	parser *parser.Parser
	port   string
}

func NewServer(cfgPath string, port string) (*Server, error) {
	cfg, err := config.ParseConfig(cfgPath)
	if err != nil {
		return nil, err
	}
	return &Server{
		config: *cfg,
		db:     db.NewDBConn(cfg),
		parser: parser.NewParser(),
		port:   port,
	}, nil
}

func (s *Server) RunServer() error {
	log.Println("Starting server")
	ln, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		s.handleMessage(conn)
	}
}

//TODO: specify package size

func (s *Server) handleMessage(conn net.Conn) {
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)

	msg := s.parser.ParseMsg(buf)
	err = s.db.CreateRecord(msg)
	if err != nil {
		log.Println(err)
	}
}
