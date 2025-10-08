package main

import (
	"client/utils"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/net/http2"
)

func main() {

	// server certificate
	certPool := x509.NewCertPool()
	svrCert, err := os.ReadFile("../certs/server.crt")
	if err != nil {
		log.Fatalf("Read search error: %v\n", err)
	}
	certPool.AppendCertsFromPEM(svrCert)
	tlsConf := &tls.Config{
		RootCAs:    certPool,
		MinVersion: tls.VersionTLS13,
		NextProtos: []string{"h2", "http/1.1"},
	}
	// transport
	transport := &http2.Transport{
		TLSClientConfig: tlsConf,

		// ForceAttemptHTTP2: false, // disable http2
	}

	// client
	client := &http.Client{
		Transport: transport,
	}

	// performance
	var stats [][4]float64
	// GET
	// Product small
	_, lost, avg, max, min := utils.GetLatency(client, 100, "https://52.62.36.23/api/product_small")
	stats = append(stats, [4]float64{lost, avg, max, min})
	// Product large
	_, lost, avg, max, min = utils.GetLatency(client, 3, "https://52.62.36.23/api/product_large")
	stats = append(stats, [4]float64{lost, avg, max, min})

	// POST
	_, lost, avg, max, min = utils.PostLatency(client, 10, "https://52.62.36.23/api/product_small", 20)
	stats = append(stats, [4]float64{lost, avg, max, min})

	fmt.Println(stats)
	// [[0 0.13391938800000003 0.4051881 0.1287043]
	//  [0 1.5628863333333334 1.7549697 1.2765632]
	//  [0 0.15959724000000003 0.4131058 0.1294764]]
}
