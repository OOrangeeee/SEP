<template>
<div>
  <div class="home-container">
    <!-- Buttons -->
    <button @click="showImage('image1')" class="button1">疾病成因</button>
    <button @click="showImage('image2')" class="button2">如何预防</button>
    <button @click="showImage('image3')" class="button3">治疗方案</button>
    <button @click="showImage('image4')" class="button4">医学突破</button>
    <button @click="showImage('image5')" class="button5">典型性状</button>
    <button @click="showImage('image6')" class="button6">疾病历史</button>

    <!-- Images -->
    <img v-if="currentImage === 'image1'" @click="hideImage" class="centered-image" src="../../assets/jibingchengyin.png" alt="Image 1">
    <img v-if="currentImage === 'image2'" @click="hideImage" class="centered-image" src="../../assets/ruheyufang.png" alt="Image 2">
    <img v-if="currentImage === 'image3'" @click="hideImage" class="centered-image" src="../../assets/zhiliaofangan.png" alt="Image 3">
    <img v-if="currentImage === 'image4'" @click="hideImage" class="centered-image" src="../../assets/yixuetupo.png" alt="Image 4">
    <img v-if="currentImage === 'image5'" @click="hideImage" class="centered-image" src="../../assets/dianxingxingzhuang.png" alt="Image 5">
    <img v-if="currentImage === 'image6'" @click="hideImage" class="centered-image" src="../../assets/jibinglishi.png" alt="Image 6">
  </div>
</div>
</template>

<script>
export default {
  data() {
    return {
      currentImage: null
    };
  },
  methods: {
    showImage(image) {
      this.currentImage = image;
    },
    hideImage() {
      this.currentImage = null;
    }
  }
}
</script>

<style scoped>

.home-container {
  background-image: url('../../assets/popularzationicon.png');
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  background-size: 70% auto; /* 修改背景大小 */
  background-position: center top 20px; /* 将背景图片距离页面顶端10px */
  background-repeat: no-repeat; /* 防止重复填充 */
  z-index: 1;
}

button {
  padding: 10px 20px;
  cursor: pointer;
  margin-top: 30px;
  z-index: 0;
  width: 180px;
  height: 50px;
  border-radius: 10px;
  background-color: transparent; /* 设置背景颜色透明 */
  border: none; /* 隐藏边框 */
  color:black;
}
.button1 {
  position: absolute;
  top: 130px;
  left: 240px;
  font-size: 22px;
}

.button2 {
  position: absolute;
  top: 70px;
  right: 570px;
  font-size: 22px;
}

.button3 {
  position: absolute;
  top: 110px;
  right: 260px;
  font-size: 22px;
}

.button4 {
  position: absolute;
  bottom: 65px;
  left: 230px;
  font-size: 22px;
}

.button5 {
  position: absolute;
  bottom: 130px;
  left: 520px;
  font-size: 22px;
}

.button6 {
  position: absolute;
  bottom: 80px;
  right: 260px;
  font-size: 22px;
}

.centered-image {
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  height: 80%;
  width: 65%;
  z-index: 2;
}
</style>