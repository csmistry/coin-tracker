package wallet

// Wallet represents a crypto that holds addresses and their txs
type Wallet struct {
	addresses []*Address
}

// Transaction for a given address
type Transaction struct {
	Id        string `json:"hash"`
	Time      int64  `json:"time"`
	AddressID string `json:"address"`
}

// Address represents a btc wallet address
type Address struct {
	Id           string         `json:"id"`
	Balance      float64        `json:"balance"`
	Transactions []*Transaction `json:"transactions"`
	IsArchived   bool
}

func NewWallet() *Wallet {
	return &Wallet{
		addresses: []*Address{},
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
	GetAddress(id string) *Address
	AddAddress(id string) *Address
	RemoveAddress(id string)
}

// ListAddresses returns a list of address IDs
func (w *Wallet) ListAddresses() []string {
	res := []string{}
	for _, addr := range w.addresses {
		if !addr.IsArchived {
			res = append(res, addr.Id)
		}
	}

	return res
}

// GetAddress returns the address details for provided address ID
func (w *Wallet) GetAddress(id string) *Address {
	// use Blockchain API to get address details
	return nil
}

// AddAddress adds a valid btc address to in-memory wallet
func (w *Wallet) AddAddress(id string) error {
	//Validate with Blockchain API if address is valid

	newAddr := &Address{
		Id:           id,
		Transactions: []*Transaction{},
	}

	w.addresses = append(w.addresses, newAddr)
	return nil
}

// RemoveAddress archives an address in the wallet
func (w *Wallet) RemoveAddress(id string) {
	for _, address := range w.addresses {
		if address.Id == id {
			address.IsArchived = true
			break
		}
	}
}

// CheckAddressExists is a helper function that checks if an address exists in the wallet
func CheckAddressExists(id string) bool {
	for _, address := range BitcoinWallet.addresses {
		if address.Id == id && !address.IsArchived {
			return true
		}
	}
	return false
}
