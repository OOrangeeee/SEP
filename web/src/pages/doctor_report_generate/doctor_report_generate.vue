<template>
  <div class="container">
    <div class="left-component">
      <h2>生成新报告</h2>
      <div class="left-buttons">
        <el-upload
          class="upload"
          action=""
          :before-upload="(e)=>beforeUpload(e,1)"
          :http-request="(e)=>uploadFile(e,1)"
          :show-file-list="false"
          :on-change="handleFileChange"
          accept="video/*"
        >

          <video :src="videoUrl" controls  v-if="videoUrl"></video>
          <div class="uploadBtn" v-else  :class="{ activeLeftButton: activeLeftButton === 1 }" style="color: black;">导入视频</div>
        </el-upload>
        <el-upload
          class="upload"
          action=""
          :before-upload="(e)=>beforeUpload(e,2)"
          :http-request="(e)=>uploadFile(e,2)"
          :show-file-list="false"
          :on-change="handleAvatarSuccess"
          accept="image/*"
        >
          <img :src="imageUrl" alt="" v-if="imageUrl">
          <div class="uploadBtn" v-else  :class="{ activeLeftButton: activeLeftButton === 2 }" style="color: black;">导入图片</div>
        </el-upload>

      </div>
      <div class="left-buttons2" v-if="activeLeftButton === 1 || activeLeftButton === 2">
        <button @click="toggleLeftButton(3)" :class="{ activeLeftButton: activeLeftButtons.includes(3) }" style="background-color: white;color:black;width: 500px;height: 38px;margin-left: 0px; margin-top: 32px;border-radius: 10px;margin-bottom: 5px;">追踪</button>
        <button @click="toggleLeftButton(4)" :class="{ activeLeftButton: activeLeftButtons.includes(4) }" style="background-color: white;color:black;width: 500px;height: 38px;margin-left: 0px; border-radius: 10px;margin-bottom: 5px;">分割</button>
        <button @click="toggleLeftButton(5)" :class="{ activeLeftButton: activeLeftButtons.includes(5) }"  style="background-color: white;color:black;width: 500px;height: 38px;margin-left: 0px; border-radius: 10px;margin-bottom: 5px;">诊断</button>
        <button @click="showConfirmationPopup" style="background-color: white;color:black;width: 250px;height: 38px;margin-left: 30px; border-radius: 10px;margin-bottom: 5px;">确认（可多选）</button>
      </div>
    </div>
    <div class="right-component">
      <img src="../../assets/avatar.png" alt="Image"> <!-- 替换为正确的图片路径 -->
      <h2>x医生</h2>
      <div class="scrollable-buttons">
        <button :class="{ activeButton: activeRightButton === 1 }" @click="handleRightButtonClick(1)">钟文唐 10<br>上一次就诊：2024.1.1 9:00-9:30</button>
        <button :class="{ activeButton: activeRightButton === 2 }" @click="handleRightButtonClick(2)">晋晨曦 20<br>上一次就诊：2024.1.1 9:30-10:00</button>
        <button :class="{ activeButton: activeRightButton === 3 }" @click="handleRightButtonClick(3)">刘朗清 30<br>上一次就诊：2024.1.1 10:00-10:30</button>
        <button :class="{ activeButton: activeRightButton === 4 }" @click="handleRightButtonClick(4)">靳文惠 40<br>上一次就诊：2024.1.1 10:30-11:00</button>
      </div>
    </div>

    <!-- 弹窗组件 -->
    <div class="popup" v-if="showPopup">
      <div class="popup-content">
        <h3>温馨提示</h3>
        <p style="font-size: 10px;">新报告已生成！</p>
        <button @click="hideConfirmationPopup">确定</button>
      </div>
    </div>
  </div>
</template>

