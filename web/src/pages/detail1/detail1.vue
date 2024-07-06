<template>
  <div class="home-container">
    <button @click="goToImage" class="button1">返回</button>
    <button @click="openModal" class="button2">下载报告</button>
    <button @click="goToDetail2" class="button3"></button>

    <!-- 模态框 -->
    <div class="modal" v-if="showModal">
      <div class="modal-content">
        <span class="close" @click="closeModal">&times;</span>
        <h2>患者信息</h2>
        <form @submit.prevent="submitForm">
          <input type="text" id="name" v-model="formData.name" style="height: 12px;" placeholder="姓名" required>
          <input type="date" id="date" v-model="formData.date" style="height: 25px; width: 89%; 
border-radius: 30px; margin-bottom: 10px; margin-top: 2px;" placeholder="就诊日期" required>
          <input type="text" id="symptom" v-model="formData.symptom" style="height: 12px;" placeholder="症状" required>
          <button type="submit" style="background-color: black; color: white; width: 22%;">下载</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      showModal: false,
      formData: {
        name: '',
        date: '',
        symptom: ''
      }
    };
  },
  methods: {
    goToImage() {
      this.$router.push({path:"/image"});
    },
    goToDetail2() {
      this.$router.push({path:"/detail2"});
    },
    openModal() {
      this.showModal = true;
    },
    closeModal() {
      this.showModal = false;
    },
    submitForm() {
      // 提交表单的逻辑
      // 这里可以将表单数据发送到服务器或执行其他操作
      // 然后关闭模态框
      this.closeModal();
    }
  }
};
</script>

<style scoped>
.home-container {
  background-image: url('../../assets/detailbackground1.png');
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-size: 60% auto; /* 修改背景大小 */
  margin-left:28%;
  background-repeat: no-repeat; /* 防止重复填充 */
  z-index: 1;
  background-position-y: 60%;
  position: relative; /* 父容器设置为相对定位 */
}
.button1{
  padding: 10px 20px;
  font-size: 13px;
  cursor: pointer;
  top:10%;
  left:-23%;
  z-index: 0;
  width:80px;
  height:40px;
  border-radius: 30px;
  background-color:#5BA98D;
  color: black;
  position: absolute; /* 设置按钮为绝对定位 */
}
.button2{
  padding: 10px 20px;
  font-size: 13px;
  cursor: pointer;
  top:10%;
  right:15%;
  z-index: 0;
  width:100px;
  height:40px;
  border-radius: 30px;
  background-color: black;
  color: white;
  position: absolute; /* 设置按钮为绝对定位 */
}
.button3 {
  background-image: url('../../assets/arrow.png');
  width: 40px;
  height: 40px;
  background-color: rgba(0, 0, 0, 0); /* 使用透明色 */
  background-size: cover;
  left:60%;
  position: absolute; /* 设置按钮为绝对定位 */
}
button:hover {
  background-color: #5BA98D;
}
.modal {
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgb(0, 0, 0);
  background-color: rgba(0, 0, 0, 0.4);
}

.modal-content {
  background: linear-gradient(to bottom left, rgba(87, 166, 140, 0.5), rgba(240, 255, 251, 0.95)); /* 从右上到左下的颜色渐变，设置透明度为0.5 */
  margin: 15% auto;
  padding: 20px;
  border: none; /* 去除原有的边框 */
  border-radius: 15px; /* 设置边框为圆角 */
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.5); /* 添加阴影效果 */
  width: 25%;
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}
</style>