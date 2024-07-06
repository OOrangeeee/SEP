<template>
  <div class="flex-col page">
    <div class="flex-col justify-start items-start relative group">
      <img
        class="image"
        src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031971002967410.png"
      />
      <img
        class="image_2 pos"
        src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031996792427666.png"
      />
    </div>
    <div class="flex-col justify-start relative group_2 mt-9">
      <div class="flex-row justify-between items-start relative group_3">
        <div class="section_2"></div>
        <img
          class="image_4"
          src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156031441909052164.png"
        />
        <img
          class="image_3 pos_3"
          src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156032035169079570.png"
        />
      </div>
      <div class="section pos_2"></div>
      <img
        class="image_5 pos_4"
        src="https://codefun-proj-user-res-1256085488.cos.ap-guangzhou.myqcloud.com/6641f10b5a7e3f0310455148/6642054ba2432f00114e6290/17156032033007994198.png"
      />
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
  .mt-9 {
    margin-top: 9px;
  }
  .page {
    padding: 27px 180px 60px;
    background-color: #e7edeb;
    width: 100%;
    overflow-y: auto;
    overflow-x: hidden;
    height: 100%;
  }
  .group {
    padding-bottom: 10px;
  }
  .image {
    margin-left: 109px;
    border-radius: 40px;
    width: 84px;
    height: 40px;
  }
  .image_2 {
    width: 1100px;
    height: 50px;
  }
  .pos {
    position: absolute;
    left: 50%;
    bottom: 0;
    transform: translateX(-50%);
  }
  .group_2 {
    padding: 87px 0 46px;
  }
  .group_3 {
    margin-left: 27px;
    margin-right: 55px;
  }
  .section_2 {
    background-image: linear-gradient(223.3deg, #ace4da 10.7%, #ace4da00 87.7%);
    border-radius: 50%;
    width: 562px;
    height: 562px;
  }
  .image_4 {
    margin-top: 355px;
    border-radius: 50%;
    width: 482px;
    height: 482px;
  }
  .image_3 {
    width: 598px;
    height: 390px;
  }
  .pos_3 {
    position: absolute;
    right: 330px;
    top: 191px;
  }
  .section {
    background-image: linear-gradient(5.8deg, #b3e6dd 9.5%, #b3e6dd00 96.6%);
    border-radius: 50%;
    height: 257px;
  }
  .pos_2 {
    position: absolute;
    left: 992px;
    right: 119px;
    top: 0;
  }
  .image_5 {
    border-radius: 50%;
    width: 250px;
    height: 250px;
  }
  .pos_4 {
    position: absolute;
    left: 0;
    bottom: 0;
  }
</style>