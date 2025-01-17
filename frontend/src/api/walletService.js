// Fetch addresses in wallet
import axios from "axios";
import { ENDPOINTS } from "./endpoints";

export const fetchAddresses = async () => {
  try {
    const response = await axios.get(ENDPOINTS.GET_ADDRESSES);
    return response.data;
  } catch (error) {
    console.error("Error fetching addresses:", error);
    throw error;
  }
};

// Fetch address by ID
export const fetchAddressById = async (id) => {
  try {
    const response = await axios.get(ENDPOINTS.GET_ADDRESS_BY_ID(id));
    return response.data;
  } catch (error) {
    console.error(`Error fetching address with ID ${id}:`, error);
    throw error;
  }
};

// Add a new address
export const addAddress = async (id) => {
  try {
    const response = await axios.post(ENDPOINTS.ADD_ADDRESS(id));
    return response.data;
  } catch (error) {
    console.error(`Error adding address ${id} to wallet:`, error);
    throw error;
  }
};

// Delete an address
export const deleteAddress = async (id) => {
  try {
    const response = await axios.delete(ENDPOINTS.DELETE_ADDRESS(id));
    return response.data;
  } catch (error) {
    console.error(`Error deleting address with ID ${id}:`, error);
    throw error;
  }
};
