import React, { useEffect, useState } from "react";
import { fetchAddresses, deleteAddress } from "../api/walletService";
import './AddressList.css'

const AddressList = () => {
  const [addresses, setAddresses] = useState([]);
  const [loading, setLoading] = useState(true);

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

  const handleDelete = async (id) => {
    try {
      await deleteAddress(id);
      setAddresses((prev) => prev.filter((addr) => addr !== id));
    } catch (error) {
      console.error(`Failed to delete address with ID ${id}`);
    }
  };

  if (loading) return <p>Loading addresses...</p>;

  return (
    <div >
      <h2 >Addresses</h2>
      <ul>
        {addresses.map((address) => (
          <li key={address}>
            {address}
            <button onClick={() => handleDelete(address)}>Delete</button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default AddressList;