package http_request

import (
	"io/ioutil"
	"errors"
	"net/http"
	"net"
	// "strconv"
	"time"
	"bytes"
)

func Send(message []byte, url, endpoint string) ([]byte, error) {
	// timeoutDuration, err := time.ParseDuration(strconv.FormatUint(uint64(timeout), 10) + "s")
	// if err != nil {
	// 	return nil, err
	// }
	timeoutDuration := time.Second * 3
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeoutDuration,
			KeepAlive: timeoutDuration,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       timeoutDuration,
		TLSHandshakeTimeout:   timeoutDuration,
		ExpectContinueTimeout: timeoutDuration,
	}
	client := &http.Client{
		Timeout:   timeoutDuration,
		Transport: transport,
	}
	buffer := bytes.NewBuffer([]byte(message))
	response, err := client.Post(url + endpoint, "application/json", buffer)
	if err != nil {
		return nil, err
	}

	respBodyRaw, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if respBodyRaw == nil {
		return nil, errors.New("Тело ответа пусто (равно nil)")
	}
	return respBodyRaw, nil
}
