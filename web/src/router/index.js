import Vue from 'vue';
import Router from 'vue-router';
import Login from '../components/Login.vue';
import Register from '../components/Register.vue';
import AccountInfo from '../components/AccountInfo.vue';
import UserRecords from '../components/UserRecords.vue';

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: '/login',
      name: 'Login',
      component: Login,
    },
    {
      path: '/register',
      name: 'Register',
      component: Register,
    },
    {
      path: '/account',
      name: 'AccountInfo',
      component: AccountInfo,
    },
    {
      path: '/records',
      name: 'UserRecords',
      component: UserRecords,
    },
    {
      path: '*',
      redirect: '/login',
    },
  ],
});
