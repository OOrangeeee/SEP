<template>
  <div class="home-container">
    <button @click="goToLogin">查看详细病灶</button>
  </div>
</template>

<script>
import {getRecordById} from "@/api";
import {mapActions} from "vuex"
export default {
  methods: {
    ...mapActions(['setRecordList']),
    goToLogin() {
      this.$router.push({path:"/detail1"});
    },
    async initDetail() {
      const promises = []
      this.ids.forEach(id => {
        promises.push(getRecordById(id))
      })
      // 获取病理图片
      Promise.all(promises).then((values) => {
          this.setRecordList(values)
      })
    }
  },
  created() {
    const id = this.$route.query.id;
    if(!id) {

      return this.$message.warning('请先选择报告')
    }
    this.ids = id.split(',')
    this.initDetail()
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
