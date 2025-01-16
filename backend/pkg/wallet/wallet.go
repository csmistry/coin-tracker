package wallet

import (
	"github.com/csmistry/coin-tracker/backend/pkg/address"
	"github.com/csmistry/coin-tracker/backend/pkg/transaction"
)

// Wallet stores address IDs in-memory
type Wallet struct {
	Addresses []*address.Address
}

func NewWallet() *Wallet {
	return &Wallet{
		Addresses: []*address.Address{},
	}
}

var BitcoinWallet *Wallet

// Initialize new in-memory wallet
func Init() {
	BitcoinWallet = NewWallet()
}

// WallerInterface defines actions on a wallet
type WalletInterface interface {
	ListAddresses() []string
	GetAddress(id string) *address.Address
	AddAddress(id string) *address.Address
	RemoveAddress(id string)
}

func (w *Wallet) ListAddresses() []string {
	res := []string{}
	for _, addr := range w.Addresses {
		if !addr.IsArchived {
			res = append(res, addr.Id)
		}
	}

	return res
}

func (w *Wallet) GetAddress(id string) *address.Address {
	// use Blockchain API to get address details
	return nil
}

func (w *Wallet) AddAddress(id string) error {
	//Validate with Blockchain API if address is valid

	newAddr := &address.Address{
		Id:           id,
		Transactions: []transaction.Transaction{},
	}

	w.Addresses = append(w.Addresses, newAddr)
	return nil
}

// RemoveAddress archives an address in the wallet
func (w *Wallet) RemoveAddress(id string) {
	for _, address := range w.Addresses {
		if address.Id == id {
			address.IsArchived = true
			break
		}
	}
}

// CheckAddressExists is a helper function that checks if an address exists in the wallet
func CheckAddressExists(id string) bool {
	for _, address := range BitcoinWallet.Addresses {
		if address.Id == id && !address.IsArchived {
			return true
		}
	}
	return false
}
