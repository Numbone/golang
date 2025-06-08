package coincap

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Client struct {
	client *http.Client
}

func NewClient(timeout time.Duration) (*Client, error) {
	if timeout == 0 {
		return nil, errors.New("timeout")
	}
	return &Client{
		client: &http.Client{
			Timeout: timeout,
			Transport: &LoggingRoundTripper{
				logger: os.Stdout,
				next:   http.DefaultTransport,
			},
		},
	}, nil
}

func (c Client) GetAssets() ([]Asset, error) {

	resp, err := c.client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r AssetsResponse
	if err = json.Unmarshal(body, &r); err != nil {
		log.Fatal(err)
	}
	return r.Assets, nil
}
