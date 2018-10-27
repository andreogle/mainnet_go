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
	defer res.Body.Close()

	if err != nil {
		fmt.Printf("Failed to GET %v", url)
	}

	return res, err
}
