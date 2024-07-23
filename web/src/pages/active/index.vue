<template>

</template>
<script >
  import {activationUser} from "../../api/index"
  export default {
    async created() {
      const token = this.$route.params.token
      console.log(token)
      console.log("==========")
       activationUser(token).then(()=>{
         this.$router.push('/')
       }).catch((error)=>{
         if(error.response.data.error_message=='用户已激活') {
           this.$router.push('/')
         }
       })

    }
  }
</script>
<style scoped>

</style>
