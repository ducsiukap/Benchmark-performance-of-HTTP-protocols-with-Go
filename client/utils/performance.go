package utils

import (
	"net/http"
	"sync"
	"time"
)

func SimplePerformanceTest(client *http.Client) [][4]float64 {

	var stats [][4]float64

	_, lost, avg, max, min := GetLatency(client, 100, "https://52.62.36.23/api/product_small")
	stats = append(stats, [4]float64{lost, avg, max, min})
	// Product large
	_, lost, avg, max, min = GetLatency(client, 3, "https://52.62.36.23/api/product_large")
	stats = append(stats, [4]float64{lost, avg, max, min})

	// POST
	_, lost, avg, max, min = PostLatency(client, 1, "https://52.62.36.23/api/product_small", 20)
	stats = append(stats, [4]float64{lost, avg, max, min})

	// for _, stat := range stats {
	// 	fmt.Println(stat)
	// }
	return stats
}

func ConcurrencyTest(client *http.Client, wg *sync.WaitGroup) []float64 {
	var times []float64

	// get Product_Small
	start := time.Now()
	wg.Add(1)
	go concurrentGet(client, 1, "https://52.62.36.23/api/product_small", wg)
	wg.Add(1)
	go concurrentGet(client, 1, "https://52.62.36.23/api/product_small", wg)
	wg.Add(1)
	go concurrentGet(client, 1, "https://52.62.36.23/api/product_small", wg)
	wg.Add(1)
	go concurrentGet(client, 1, "https://52.62.36.23/api/product_small", wg)
	wg.Add(1)
	go concurrentGet(client, 1, "https://52.62.36.23/api/product_small", wg)
	wg.Wait()
	times = append(times, float64(time.Since(start).Seconds()))

	// get Product_Large
	start = time.Now()
	wg.Add(1)
	go concurrentGet(client, 1, "https://52.62.36.23/api/product_large", wg)
	wg.Add(1)
	go concurrentGet(client, 1, "https://52.62.36.23/api/product_large", wg)
	wg.Wait()
	times = append(times, float64(time.Since(start).Seconds()))

	// Post
	start = time.Now()
	wg.Add(1)
	go concurrentPost(client, 1, "https://52.62.36.23/api/product_small", wg, 10)
	wg.Add(1)
	go concurrentPost(client, 1, "https://52.62.36.23/api/product_small", wg, 10)
	wg.Add(1)
	go concurrentPost(client, 1, "https://52.62.36.23/api/product_small", wg, 10)
	wg.Add(1)
	go concurrentPost(client, 1, "https://52.62.36.23/api/product_small", wg, 10)
	wg.Add(1)
	go concurrentPost(client, 1, "https://52.62.36.23/api/product_small", wg, 10)
	wg.Wait()
	times = append(times, float64(time.Since(start).Seconds()))

	return times
}
