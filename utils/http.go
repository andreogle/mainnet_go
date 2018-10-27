package utils

import (
	"fmt"
	"net/http"
	"time"
)

var netClient = http.Client{
	Timeout: time.Second * 5,
}

func HttpGet(url string) (*http.Response, error) {
	res, err := netClient.Get(url)
	if err != nil {
		fmt.Printf("Failed to GET %v", url)
	}
	defer res.Body.Close()

	return res, err
}
