package blockchain

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

// Blockchain Address API response
type APIResponse struct {
	Address      string `json:"address"`
	FinalBalance int64  `json:"final_balance"`
	Txs          []struct {
		Hash string `json:"hash"`
		Time int64  `json:"time"`
	} `json:"txs"`
}

// GetBlockchainAddress calls the Blockchain API to retrieve address info
func GetBlockchainAddress(address string) (*APIResponse, error) {
	url := fmt.Sprintf("https://blockchain.info/rawaddr/%s?limit=5", address)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch address info: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Parse the JSON response
	apiResp := APIResponse{}
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &apiResp, nil
}

func ValidateBitcoinAddress(address string) bool {
	_, err := btcutil.DecodeAddress(address, &chaincfg.MainNetParams)
	return err == nil
}
