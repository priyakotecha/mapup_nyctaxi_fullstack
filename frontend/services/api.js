import axios from 'axios';

const API_URL = 'http://localhost:3000'; // Your backend API

export const fetchAllData = async () => {
  const res = await axios.get(`${API_URL}/admin/getdata`);
  return res.data;
};

export const fetchManagerData = async () => {
  const res = await axios.get(`${API_URL}/manager/getdata`);
  return res.data;
};

export const fetchUserData = async () => {
  const res = await axios.get(`${API_URL}/user/getdata`);
  return res.data;
};
