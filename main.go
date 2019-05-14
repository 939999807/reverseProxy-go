package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

var cfg Config

func main() {
	cfg = parseConfig()

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Panic(err)
	}

	for {
		client, err := l.Accept()
		if err != nil {
			log.Panic(err)
		}

		go handleClientRequest(client)
	}
}

func handleClientRequest(client net.Conn) {
	if client == nil {
		return
	}
	defer client.Close()

	var b [1024]byte
	n, err := client.Read(b[:])
	if err != nil {
		log.Println(err)
		return
	}

	reader := bufio.NewReader(strings.NewReader(string(b[:])))
	req, err := http.ReadRequest(reader)
	if err != nil {
		log.Println(err)
	}

	var domain string
	i := strings.IndexByte(req.Host, ':')
	if i < 0 {
		domain = req.Host
	} else {
		domain = req.Host[0:i]
	}
	port, ok := cfg.Route[domain];
	if !ok {
		log.Println(domain + " not found")
		return
	}
	//获得了请求的host和port，就开始拨号吧
	server, err := net.Dial("tcp", domain+":"+port)
	if err != nil {
		log.Println(err)
		return
	}
	server.Write(b[:n])
	//进行转发
	go io.Copy(server, client)
	io.Copy(client, server)
}
