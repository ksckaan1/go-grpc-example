package chat

import (
	"log"

	"golang.org/x/net/context"
)

type Server struct {
	//Burası şuanlık boş olacak
}

//server nesnemize iliştirilecek fonksiyonumuz
func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {

	//Client'ten gelen mesajı ekrana bastıralım
	log.Printf("Clientten mesaj alındı: %s", message.Body)

	//mesaj gelince değer döndürelim (mesaj ve hata)
	return &Message{Body: "Server'dan merhaba!"}, nil //hata nil olsun
}
