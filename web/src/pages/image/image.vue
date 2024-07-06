<template>
  <div class="home-container">
    <button @click="goToLogin">查看详细病灶</button>
  </div>
</template>

<script>
export default {
  methods: {
    goToLogin() {
      this.$router.push({path:"/detail1"});
    }
  }
}
</script>

<style scoped>
.home-container {
  background-image: url('../../assets/resultbackground.png');
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-size: 70% auto; /* 修改背景大小 */
  margin-left:22%;
  background-repeat: no-repeat; /* 防止重复填充 */
  z-index: 1;
  background-position-y: 30%;
}
button {
  padding: 10px 20px;
  font-size: 18px;
  cursor: pointer;
  margin-top:40%;
  margin-left:-30%;
  z-index: 0;
  width:160px;
  height:50px;
  border-radius: 30px;
  background-color: black;
  color: white;
}
button:hover {
  background-color: #5BA98D;
}
</style scoped>