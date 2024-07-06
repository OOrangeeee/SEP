<template>
  <div class="flex-col page">
    <div class="flex-col justify-start items-end relative group">
      <img
        class="button image_2"
        src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031442201627424.png"
      />
      <img
        class="image pos"
        src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031809393924174.png"
      />
    </div>
    <div class="flex-col mt-57">
      <div class="flex-row justify-between">
        <img
          class="button_2"
          src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031801925313762.png"
        />
        <img
          class="image_3"
          src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031802393555831.png"
        />
      </div>
      <div class="mt-6 flex-col justify-start items-center relative group_2">
        <img
          class="image_5"
          src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031790202399773.png"
        />
        <img
          class="image_6 pos_4"
          src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031803874466184.png"
        />
        <img
          class="image_7 pos_5"
          src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031804189825711.png"
        />
        <img
          class="image_8 pos_6"
          src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031805242702447.png"
        />
        <img
          class="image_4 pos_2"
          src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031808372842135.png"
        />
        <div class="flex-col section pos_3">
          <span class="self-center text">患者信息</span>
          <div class="mt-36 flex-col self-stretch">
            <div class="flex-col justify-start items-start self-stretch text-wrapper">
              <span class="font text_2 text_3">姓名</span>
            </div>
            <div class="mt-38 flex-col justify-start items-start self-stretch text-wrapper">
              <span class="font text_2">就诊日期</span>
            </div>
            <div class="mt-38 flex-col justify-start items-start self-stretch text-wrapper">
              <span class="font text_2 text_4">症状</span>
            </div>
            <div class="mt-38 flex-col justify-start items-center button_3 text-wrapper_2">
              <span class="font text_5">下载</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    components: {},
    props: {},
    data() {
      return {};
    },

    methods: {},
  };
</script>

<style scoped lang="css">
  .mt-57 {
    margin-top: 57px;
  }
  .page {
    padding: 28px 99px 74px 113px;
    background-color: #e7edeb;
    width: 100%;
    overflow-y: auto;
    overflow-x: hidden;
    height: 100%;
  }
  .group {
    padding-bottom: 10px;
  }
  .button {
    margin-right: 549px;
  }
  .image_2 {
    border-radius: 40px;
    width: 103px;
    height: 40px;
  }
  .image {
    width: 1100px;
    height: 50px;
  }
  .pos {
    position: absolute;
    left: 201px;
    bottom: 0;
  }
  .button_2 {
    width: 124px;
    height: 57px;
  }
  .image_3 {
    width: 135px;
    height: 57px;
  }
  .group_2 {
    padding-top: 5px;
  }
  .image_5 {
    filter: blur(20px);
    border-radius: 50%;
    width: 984px;
    height: 840px;
  }
  .image_6 {
    width: 467px;
    height: 492px;
  }
  .pos_4 {
    position: absolute;
    left: 519px;
    top: 50%;
    transform: translateY(-50%);
  }
  .image_7 {
    width: 90px;
    height: 79px;
  }
  .pos_5 {
    position: absolute;
    right: 249px;
    bottom: 362px;
  }
  .image_8 {
    width: 541px;
    height: 209px;
  }
  .pos_6 {
    position: absolute;
    left: 49px;
    bottom: 116px;
  }
  .image_4 {
    width: 533px;
    height: 279px;
  }
  .pos_2 {
    position: absolute;
    right: 66px;
    top: 0;
  }
  .section {
    padding: 42px 92px 30px 96px;
    background-image: linear-gradient(120.9deg, #f0fffac9 3.1%, #57a68be8 96.5%);
    border-radius: 44px;
    box-shadow: 0px 4px 4px #00000040;
    border-image-slice: 1;
    backdrop-filter: blur(5.5px);
    border-left: solid 2px #408e72c2;
    border-right: solid 2px #408e72c2;
    border-top: solid 2px #408e72c2;
    border-bottom: solid 2px #408e72c2;
  }
  .pos_3 {
    position: absolute;
    left: 450px;
    top: 141px;
  }
  .text {
    color: #000000;
    font-size: 36px;
    font-family: PingFang SC;
    font-weight: 700;
    line-height: 33.5px;
  }
  .text-wrapper {
    margin-right: 6px;
    padding: 12px 0;
    background-color: #e7edeb;
    border-radius: 33px;
    width: 404px;
  }
  .font {
    font-size: 20px;
    font-family: PingFang SC;
    line-height: 18.5px;
    font-weight: 700;
    color: #787878;
  }
  .text_2 {
    margin-left: 26px;
  }
  .text_3 {
    line-height: 18.5px;
  }
  .text_4 {
    line-height: 18.5px;
  }
  .button_3 {
    align-self: center;
  }
  .text-wrapper_2 {
    padding: 14px 0;
    background-color: #000000;
    border-radius: 33px;
    width: 156px;
  }
  .text_5 {
    color: #ffffff;
    line-height: 18.5px;
  }
</style>