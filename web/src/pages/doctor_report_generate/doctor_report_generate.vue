<template>
  <div class="container">
    <div class="left-component">
      <h2>生成新报告</h2>
      <div class="left-buttons">
        <button @click="handleLeftButtonClick(1)">导入视频</button>
        <button @click="handleLeftButtonClick(2)">导入图片</button>
      </div>
      <div class="left-buttons2">
        <button @click="handleLeftButtonClick(3)" style="background-color: white;color:black;width: 500px;height: 38px;margin-left: 0px; margin-top: 32px;border-radius: 10px;margin-bottom: 5px;">追踪</button>
        <button @click="handleLeftButtonClick(4)" style="background-color: white;color:black;width: 500px;height: 38px;margin-left: 0px; border-radius: 10px;margin-bottom: 5px;">分割</button>
        <button @click="handleLeftButtonClick(5)" style="background-color: white;color:black;width: 500px;height: 38px;margin-left: 0px; border-radius: 10px;margin-bottom: 5px;">诊断</button>
        <button @click="handleLeftButtonClick(6)" style="background-color: white;color:black;width: 250px;height: 38px;margin-left: 30px; border-radius: 10px;margin-bottom: 5px;">确认（可多选）</button>
      </div>
    </div>
    <div class="right-component">
      <img src="../../assets/avatar.png" alt="Image"> <!-- 替换为正确的图片路径 -->
      <h2>x医生</h2>
      <div class="scrollable-buttons">
        <button :class="{activeButton: activeRightButton === 1}" @click="handleRightButtonClick(1)">Button 1</button>
        <button :class="{activeButton: activeRightButton === 2}" @click="handleRightButtonClick(2)">Button 2</button>
        <button :class="{activeButton: activeRightButton === 3}" @click="handleRightButtonClick(3)">Button 3</button>
        <button :class="{activeButton: activeRightButton === 4}" @click="handleRightButtonClick(4)">Button 4</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      activeRightButton: null
    };
  },
  methods: {
    handleLeftButtonClick(buttonNumber) {
      // 左边按钮点击事件处理逻辑
      console.log('Left button ' + buttonNumber + ' clicked');
    },
    handleRightButtonClick(buttonNumber) {
      // 右边按钮点击事件处理逻辑
      console.log('Right button ' + buttonNumber + ' clicked');
      this.activeRightButton = buttonNumber;
    }
  }
};
</script>
<style>
.container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 100vh;
  background-image: url('../../assets/reporterbackground.png'); /* 替换为正确的路径 */
  background-size: 20%;
  background-repeat: no-repeat;
  background-position-x: 300px;
  background-position-y: 150px;
}
.left-component {
  text-align: center;
  width: 500px;
  height: 260px;
  background-color: white;
  margin-left: 60px;
  margin-top: -180px;
  border-radius: 30px;
}
.right-component {
  text-align: center;
  width: 500px;
  height: 400px;
  background-color: white;
  margin-right: 60px;
  margin-top: -40px;
  border-radius: 30px;
  overflow-y: auto; /* 允许垂直滚动 */
}
.left-component h2 {
  color: black;
}
.left-buttons {
  display: flex;
  justify-content: space-between;
}
.left-component button {
  margin: 10px;
  margin-left: 40px;
  margin-right:40px;
  padding: 10px;
  border-radius: 5px;
  background-color: #E7EDEB;
  color: white;
  cursor: pointer;
  height:150px;
  width:200px;
}
.scrollable-buttons {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
}

.right-component button {
  width: 420px;
  height: 90px;
  margin: 10px 0;
  padding: 10px;
  border: none;
  border-radius: 20px;
  background-color: #E7EDEB;
  color: white;
  cursor: pointer;
}
.right-component img {
  max-width: 20%;
  height: auto;
  margin-top:40px;
}
.left-component button:hover {
  background-color:#5BA98D;
}
.activeButton {
  background-color: #5BA98D !important; /* 使用 !important 来确保优先级 */
}
</style>