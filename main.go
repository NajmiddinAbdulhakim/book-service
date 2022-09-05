package main

import (
	"log"
	"net"

	pb "github.com/NajmiddinAbdulhakim/ude/book-service/genproto"
	"google.golang.org/grpc"

	"github.com/NajmiddinAbdulhakim/ude/book-service/config"
	"github.com/NajmiddinAbdulhakim/ude/book-service/db"
	"github.com/NajmiddinAbdulhakim/ude/book-service/service"
)

func main() {
	cfg := config.Load()

	conn, err := db.ConnectDB(cfg)
	if err != nil {
		log.Fatal(`Failed to connection postgeres:`, err)
	}

	service := service.NewBookService(conn)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal(`Error while listening: `, err)
	}
	s := grpc.NewServer()
	pb.RegisterBookServcieServer(s, service)
	if err := s.Serve(lis); err != nil {
		log.Fatal(`Error while serve: `, err)
	}
}
