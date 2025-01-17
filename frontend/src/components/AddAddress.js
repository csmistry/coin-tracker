import React, { useState } from "react";
import { addAddress } from "../api/walletService";
import './AddAddress.css'

const AddAddress = () => {
  const [newAddress, setNewAddress] = useState("");

  const handleAdd = async () => {
    if (!newAddress) return alert("Address cannot be empty!");
    try {
      await addAddress(newAddress);
      alert("Address added successfully!");
      setNewAddress("");
    } catch (error) {
      console.error("Failed to add address.");
    }
  };

  return (
    <div>
      <input
        type="text"
        value={newAddress}
        onChange={(e) => setNewAddress(e.target.value)}
        placeholder="Enter address"
        className="text-input"
      />
      <button onClick={handleAdd}>Add</button>
    </div>
  );
};

export default AddAddress;