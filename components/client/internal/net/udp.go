package net

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"net"
)

func UDPSend(message string) {
	p :=  make([]byte, 2048)
	conn, err := net.Dial("udp", viper.GetString("host.udp"))
	if err != nil {
		fmt.Printf("[UDP] Could not dial UDP: %v\n", err)
		return
	}
	fmt.Fprintf(conn, message)
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("[UDP] Received response: %s\n", p)
	} else {
		fmt.Printf("[UDP] Could not read the response: %v\n", err)
	}
}
