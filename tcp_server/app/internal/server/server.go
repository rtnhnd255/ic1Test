package server

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"net"
	"tcp_server/app/internal/config"
	"tcp_server/app/internal/db"
	"tcp_server/app/internal/model"
)

type Server struct {
	config config.Config
	db     *db.Db
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
	var buf []byte
	_, err := conn.Read(buf)

	record := parsePackage(buf)
	err = s.db.CreateRecord(*record)
	if err != nil {
		log.Println(err)
	}
}

func parsePackage(msg []byte) *model.RecordDTO {
	var result model.RecordDTO
	offset := 0
	if len(msg) < 4 {
		log.Println("Corrupted package")
		return nil
	}

	packageSize := binary.LittleEndian.Uint16(msg[0:4])
	if (uint16)(len(msg)-4) < packageSize {
		log.Println("Corrupted package")
		return nil
	}

	message := msg[4:packageSize]
	var uidEndIndex int
	for i, b := range message {
		if b == 0 {
			uidEndIndex = i
		}
	}
	result.DeviceID = fmt.Sprint(message[0:uidEndIndex])
	offset = uidEndIndex + 1

	result.PointTime = binary.LittleEndian.Uint16(message[offset : offset+4])
	offset += 8

	for {
		packageSize = binary.BigEndian.Uint16(message[offset+2 : offset+4])
		if binary.BigEndian.Uint16(message[offset:offset+2]) != 0x0BBB {
			log.Println("Corrupted block")
			offset += int(packageSize)
			continue
		}

		packageBuf := message[offset+4 : offset+int(packageSize)]
		flag, lon, lat := parsePackageBlock(packageBuf)
		if !flag {
			continue
		} else {
			result.Latitude = lat
			result.Longitude = lon
			break
		}
	}

	return &result
}

func parsePackageBlock(pkg []byte) (flag bool, lon float64, lat float64) {
	var nameEndIndex int
	var offset = 2
	for i, b := range pkg {
		if b == 0 {
			nameEndIndex = i
		}
	}
	pkgName := fmt.Sprint(pkg[offset:nameEndIndex])
	if pkgName != "posinfo" {
		return false, 0, 0
	}
	offset = nameEndIndex + 1

	lon = float64FromBytes(pkg[offset : offset+8])
	offset += 8
	lat = float64FromBytes(pkg[offset : offset+8])

	return true, lon, lat
}

func float64FromBytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
