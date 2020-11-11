package main

import (
	"log"

	//chat klasörünü kullanacağımız için
	//go mod init üzerinde belirlediğimiz deponun
	//sonuna /chat ekliyoruz.
	"github.com/ksckaan1/grpcExample/chat"
	"golang.org/x/net/context"

	//grpc kütüphanemiz
	"google.golang.org/grpc"
)

func main() {

	//grpc client bağlantı nesnesi
	var conn *grpc.ClientConn

	//grpc ye 9080 portunu dinlemesini söyleyelim
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("9080 portu dinlenirken hata oluştu: %v", err)
	}

	//bağlantığı kapatmayı unutmayalım
	defer conn.Close()

	//Client nesnesi oluşturalım
	c := chat.NewChatServiceClient(conn)

	//server'a gönderilecek mesajı belirleyelim
	message := chat.Message{
		Body: "Client'ten merhabalar!",
	}

	//Server'dan dönen cevabı alalım
	response, err := c.SayHello(context.Background(), &message)

	//hata kontrolü yapalım.
	if err != nil {
		log.Fatalf("SayHello fonksiyonu çağırılıken hata oluştu: %v", err)
	}

	//server'dan gelen mesajı ekrana bastıralım
	log.Printf("Server'dan gelen cevap: %s", response.Body)
}
