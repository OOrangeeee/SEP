<template>
  <div id="app">
    <header class="header">
      <router-link to="/home" class="header-button" :class="{ 'active': $route.path === '/home' }">首页</router-link>
      <router-link to="/doctor_report_generate" class="header-button" :class="{ 'active': $route.path === '/doctor_report_generate' }">息肉诊断</router-link>
      <router-link to="/image" class="header-button" :class="{ 'active': $route.path === '/image' }">诊断报告</router-link>
      <router-link to="/tabbar" class="header-button" :class="{ 'active': $route.path === '/tabbar' }">关于息肉</router-link>
    </header>
    <router-view></router-view>
  </div>
</template>

<script>
import {getCsrf} from "./api/index"
export default {
  name: 'App',
  // 在这里可以添加一些全局逻辑或生命周期钩子函数
  async created() {
   const result = await getCsrf()
   console.log(result)

  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background-color: #F5F5F5;
}
.header {
  display: flex;
  align-items: center;
  justify-content: center;
  position: fixed;
}
body{
	margin:0;
	padding:0;
	border:0
}
.header-button {
  margin-left: 230px;
  margin-right: -130px;
  width: 100px;
  height: 30px;
  border-radius: 20px;
  padding-top: 10px;
  text-align: center;
  background-color: transparent; /* 初始时背景透明 */
  color: black; /* 初始时字体为黑色 */
  transition: background-color 0.3s, color 0.3s; /* 添加过渡效果 */
}

.header-button:hover {
  background-color: #5BA98D; /* 鼠标悬停时背景为深绿色 */
  color: white; /* 鼠标悬停时字体为白色 */
}

.active {
  /* 选中状态的样式设置，比如深绿色 */
  background-color: #5BA98D; /* 选中时背景为深绿色 */
  color: white; /* 选中时字体为白色 */
}
</style>