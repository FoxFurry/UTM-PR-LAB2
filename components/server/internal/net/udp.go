package net

import (
	"context"
	"fmt"
	"github.com/foxfurry/university/PRLab2/components/server/internal/infrastructure/word"
	"github.com/spf13/viper"
	"log"
	"net"
	"time"
)

type UDPServer struct {

}

func NewUDP() Server {
	return &UDPServer{}
}

func (s *UDPServer) Listen(ctx context.Context) {
	p := make([]byte, 2048)
	addr := net.UDPAddr{
		Port: viper.GetInt("host.udp"),
		IP: net.ParseIP("127.0.0.1"),
	}
	ser, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Panicf("[UDP] Failed to listen: %v", err)
	}else{
		log.Printf("[UDP] Listening to %v", addr.String())
	}

	readerTick := time.Tick(time.Millisecond * 10)

	for {
		select {
		case <-readerTick:
			_,remoteaddr,err := ser.ReadFromUDP(p)
			fmt.Printf("[UDP] Read a message from %v %s \n", remoteaddr, p)
			if err !=  nil {
				log.Fatalf("Read from remote address failed: %v", err)
			}
			go sendResponse(ser, remoteaddr, word.Say())

		case <-ctx.Done():
			log.Printf("[UDP] Closing listener...")
			return
		}
	}
}

func (s *UDPServer) Shutdown() error {
	return nil
}

func sendResponse(conn *net.UDPConn, addr *net.UDPAddr, message string) {
	_,err := conn.WriteToUDP([]byte(message), addr)
	if err != nil {
		fmt.Printf("[UDP] Couldn't send response %v", err)
	}
}