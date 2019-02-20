package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

func now() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func httpget(wg *sync.WaitGroup, id string, config *Config) {
	defer wg.Done()
	client := &http.Client{}

	for i := 0; i < config.Iterations; i++ {
		req, _ := http.NewRequest("GET", config.URL, nil)

		// Add the APIKey if there is one
		if len(config.APIKey) > 0 {
			req.Header.Set("Authorization", "Apikey "+config.APIKey)
		}
		// Add additional headers
		for h := range config.Headers {
			header := strings.SplitN(config.Headers[h], ":", 2)
			req.Header.Set(header[0], header[1])
		}

		start := now()
		res, err := client.Do(req)
		end := now()

		if err != nil {
			fmt.Printf("%s-%d: (%d) %s\n", id, i, end-start, err)
		} else {
			status := res.StatusCode
			reqID := res.Header.Get("request-id")
			// if len(reqID) == 0 {
			// 	bodyBytes, _ := ioutil.ReadAll(res.Body)
			// 	bodyString := string(bodyBytes)
			// 	fmt.Printf("%+v\n", bodyString)
			// }
			fmt.Printf("%s-%d: (%d) %d %s\n", id, i, end-start, status, reqID)
		}
	}
}

func perf(config *Config) {
	var wg sync.WaitGroup
	wg.Add(config.Threads)

	for i := 0; i < config.Threads; i++ {
		go httpget(&wg, "t"+strconv.Itoa(i), config)
	}

	wg.Wait()
}
