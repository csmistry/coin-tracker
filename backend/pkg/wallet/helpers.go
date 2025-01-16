package wallet

import (
	"time"

	"github.com/csmistry/coin-tracker/backend/pkg/blockchain"
)

// CheckAddressExists is a helper function that checks if an address exists in the wallet
func CheckAddressExists(id string) bool {
	for addrId, address := range BitcoinWallet.addresses {
		if addrId == id && !address.IsArchived {
			return true
		}
	}
	return false
}

// APIRespnseToAddress extracts required fields from blockchain API response
func APIRespnseToAddress(resp *blockchain.APIResponse) *Address {
	addr := &Address{
		Id:           resp.Address,
		Balance:      float64(resp.FinalBalance) / 1e8, // represent value as BTC
		Transactions: []*Transaction{},
	}

	// add transactions
	for _, tx := range resp.Txs {
		// convert unix time to readable format
		t := time.Unix(tx.Time, 0)
		formattedTime := t.Format("2006-01-02 15:04:05 MST")

		addr.Transactions = append(addr.Transactions, &Transaction{
			Id:        tx.Hash,
			Time:      formattedTime,
			AddressID: resp.Address,
		})
	}
	return addr
}
