package utils

import (
	"encoding/json"
	"net/http"
	"time"
)

var netClient = http.Client{
	Timeout: time.Second * 5,
}

func HttpGet(url string, target interface{}) error {
	res, err := netClient.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
