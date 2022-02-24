package monitor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func MonitorWeb(url string, basicUser string, basicPassword string) (status string, statusCode int, body string, proto string, elapsed int64) {
	// var status string
	// var statusCode int
	// var proto string
	// var elapsed time.Duration
	t := time.Now()
	c := http.Client{
		Timeout: time.Second * 2,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	if basicUser == "" && basicPassword == "" {
		req.SetBasicAuth(basicUser, basicPassword)
	}
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		defer resp.Body.Close()
		status = resp.Status
		statusCode = resp.StatusCode
		proto = resp.Proto
	} else {
		status = ""
		statusCode = -1
		proto = ""
	}
	var bodyStr string
	if statusCode != -1 {
		body, _ := ioutil.ReadAll(resp.Body)
		bodyStr = string(body)
	} else {
		bodyStr = ""
	}
	elapsedTime := time.Since(t) / 1e6
	return status, statusCode, bodyStr, proto, int64(elapsedTime)
}
