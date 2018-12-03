package main

import (
	"context"
	"github.com/bogdanovich/dns_resolver"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"
)

type handle struct {
	reverseProxy string
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
	remote, err := url.Parse(h.reverseProxy)
	if err != nil {
		log.Fatalln(err)
	}
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}
	http.DefaultTransport.(*http.Transport).DialContext = func(ctx context.Context, network, addr string) (net.Conn, error) {
		remote := strings.Split(addr, ":")
		if cfg.Ip == "" {
			resolver := dns_resolver.New(cfg.Dns)
			resolver.RetryTimes = 5
			ip, err := resolver.LookupHost(remote[0])
			if err != nil {
				log.Println(err)
			}
			cfg.Ip = ip[0].String()
		}
		addr = cfg.Ip + ":" + remote[1]
		return dialer.DialContext(ctx, network, addr)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	r.Host = remote.Host
	proxy.ServeHTTP(w, r)
}
