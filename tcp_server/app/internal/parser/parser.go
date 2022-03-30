package parser

import "tcp_server/app/internal/model"

// TODO: everything

type Parser struct {
}

func NewParser() *Parser {
	return nil
}

func (p *Parser) ParseMsg(msg []byte) model.RecordDTO {
	return model.RecordDTO{}
}
