<template>
  <div>
    <h2>Login</h2>
    <form @submit.prevent="login">
      <input v-model="userName" placeholder="Username" />
      <input v-model="userPassword" type="password" placeholder="Password" />
      <button type="submit">Login</button>
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
      message: '',
    };
  },
  methods: {
    login() {
      auth.getCsrfToken().then(() => {
        auth.login(this.userName, this.userPassword).then(response => {
          this.message = response.success_message;
          this.$router.push('/account');
        }).catch(error => {
          this.message = error.response.data.error_message;
        });
      });
    },
  },
};
</script>
