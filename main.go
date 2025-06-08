package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"time"
)

type assetsResponse struct {
	Assets    []assetData `json:"data"`
	TimeStamp int64       `json:"timestamp"`
}

type assetResponse struct {
	Asset     assetData `json:"data"`
	TimeStamp int64     `json:"timestamp"`
}

type assetData struct {
	ID           string `json:"id"`
	Rank         string `json:"rank"`
	Symbol       string `json:"symbol"`
	Name         string `json:"name"`
	Supply       string `json:"supply"`
	MaxSupply    string `json:"maxSupply"`
	MarketCapUSD string `json:"marketCapUSD"`
	VolumeUSD24h string `json:"volumeUsd24Hr"`
	PriceUSD     string `json:"priceUsd"`
}

func (d assetData) Info() string {
	return fmt.Sprintf("[ID] %s  [RANK] %s", d.ID, d.Rank)
}

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (rt *loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	fmt.Fprintf(rt.logger, "%s %s %s\n", time.Now().Format(time.ANSIC), req.Method, req.URL)
	return rt.next.RoundTrip(req)
}

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	//jar.SetCookies(url.URL{
	//	Host: "",
	//
	//}("http://localhost:8000"),[]*http.Cookie{})
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println(req.Response.StatusCode)
			fmt.Println("REDIRECT")
			return nil
		},
		Transport: &loggingRoundTripper{
			logger: os.Stdout,
			next:   http.DefaultTransport,
		},
		Jar: jar,
	}
	resp, err := client.Get("https://api.coincap.io/v2/assets")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r assetResponse
	if err = json.Unmarshal(body, &r); err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	//for _, d := range r.Asset {
	//	fmt.Println(d.Info())
	//}
}
