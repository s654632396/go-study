package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	log.Println(get_local_ip())
}

func get_local_ip() string {
	conn, _ := net.Dial("udp", "8.8.8.8:80")
	defer conn.Close()
	ip := conn.LocalAddr().(*net.UDPAddr).IP
	log.Println(ip)
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}
