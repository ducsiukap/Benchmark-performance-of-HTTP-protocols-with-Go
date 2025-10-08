package main

import (
	"client/utils"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
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
	}
	// transport
	transport := &http.Transport{
		TLSClientConfig:   tlsConf,
		ForceAttemptHTTP2: false, // disable http2
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
	// [[0 0.4211929879999999 1.6203766 0.4003979]
	//  [0 1.6282791333333335 1.7673731 1.467451]
	//   [0 0.40578521 0.4083994 0.4014611]]
}
