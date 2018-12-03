package main

import (
	"log"
	"net/http"
)

var cfg Config
var srv http.Server

func StartServer(port string, remote string) {
	log.Printf("Listening on %s, forwarding to %s", port, remote)
	h := &handle{reverseProxy: remote}
	srv.Addr = ":" + port
	srv.Handler = h
	//go func() {
	err := srv.ListenAndServe();
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
	//}()
}

//func StopServer() {
//	if err := srv.Shutdown(nil); err != nil {
//		log.Println(err)
//	}
//}

func main() {
	cfg = parseConfig()
	StartServer(cfg.Port, cfg.Remote)
}
