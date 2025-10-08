package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"server/handler"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	// route
	route := http.NewServeMux()
	route.HandleFunc("/api/product_small", handler.ProductSmallHandler)

	// tls certificates
	tlsConf := tls.Config{
		MinVersion: tls.VersionTLS13,
	}
	// load certificate
	// cert, err := tls.LoadX509KeyPair("../tls_certificates/server.crt", "../tls_certificates/server.key")
	cert, err := tls.LoadX509KeyPair("/home/ubuntu/httpserver/server/tls_certificates/server.crt",
		"/home/ubuntu/httpserver/server/tls_certificates/server.key")

	if err != nil {
		log.Fatalf("TLS config failed: %v", err)
	}
	// add certificate into tlsCof
	tlsConf.Certificates = []tls.Certificate{cert}

	// server
	server := http3.Server{
		Addr:      ":443",
		TLSConfig: &tlsConf,
		Handler:   route,
	}

	// listening
	fmt.Println("server is running on", server.Addr)
	err = server.ListenAndServe()

	if err != nil {
		fmt.Printf("http/3 server error %v", err)
	}
}
