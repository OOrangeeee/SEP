<template>
  <div>
    <h2>Register</h2>
    <form @submit.prevent="register">
      <input v-model="userName" placeholder="Username" />
      <input v-model="userPassword" type="password" placeholder="Password" />
      <input v-model="userEmail" placeholder="Email" />
      <input v-model="userNickname" placeholder="Nickname" />
      <input v-model="userAdminSecret" placeholder="Admin Secret" />
      <button type="submit">Register</button>
    </form>
    <p v-if="message">{{ message }}</p>
  </div>
</template>

<script>
import auth from '../services/auth';

export default {
  data() {
    return {
      userName: '',
      userPassword: '',
      userEmail: '',
      userNickname: '',
      userAdminSecret: '',
      message: '',
    };
  },
  methods: {
    register() {
      auth.getCsrfToken().then(() => {
        auth.register(this.userName, this.userPassword, this.userEmail, this.userNickname, this.userAdminSecret).then(response => {
          this.message = response.success_message;
          this.$router.push('/login');
        }).catch(error => {
          this.message = error.response.data.error_message;
        });
      });
    },
  },
};
</script>
