<template>
  <div class="login-container">
    <div class="left-panel">
      <img src="../../assets/loginback-leftside.png" alt="Login Image" class="login-image1" />
    </div>
    <div class="right-panel">
      <form class="login-form">
                <img src="../../assets/loginbackground2.png" alt="Login Image" class="login-image2" />
        <div class="login-box">
          <h2 style="color: white; margin-top: 0px;  font-size: 22px;">注册</h2>
          <form @submit.prevent="login">
            <div class="form-group">
              <input type="text" v-model="username" placeholder="请输入用户名" >
            </div>
            <div class="form-group">
              <input type="password" v-model="password" placeholder="请输入密码" >
            </div>
            <div class="form-group">
              <input type="text" v-model="email" placeholder="请输入邮箱" >
            </div>
            <div class="form-group">
              <input type="text" v-model="nickName" placeholder="请输入昵称" >
            </div>
            <div class="form-group">
              <button type="submit">立即注册</button>
            </div>
          </form>

        </div>
      </form>
    </div>
  </div>
</template>
<script>
import {register,activationUser,getUserInfo} from "../../api/index"
import {mapActions} from "vuex"
export default {
  data() {
    return {
      username: '',
      password: '',
      email: '',
      nickName:''
    };
  },
  computed:{

  },
  methods: {
    ...mapActions(['setUserInfo']),
    async login() {
      // const phoneRegex = /^1[3-9]\d{9}$/;
      // if (!phoneRegex.test(this.username)) {
      //  return this.$message.warning('请输入正确的手机号');
      // }
      if (!this.username) {
        return this.$message.warning('用户名不能为空')
      }
      if (!this.password) {
       return  this.$message.warning('请输入密码');
      }
      // 邮箱正则表示式
      const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;

      if (!emailRegex.test(this.email)) {
        return this.$message.warning('请输入正确的邮箱');
      }

      if (!this.nickName) {
        return this.$message.warning('请输入昵称');
      }
      const requestBody = {
        'user-name': this.username,
        'user-password': this.password,
        'user-email':this.email,
        'user-nickname':this.nickName
      };
      await register(requestBody)
      this.$message.success('注册成功');
      this.$router.push('/login');
    },

  }
};
</script>
<style scoped>

.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-right: 0px;
  margin-left: 20px;
  border: none; /* 设置边框为无 */
}
.left-panel {

  text-align: center;
  width: 50%; /* 让左侧面板占50%宽度 */
  height: 100vh;
}
.login-image1 {
  width: 100%;
  height: calc(100% - 10px);
}
.right-panel {
  flex: 1;
  text-align: center;

  right: 0;
  height: 100vh;
  position: relative; /* 设置相对定位 */
  z-index: 1; /* 确保背景在底部 */
  border: none; /* 设置边框为无 */
  display: flex;
  align-items: center;
  justify-content: center;
}
.login-image2 {
  position: absolute;
  width: 100%;
  height: 100%;
  left: 0;
  top:0;
}
.login-box {
  margin: 0 auto; /* 让表单居中 */
  border: 0; /* 隐藏边框 */
  padding: 20px; /* 添加内边距 */
  background: linear-gradient(to bottom right, rgba(255, 255, 255, 0), rgba(255, 255, 255, 0.1));
  //position: fixed;
  top: 40%; /* 设置距离顶部的距离为40% */
  left: 71%; /* 设置距离左侧的距离为50%，使其水平居中 */
  border-radius: 20px;
  //transform: translate(-50%, -50%); /* 使用transform属性使元素水平垂直居中 */
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.5); /* 将阴影放置在左下角 */
  backdrop-filter: blur(8px); /* 添加背景模糊 */
  margin-top:50px;
  width:300px;

}
.login-form input[type='text'], .login-form input[type="password"]  {
  width: 100%;
  padding: 10px;

  border-radius: 20px;
  margin-bottom: 5px;
  height:35px;
  margin-top:-10px;
  box-sizing: border-box;
}
.login-form button {
  width: 90%;
  padding: 5px;
  background-color: black; /* 蓝色背景 */
  color: #fff; /* 白色文字 */
  border: none;
  margin-top:-10px;
  border-radius: 30px;
  cursor: pointer;
  height: 36px;
}
.login-form button:hover {
  background-color: black; /* 深蓝色背景 */
  cursor: pointer;
}
.remember-password {
  display: flex;
  margin-top:15px;
  margin-left:140px;
  align-items: center;
  justify-content: flex-end;
}
.remember-password label {

  font-size: 12px;
}
.login-form{
  height: 100%;
  width: 100%;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;

}
.form-group{
  margin-top: 20px;
}
.operation{
  cursor: pointer;
  color: #fff;
}
</style>
