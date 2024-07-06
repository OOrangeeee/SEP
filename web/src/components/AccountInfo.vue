<template>
    <div>
      <h2>Account Info</h2>
      <div v-if="userInfo">
        <p><strong>Username:</strong> {{ userInfo.UserName }}</p>
        <p><strong>Email:</strong> {{ userInfo.UserEmail }}</p>
        <p><strong>Nickname:</strong> {{ userInfo.UserNickName }}</p>
        <p><strong>Is Admin:</strong> {{ userInfo.UserIsAdmin }}</p>
      </div>
      <p v-else>{{ message }}</p>
    </div>
  </template>
  
  <script>
  import auth from '../services/auth';
  
  export default {
    data() {
      return {
        userInfo: null,
        message: '',
      };
    },
    created() {
      auth.getAccountInfo().then(response => {
        this.userInfo = response.userInfo;
      }).catch(error => {
        this.message = error.response.data.error_message;
      });
    },
  };
  </script>
  