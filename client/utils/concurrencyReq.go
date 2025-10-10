package utils

import (
	"net/http"
	"sync"
)

func concurrentGet(client *http.Client, nRequest int, api string, wg *sync.WaitGroup) {
	defer wg.Done()
	GetLatency(client, nRequest, api)
}

func concurrentPost(client *http.Client, nRequest int, api string, wg *sync.WaitGroup, nProduct ...int) {
	defer wg.Done()
	PostLatency(client, nRequest, api, nProduct...)
}
