

import axios from 'axios';
import {GET_TOKEN,GET_COOKIE,SET_COOKIE} from "../utils/token"
import {Message} from "element-ui"
// 创建axios实例
const service = axios.create({
  baseURL: "/api", // api的base_url
  timeout: 50000 ,// 请求超时时间
   headers: {

     'Content-Type': 'application/x-www-form-urlencoded',

   }
});

// 请求拦截器
service.interceptors.request.use(
  config => {
    // 可以在这里添加请求头等信息
    const token = GET_TOKEN()
    config.headers.Authorization ='Bearer '+ GET_TOKEN()
    config.headers['X-Csrf-Token'] = GET_COOKIE()
    return config;
  },
  error => {
    // 请求错误处理
    console.log(error); // for debug
    Promise.reject(error);
  }
);

// 响应拦截器
service.interceptors.response.use(
  response => {
    if (response.config.url=='/csrf-token'){
      debugger
      SET_COOKIE(response.headers['set-cookie'])
    }
    // 对响应数据做处理，例如只返回data部分
    const res = response.data;
    // 根据返回的状态码做相应处理，例如401未授权等
    return res;
  },
  error => {
    Message.error(error.response.data.error_message)
    // 响应错误处理
    console.log('err' + error); // for debug
    return Promise.reject(error);
  }
);


export default service;
