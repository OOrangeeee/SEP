import apiClient from './api';

export default {
  getCsrfToken() {
    return apiClient.get('/csrf-token').then(response => {
      const csrfToken = response.headers['x-csrf-token'];
      localStorage.setItem('csrfToken', csrfToken);
    });
  },

  login(userName, userPassword) {
    return apiClient.post('/users/login', {
      'user-name': userName,
      'user-password': userPassword,
    }).then(response => {
      localStorage.setItem('authToken', response.data.token);
      return response.data;
    });
  },

  register(userName, userPassword, userEmail, userNickname, userAdminSecret) {
    return apiClient.post('/users/account', {
      'user-name': userName,
      'user-password': userPassword,
      'user-email': userEmail,
      'user-nickname': userNickname,
      'user-admin-secret': userAdminSecret,
    }).then(response => response.data);
  },

  getAccountInfo() {
    const authToken = localStorage.getItem('authToken');
    return apiClient.get('/users/account', {
      headers: {
        Authorization: `Bearer ${authToken}`,
      },
    }).then(response => response.data);
  },

  getUserRecords() {
    const authToken = localStorage.getItem('authToken');
    return apiClient.get('/users/records-all', {
      headers: {
        Authorization: `Bearer ${authToken}`,
      },
    }).then(response => response.data);
  },
};
