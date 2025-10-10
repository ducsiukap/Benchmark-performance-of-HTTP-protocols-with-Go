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

	var wg sync.WaitGroup
user_select:
	fmt.Println("Select:\n[1]: Performance of protocol HTTP/1.1\n[2]: Concurrency test\n[0]: All")
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
	// [[0 0.4211929879999999 1.6203766 0.4003979]
	//  [0 1.6282791333333335 1.7673731 1.467451]
	//   [0 0.40578521 0.4083994 0.4014611]]
}
