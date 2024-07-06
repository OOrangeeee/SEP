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
    <div class="mt-70 flex-row items-center relative group_2">
      <img
        class="image_3"
        src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031898388514407.png"
      />
      <span class="text">xx年x月x日</span>
      <div class="flex-col justify-start items-start group_3 pos_2">
        <img
          class="image_5"
          src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031873027189976.png"
        />
        <div class="flex-col section pos_3">
          <img
            class="self-center image_4"
            src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031895098645094.png"
          />
          <span class="self-center text_2">刘朗清</span>
          <div class="flex-col self-stretch group_4">
            <div class="flex-row items-center">
              <img
                class="image_6"
                src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031895361204081.png"
              />
              <span class="text_3 ml-13">查看往期病史</span>
            </div>
            <img
              class="mt-20 image_7"
              src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031890114620856.png"
            />
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
  .ml-13 {
    margin-left: 13px;
  }
  .page {
    padding: 26px 51px 0 92px;
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
    margin-right: 598px;
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
    left: 222px;
    bottom: 0;
  }
  .group_2 {
    padding-bottom: 935px;
  }
  .image_3 {
    width: 36px;
    height: 36px;
  }
  .text {
    margin-left: 26px;
    color: #000000;
    font-size: 36px;
    font-family: PingFang SC;
    font-weight: 700;
    line-height: 34px;
    letter-spacing: 5px;
  }
  .group_3 {
    padding-top: 83px;
  }
  .pos_2 {
    position: absolute;
    left: 0;
    right: 0;
    top: 0;
    bottom: 0;
  }
  .image_5 {
    border-radius: 64px;
    width: 1133px;
    height: 888px;
  }
  .section {
    padding: 32px 58px 0;
    border-radius: 64px;
    background-color: #ffffff;
  }
  .pos_3 {
    position: absolute;
    right: 0;
    top: 0;
  }
  .image_4 {
    border-radius: 51px;
    width: 106px;
    height: 102px;
  }
  .text_2 {
    margin-top: 30px;
    color: #000000;
    font-size: 36px;
    font-family: Overpass;
    line-height: 33.5px;
  }
  .group_4 {
    margin-top: 52px;
  }
  .image_6 {
    width: 24px;
    height: 24px;
  }
  .text_3 {
    color: #004a30;
    font-size: 20px;
    font-family: Overpass;
    line-height: 18.5px;
  }
  .image_7 {
    width: 350px;
    height: 614px;
  }
</style>