package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/csmistry/coin-tracker/backend/pkg/wallet"
	"github.com/gorilla/mux"
)

// ListAddresses returns a list of address added to the wallet
func ListAddresses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	addrs := wallet.BitcoinWallet.ListAddresses()
	json.NewEncoder(w).Encode(addrs)
}

// GetAddress returns address information for a specific address
func GetAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "No address ID specified", http.StatusBadRequest)
		return
	}

	exists := wallet.CheckAddressExists(id)
	if !exists {
		http.Error(w, "Invalid address ID", http.StatusBadRequest)
		return
	}

	// Fetch Address details
	address := wallet.BitcoinWallet.GetAddress(id)
	json.NewEncoder(w).Encode(address)
}

// AddAddress adds a new address to the wallet
func AddAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "No address ID specified", http.StatusBadRequest)
		return
	}

	// Wallet must not have duplicate addresses
	exists := wallet.CheckAddressExists(id)
	if exists {
		http.Error(w, "Address already exists", http.StatusConflict)
		return
	}

	// Add new address to wallet
	err := wallet.BitcoinWallet.AddAddress(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Println(w, "Address added with ID: %d", id)
}

// RemoveAddress archives an address from the wallet
func RemoveAddress(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		http.Error(w, "No address ID specified", http.StatusBadRequest)
		return
	}

	exists := wallet.CheckAddressExists(id)
	if !exists {
		http.Error(w, "Invalid address ID", http.StatusBadRequest)
		return
	}

	// archive address
	wallet.BitcoinWallet.RemoveAddress(id)
	w.WriteHeader(http.StatusOK)
}
