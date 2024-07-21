import Vue from 'vue';
import VueRouter from 'vue-router';
import Home from '../pages/home/home.vue';
import Name from '../pages/name/name.vue';
import Login from '../pages/login/login.vue';
import Register from '../pages/register/register.vue';
import Patient_reporter from '../pages/patient_reporter/patient_reporter.vue';
import Image from '../pages/image/image.vue';
import Patient_detail from '../pages/patient_detail/patient_detail.vue';
import Tabbar from '../pages/tabbar/tabbar.vue';
import Detail1 from '../pages/detail1/detail1.vue';
import Detail2 from '../pages/detail2/detail2.vue';
import Detail_download from '../pages/detail_download/detail_download.vue';
import Detail_form from '../pages/detail_form/detail_form.vue';
import Doctor_report_loading from '../pages/doctor_report_loading/doctor_report_loading.vue';
import Doctor_report_prompt from '../pages/doctor_report_prompt/doctor_report_prompt.vue';
import Doctor_report_finish from '../pages/doctor_report_finish/doctor_report_finish.vue';
import Doctor_report_generate from '../pages/doctor_report_generate/doctor_report_generate.vue';
import {Message} from "element-ui"
import {GET_TOKEN} from "../utils/token"
Vue.use(VueRouter);

let routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
  },
  {
    path: '/home',
    name: 'home',
    component: Home,
  },
  {
    path: '/name',
    name: 'name',
    component: Name,
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
  },
  {
    path: '/register',
    name: 'register',
    component: Register,
  },
  {
    path: '/patient_reporter',
    name: 'patient_reporter',
    component: Patient_reporter,
  },
  {
    path: '/image',
    name: 'image',
    component: Image,
  },
  {
    path: '/patient_detail',
    name: 'patient_detail',
    component: Patient_detail,
  },
  {
    path: '/tabbar',
    name: 'tabbar',
    component: Tabbar,
  },
  {
    path: '/detail1',
    name: 'detail1',
    component: Detail1,
  },
  {
    path: '/detail2',
    name: 'detail2',
    component: Detail2,
  },
  {
    path: '/detail_download',
    name: 'detail_download',
    component: Detail_download,
  },
  {
    path: '/detail_form',
    name: 'detail_form',
    component: Detail_form,
  },
  {
    path: '/doctor_report_loading',
    name: 'doctor_report_loading',
    component: Doctor_report_loading,
  },
  {
    path: '/doctor_report_prompt',
    name: 'doctor_report_prompt',
    component: Doctor_report_prompt,
  },
  {
    path: '/doctor_report_finish',
    name: 'doctor_report_finish',
    component: Doctor_report_finish,
  },
  {
    path: '/doctor_report_generate',
    name: 'doctor_report_generate',
    component: Doctor_report_generate,
  },
];

const router = new VueRouter({
  mode: 'hash',
  base: process.env.BASE_URL,
  routes,
});

// 路由拦截器，若用户未登录则跳转到登录页面
router.beforeEach( (to,from,next) =>{
  const token = GET_TOKEN()
  // if(!token){
  //     if(to.name==='/login'){
  //         next()
  //         return
  //     }
  //     console.log("路由拦截器生效")
  //     next("/login")
  //
  //     Message.error("您还未登录，请先登录",1)
  //     return
  // }
  // if(to.name==='/admin') {
  //     // 如果跳转的是管理页面，验证权限
  //     loginCheck()
  //         .then(res => {
  //             console.log(res.data)
  //             if (!res.data.isAdmin) {
  //                 this.$message.error("您不是管理员", 1)
  //                 next(from)
  //             }else {
  //                 next()
  //             }
  //         })
  // }else if(to.name==='/login'){
  //     next('/home')
  // }else {
  //     next()
  // }
  // next()


  if(token){

    //已登录去首页
    if(to.path=='/login'){
      next({path:'/'})
    }else{
      next()
    }
  }
  else{

    //改进后的逻辑
    if(to.path=='/login'){
      //重新触发的时候会进入到这里,因为to.path已经变为/login了，然后直接跳转到登录页
      next()
    }else{
      //第一次会走到这个位置，然后重新触发路由钩子函数beforeEach,
      //当重新触发钩子函数的时候这时的to.path已经变为/login
      if (to.path=='/home' || to.path == '/' || to.path == '/register') {
        next()
      }else {
        next('/login')
        Message.error("您还未登录，请先登录")
      }

    }
  }
  next()
})

export default router