<script>
import {segmentation,detection,track,getRecords} from "@/api";
import {mapActions} from "vuex";
export default {
  data() {
    return {
      activeRightButton: null,
      activeLeftButton: null,
      activeLeftButtons: [],
      showPopup: false, // 控制弹窗显示与隐藏的状态
      videoUrl:null,
      imageUrl:null,
      recordList:[],
      imageFormData:null,
      videoFormData:null,
      imageSize:true,
      videoSize:true,
      resultsId:''
    };
  },
  created() {
    this.initRecordList();
  },
  methods: {
    ...mapActions(['setRecordList']),
    // 获取记录
    async initRecordList() {
      const {records} = await getRecords();
      this.recordList = records;
    },
    handleLeftButtonClick(buttonNumber) {
      // 左边按钮点击事件处理逻辑
      console.log('Left button ' + buttonNumber + ' clicked');
      // 设置当前选中的左侧按钮
      this.activeLeftButton = buttonNumber;
      // Toggle active state for tracking, segmentation, diagnosis, and confirmation buttons
      if (buttonNumber === 1 || buttonNumber === 2) {
        this.toggleLeftButton(3);
        this.toggleLeftButton(4);
        this.toggleLeftButton(5);
        this.toggleLeftButton(6);
      }
    },
    toggleLeftButton(buttonNumber) {
      // 切换左边按钮的选中状态
      const index = this.activeLeftButtons.indexOf(buttonNumber);
      if (index === -1) {
        // 如果按钮未选中，则添加到选中列表中
        this.activeLeftButtons.push(buttonNumber);
      } else {
        // 如果按钮已选中，则从选中列表中移除
        this.activeLeftButtons.splice(index, 1);
      }
    },
    handleRightButtonClick(buttonNumber) {
      // 右边按钮点击事件处理逻辑
      console.log('Right button ' + buttonNumber + ' clicked');
      this.activeRightButton = buttonNumber;
      this.$router.push({
        path:'/image',
        query:{
          id:buttonNumber,
        }
      })
    },

    // 提交文件获取确诊结果
    showConfirmationPopup() {
      let promises = [];
      if (this.activeLeftButtons.includes(3) || this.activeLeftButtons.includes(4) || this.activeLeftButtons.includes(5)) {
        //功能根据视频和图片来划分
        // 视频功能
        if (this.activeLeftButtons.includes(3) && this.videoUrl) {
          // 只有追踪
          promises.push(track(this.videoFormData));
        }
        // 图片功能
        if (this.imageUrl) {
          // 分割
          if (this.activeLeftButtons.includes(4)) {
            promises.push(segmentation(this.imageFormData));
          }
          // 诊断
          if (this.activeLeftButtons.includes(5)) {
            promises.push(detection(this.imageFormData));
          }
        }
        if (promises.length==0) {
          return this.$message.warning('请选择正确的功能');
        }

        // 等待所有的结果全部出来以后，才提示
        Promise.all(promises).then((values)=>{
          // 显示弹窗
          this.showPopup = true;
          // values 装的是每一次结果
          // id 根据values去取，这里只是静态数据
          this.resultsId = '1,2,3';

        })
      }else{
        this.$message.warning('请选择功能');
      }


    },
    hideConfirmationPopup() {

      // 隐藏弹窗
      this.showPopup = false;
      this.$router.push({
        path:'/image',
        query:{
          id:this.resultsId
        }

      })
    },
    beforeUpload (file,index) {

      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isLt2M) {
          this.$message.warning('上传文件不能超过2MB')
      }
      return isLt2M
    },
    uploadFile (option,index) {
      const formData = new FormData();
      if (index===1) {
        formData.append('file', option.file);
        formData.append('patient-name','张三')
        this.videoFormData = formData

      }else if (index===2) {
        formData.append('file', option.file);
        formData.append('patient-name','张三')
        this.imageFormData = formData
      }
      this.handleLeftButtonClick(index)
    },
    handleFileChange(file, fileList) {
      // 假设我们只处理第一个文件
      if (fileList.length > 0) {
        const fileObj = fileList[0].raw; // 获取原始的 File 对象
        this.createVideoUrl(fileObj);
      }
    },
    createVideoUrl(file) {
      if (file) {
        const isLt2M = file.size / 1024 / 1024 < 2;
        if (!isLt2M) {
          this.videoSize = false
          return
        }else {
          this.videoSize = true
        }
        this.videoUrl = URL.createObjectURL(file); // 创建一个指向该文件的 URL
      }
    },
    handleAvatarSuccess(file, fileList) {
      const isLt2M = file.size / 1024 / 1024 < 2;
      if (!isLt2M) {
        this.imageSize = false
        return
      }else {
        this.imageSize = true
      }
      // 创建一个FileReader实例
      const reader = new FileReader();

      // 读取文件完成后执行的回调函数
      reader.onload = (e) => {
        // 读取完成后的结果就是一个base64字符串
        this.imageUrl = e.target.result;
      };

      // 使用readAsDataURL方法读取file对象
      reader.readAsDataURL(file.raw ? file.raw : file); // 假设file.raw是文件对象，否则直接使用file
    },

  },
  beforeDestroy() {
    if (this.videoUrl) {
      URL.revokeObjectURL(this.videoUrl); // 释放 URL 对象
    }
  },
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
  padding: 0 120px;
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
  max-height:calc(100vh - 200px);
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
  color:black;
  padding: 0 40px;
}

.left-component .uploadBtn,.left-component video,.left-component img {

  border-radius: 5px;
  background-color: #E7EDEB;
  cursor: pointer;
  height: 150px;
  width: 170px;
  box-sizing: border-box;
  display: flex;
  align-items:center;
  justify-content: center;
}
.scrollable-buttons {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  color:black;
}

.right-component button {
  width: 420px;
  height: 90px;
  margin: 10px 0;
  padding: 10px;
  border: none;
  border-radius: 20px;
  background-color: #E7EDEB;
  color: black;
  cursor: pointer;
}
.right-component img {
  max-width: 20%;
  height: auto;
  margin-top: 40px;
}
.left-component button:hover {
  background-color: #5BA98D;
}
.activeButton {
  background-color: #5BA98D !important; /* 使用 !important 来确保优先级 */
}
.activeLeftButton {
  background-color: #5BA98D !important; /* 使用 !important 来确保优先级 */
}
/* 省略已有样式，添加弹窗样式 */
.popup {
  position: fixed;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
  background: linear-gradient(to bottom left, rgba(87, 166, 140, 0.5), rgba(240, 255, 251, 0.95)); /* 从右上到左下的颜色渐变，设置透明度为0.5 */
  width: 300px;
  padding: 0px;
  border-radius: 20px;
  z-index: 9999;
  box-shadow: 5px 5px 10px rgba(0, 0, 0, 0.5); /* 添加阴影效果 */
  border: none; /* 去除原有的边框 */
}

.popup-content {
  background: linear-gradient(to bottom left, rgba(87, 166, 140, 0.5), rgba(240, 255, 251, 0.95));
  padding: 20px;
  border-radius: 20px;
  text-align: center;
}

.popup button {
  padding: 10px 20px;
  background-color: black;
  color: white;
  border: none;
  border-radius: 25px;
  cursor: pointer;
  margin-top: 20px;
  width: 120px;
}

.popup button:hover {
  background-color: #44956B;
}
.upload{
  //border:1px solid #000
}
</style>
