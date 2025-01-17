const API_BASE_URL = "http://localhost:8080";

export const ENDPOINTS = {
  GET_ADDRESSES: `${API_BASE_URL}/addresses`,
  GET_ADDRESS_BY_ID: (id) => `${API_BASE_URL}/addresses/${id}`,
  ADD_ADDRESS: (id) => `${API_BASE_URL}/addresses/${id}`,
  DELETE_ADDRESS: (id) => `${API_BASE_URL}/addresses/${id}`,
};
