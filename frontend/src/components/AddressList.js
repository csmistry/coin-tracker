// src/components/AddressList.js
import React, { useEffect, useState } from "react";
import { fetchAddresses, fetchAddressById, deleteAddress} from "../api/walletService";

const AddressList = () => {
  const [addresses, setAddresses] = useState([]);
  const [loading, setLoading] = useState(true);
  const [selectedAddress, setSelectedAddress] = useState(null);
  const [selectedAddressBalance, setSelectedAddressBalance] = useState(null);
  const [transactions, setTransactions] = useState([]);

  useEffect(() => {
    const loadAddresses = async () => {
      try {
        const data = await fetchAddresses();
        setAddresses(data);
      } catch (error) {
        console.error("Failed to load addresses.");
      } finally {
        setLoading(false);
      }
    };

    loadAddresses();
  }, []);

  const handleClick = async (id) => {
    try {
      setSelectedAddress(id);
      const data = await fetchAddressById(id);
      setTransactions(data.transactions);
      setSelectedAddressBalance(data.balance)
    } catch (error) {
      console.error(`Failed to fetch transactions for address ${id}`);
    }
  };

    // Handle address deletion
  const handleDelete = async (id) => {
    try {
        await deleteAddress(id);
        setAddresses((prevAddresses) => prevAddresses.filter((addr) => addr !== id));
        if (selectedAddress === id) {
        setSelectedAddress(null);
        setSelectedAddressBalance(null);
        setTransactions([]);
        }
        alert("Address deleted successfully.");
    } catch (error) {
        console.error(`Failed to delete address with ID ${id}`);
        alert("Failed to delete address.");
    }
    };

  if (loading) return <p>Loading addresses...</p>;

  return (
    <div>
      <h2>Addresses</h2>
      <ul>
        {addresses.map((address) => (
        <div>
          <li
            key={address}
            style={{ cursor: "pointer", color: selectedAddress === address ? "blue" : "black" }}
            onClick={() => handleClick(address)}
          >
            {address}
          </li>
        <button onClick={() => handleDelete(address)} style={{ marginLeft: "10px" }}>
        Delete
        </button>
        </div>
        ))}
      </ul>

      {selectedAddress && (
        <div>
          <h2>Balance: {selectedAddressBalance} BTC</h2>
          <h2>Transactions for Address: {selectedAddress}</h2>
          {transactions.length > 0 ? (
            <ul>
              {transactions.map((tx, index) => (
                <li key={index}>
                  <strong>ID:</strong> {tx.hash}, <strong>Time:</strong> {tx.time}
                </li>
              ))}
            </ul>
          ) : (
            <p>No transactions found for this address.</p>
          )}
        </div>
      )}
    </div>
  );
};

export default AddressList;