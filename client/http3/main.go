package main

import (
	"client/utils"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/quic-go/quic-go/http3"
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
	transport := &http3.Transport{
		TLSClientConfig: tlsConf,

		// ForceAttemptHTTP2: false, // disable http2
	}
	defer transport.Close()

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
	// [[0 0.13355490599999997 0.278496 0.1302149]
	// [0 0.1298176666666667 0.1301002 0.1294462]
	//  [0 0.13162401 0.1341455 0.130302]]
}
