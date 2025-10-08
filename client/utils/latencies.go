package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// @GET
func GetLatency(client *http.Client, nRequest int, api string) ([]float64, float64, float64, float64, float64) {
	// send n request to api
	// var avg float64
	var latencies []float64
	total := 0.0
	lost := 0
	max := 0.0
	min := 999999999999999999.0

	//
	ok := 0

	for i := 0; i < nRequest; i++ {
		// mark start time
		start := time.Now()

		resp, err := client.Get(api)

		// lost packet
		if err != nil {
			lost += 1
			latencies = append(latencies, 0)
			continue
		}
		// close
		resp.Body.Close()

		// take time
		duration := time.Since(start).Seconds()

		latencies = append(latencies, duration)

		total += duration
		ok++

		if duration > max {
			max = duration
		}
		if duration < min {
			min = duration
		}
	}

	// latencies, lost, avg, max, min
	return latencies, float64(lost) / float64(nRequest), total / float64(ok), max, min
}

// Product
type Product struct {
	id    int     `json:"id"`
	name  string  `json:"name"`
	price float64 `json:"price"`
}

var products []Product

func init() {
	for i := 1; i < 1000000; i++ {
		p := Product{
			name:  fmt.Sprintf("Product %d", i),
			price: float64(i * 1000),
		}
		products = append(products, p)
	}
}

// @POST
func PostLatency(client *http.Client, nRequest int, api string, nProduct ...int) ([]float64, float64, float64, float64, float64) {

	var latencies []float64
	total := 0.0
	lost := 0
	max := 0.0
	min := 999999999999999999.0

	// success request
	ok := 0

	for i := 0; i < nRequest; i++ {

		// prepare data
		amount := nProduct[0]
		if i < len(nProduct) {
			amount = nProduct[i]
		}
		dataJSON, err := json.Marshal(products[:amount])
		if err != nil {
			latencies = append(latencies, 0)
			continue
		}

		// config request
		request, err := http.NewRequest("POST", api, bytes.NewBuffer(dataJSON))
		if err != nil {
			latencies = append(latencies, 0)
			continue
		}
		request.Header.Add("Content-Type", "application/json")

		// post
		start := time.Now()
		resp, err := client.Do(request)
		if err != nil {
			lost++
			continue
		}
		resp.Body.Close()

		// take time
		duration := time.Since(start).Seconds()

		latencies = append(latencies, duration)

		ok++
		total += duration

		if duration > max {
			max = duration
		}
		if duration < min {
			min = duration
		}
	}

	return latencies, float64(lost) / float64(nRequest), total / float64(ok), max, min
}
