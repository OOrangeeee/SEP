<template>
  <div class="home-container">
    <button @click="goToLogin">开始使用</button>
  </div>
</template>

<script>
export default {
  methods: {
    goToLogin() {
      this.$router.push({path:"/login"})
    }
  }
}
</script>

<style scoped>
.home-container {
  background-image: url('../../assets/homebackground.png');
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-size: 60% auto; /* 修改背景大小 */
  background-position: center;
  background-repeat: no-repeat; /* 防止重复填充 */
  z-index: 1;
}
button {
  padding: 10px 20px;
  font-size: 18px;
  cursor: pointer;
  margin-top:30%;
  z-index: 0;
  width:160px;
  height:50px;
  border-radius: 30px;
  background-color: black;
  color: white;
}
button:hover {
  background-color: rgb(144, 222, 202);
}
</style scoped>