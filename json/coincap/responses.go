package coincap

import "fmt"

type AssetsResponse struct {
	Assets    []Asset `json:"data"`
	TimeStamp int64   `json:"timestamp"`
}

type AssetResponse struct {
	Asset     Asset `json:"data"`
	TimeStamp int64 `json:"timestamp"`
}

type Asset struct {
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

func (d Asset) Info() string {
	return fmt.Sprintf("[ID] %s  [RANK] %s", d.ID, d.Rank)
}
