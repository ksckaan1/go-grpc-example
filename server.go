package main

import (
	"log"
	"net"

	"github.com/ksckaan1/grpcExample/chat"
	"google.golang.org/grpc"
)

func main() {

	//tcp dinleyelim
	lis, err := net.Listen("tcp", ":9080")

	//hata kontrolü yapalım
	if err != nil {
		log.Fatalf("9080 portu dinlenirken hata oluştu: %v \n", err)
	}

	//chat server nesnesi oluşturalım
	s := chat.Server{}

	//grpc server'ı oluşturalım.
	grpcServer := grpc.NewServer()

	//tcp sunucu ile grpc'yi bağlayalım.
	chat.RegisterChatServiceServer(grpcServer, &s)

	//hata kontrolü
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpcServer oluşturulurken hata : %v", err)
	}
}
