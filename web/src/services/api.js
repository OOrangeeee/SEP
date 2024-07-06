import axios from 'axios';

const apiClient = axios.create({
  baseURL: 'http://localhost:714',
  withCredentials: true, // 允许携带cookie
  headers: {
    'Content-Type': 'application/json',
  },
});

apiClient.interceptors.request.use((config) => {
  const csrfToken = localStorage.getItem('csrfToken');
  if (csrfToken) {
    config.headers['X-Csrf-Token'] = csrfToken;
  }
  return config;
});

export default apiClient;
