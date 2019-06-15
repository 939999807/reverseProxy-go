package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strings"
)

var cfg Config

func main() {
	director := func(req *http.Request) {
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
		req.URL.Host = fmt.Sprintf("%s:%s", domain, port)
	}
	proxy := &httputil.ReverseProxy{Director: director}
	err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), proxy)
	if err != nil {
		log.Panic(err)
		return
	}
	log.Println("Listening to {}", cfg.Port)
}
