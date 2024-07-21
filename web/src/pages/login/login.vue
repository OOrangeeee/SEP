<template>
  <div class="login-container">
    <div class="left-panel">
      <img src="../../assets/loginback-leftside.png" alt="Login Image" class="login-image1" />
    </div>
    <div class="right-panel">
      <form class="login-form">
<!--        <img src="../../assets/loginbackground2.png" alt="Login Image" class="login-image2" />-->
        <div class="login-box">
          <h2 style="color: white; margin-top: 0px;  font-size: 22px;">登录</h2>
          <form @submit.prevent="loginForm">
            <div class="form-group">
              <input type="text" v-model="username" placeholder="用户名" >
            </div>
            <div class="form-group">
              <input type="password" v-model="password" placeholder="密码" >
            </div>
            <div class="form-group remember-password flex">
              <input type="checkbox" v-model="rememberPassword">
              <label>记住密码</label>
            </div>
            <div class="form-group">
              <button type="submit">立即登录</button>
            </div>
          </form>
          <p>
            <span class="operation"  @click="gotoForgotPassword" style="font-size: 12px;">忘记密码</span>
            <span class="operation"> | </span>
            <span class="operation" @click="gotoRegister" style="font-size: 12px;">没有账号？立即注册</span>
          </p>
        </div>
      </form>
    </div>
  </div>
</template>
<script>
import {login} from "../../api/index"
import {mapActions,mapState} from "vuex";
import { SET_TOKEN } from "../../utils/token";
export default {
  data() {
    return {
      username: '',
      password: '',
      rememberPassword: false
    };
  },
  computed:{
      ...mapState(['userInfo'])
  },
  methods: {
    ...mapActions(['setUserInfo']),
    async loginForm() {
      

      if (!this.username) {
        return this.$message.warning('用户名不能为空')
      }
      if (!this.password) {

       return this.$message.warning('请输入密码');
      }
      const requestBody = {
        'user-name': this.username,
        'user-password': this.password
      };
    
      const res = await login(requestBody)
      const {token} = res
      SET_TOKEN(token)
      this.$message.success('登录成功');
      await this.setUserInfo();
      this.$router.replace('/home')
      
    },
    gotoRegister() {
      this.$router.push({ path: "/register" });
    },
    gotoForgotPassword() {
    }
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
}
.login-image1 {
 width: 85%;
  height: auto;
}
.right-panel {
  flex: 1;
  text-align: center;
  height: 700px;
  right: 0;
  width:700px;
  width: 50%; /* 让右侧面板占50%宽度 */
  position: relative; /* 设置相对定位 */
  z-index: 1; /* 确保背景在底部 */
  border: none; /* 设置边框为无 */
  display: flex;
  align-items: center;
  justify-content: center;
}
.login-image2 {
  //width: 90%; /* 设置背景图片宽度为100% */
  //height: 90%; /* 设置背景图片高度为100% */
  width: 70%;
  //position: absolute; /* 设置绝对定位 */
  top: -20px; /* 顶部对齐 */
  right: -60px; /* 右侧对齐 */
  z-index: -1; /* 确保背景在底部 */
  border: none; /* 设置边框为无 */
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
  background: url("../../assets/loginbackground2.png");
  height: 890px;
  width: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
  background-size: contain;
}
.form-group{
  margin-top: 20px;
}
.operation{
  cursor: pointer;
  color: #fff;
}
</style>
