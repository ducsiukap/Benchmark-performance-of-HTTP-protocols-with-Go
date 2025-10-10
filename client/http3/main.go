package main

import (
	"client/utils"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

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

	var wg sync.WaitGroup
user_select:
	fmt.Println("Select:\n[1]: Performance of protocol HTTP/3\n[2]: Concurrency test\n[0]: All")
	var choice string
	// fmt.Scan(&choice)
	choice = os.Args[1]
	fmt.Println("your choice:", choice)

	switch choice {
	case "0":
		stat := utils.SimplePerformanceTest(client)
		times := utils.ConcurrencyTest(client, &wg)

		fmt.Println("simple performance test:")
		fmt.Println("lost\t avg\t max\t min")
		fmt.Println(stat)
		fmt.Println()
		fmt.Println("concurrency test:")
		fmt.Println("time")
		fmt.Println(times)

	case "1":
		stat := utils.SimplePerformanceTest(client)

		fmt.Println("simple performance test:")
		fmt.Println("lost\t avg\t max\t min")
		fmt.Println(stat)
	case "2":
		times := utils.ConcurrencyTest(client, &wg)

		wg.Wait()

		fmt.Println("concurrency test:")
		fmt.Println("time")
		fmt.Println(times)

	default:
		fmt.Println("invalid choice:", choice)
		goto user_select
	}
}
