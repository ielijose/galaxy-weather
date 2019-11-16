import axios from 'axios';
const host = window.location.host || 'localhost';
const url = `http://${host}:8080/api/galaxy-weather/`;

const api = axios.create({
  baseURL: url,
});

export default api;
