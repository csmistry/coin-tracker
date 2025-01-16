package wallet

import (
	"fmt"

	"github.com/csmistry/coin-tracker/backend/pkg/blockchain"
)

// Wallet represents a crypto that holds addresses and their txs
type Wallet struct {
	addresses map[string]*Address
}

// Transaction for a given address
type Transaction struct {
	Id        string `json:"hash"`
	Time      string `json:"time"`
	AddressID string `json:"address"`
}

// Address represents a btc wallet address
type Address struct {
	Id           string         `json:"id"`
	Balance      float64        `json:"balance"`
	Transactions []*Transaction `json:"transactions"`
	IsArchived   bool           `json:"is_archived"`
}

func NewWallet() *Wallet {
	return &Wallet{
		addresses: map[string]*Address{},
	}
}

// BitcoinWallet is in-memory wallet managed by the server
var BitcoinWallet *Wallet

// Initialize new in-memory wallet
func Init() {
	BitcoinWallet = NewWallet()
}

// WallerInterface defines actions on a wallet
type WalletInterface interface {
	ListAddresses() []string
	GetAddress(id string) (*Address, error)
	AddAddress(id string) *Address
	RemoveAddress(id string)
}

// ListAddresses returns a list of address IDs
func (w *Wallet) ListAddresses() []string {
	res := []string{}
	for addrId, addr := range w.addresses {
		if !addr.IsArchived {
			res = append(res, addrId)
		}
	}

	return res
}

// GetAddress returns the address details for provided address ID
func (w *Wallet) GetAddress(id string) *Address {
	return w.addresses[id]
}

// AddAddress adds a valid btc address to in-memory wallet
func (w *Wallet) AddAddress(id string) error {
	// validate address
	valid := blockchain.ValidateBitcoinAddress(id)
	if !valid {
		return fmt.Errorf("invalid btc address")
	}

	// fetch details from blockchain
	resp, err := blockchain.GetBlockchainAddress(id)
	if err != nil {
		return fmt.Errorf("failed to address details for ID: [%s] err: %w", id, err)
	}

	// update in-memory wallet
	w.addresses[id] = APIRespnseToAddress(resp)
	fmt.Println("Address added with ID: ", id)
	return nil
}

// RemoveAddress archives an address in the wallet
func (w *Wallet) RemoveAddress(id string) {
	for addrId, address := range w.addresses {
		if addrId == id {
			address.IsArchived = true
			break
		}
	}
	fmt.Println("Archived address: ", id)
}
