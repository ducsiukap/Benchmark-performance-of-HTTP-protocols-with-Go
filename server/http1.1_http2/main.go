package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"server/handler"
)

func main() {
	// route
	route := http.NewServeMux()
	route.HandleFunc("/api/product_small", handler.ProductSmallHandler)
	route.HandleFunc("/api/product_large", handler.ProductLargeHandler)

	// tls certificates
	tlsConf := tls.Config{
		NextProtos: []string{"h2", "http/1.1"},
		MinVersion: tls.VersionTLS13,
	}
	// load certificate
	cert, err := tls.LoadX509KeyPair("../tls_certificates/server.crt", "../tls_certificates/server.key")
	if err != nil {
		log.Fatalf("TLS config failed: %v", err)
	}
	// add certificate into tlsCof
	tlsConf.Certificates = []tls.Certificate{cert}

	// server
	server := http.Server{
		Addr:      ":1303",
		TLSConfig: &tlsConf,
		Handler:   route,
	}

	// listening
	fmt.Println("server is running on", server.Addr)
	err = server.ListenAndServeTLS("", "")

	if err != nil {
		fmt.Println("server error")
	}
}
