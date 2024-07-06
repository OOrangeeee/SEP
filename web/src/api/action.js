import axios from 'axios';
import cookie from 'vue-cookie';
import { message } from 'ant-design-vue';
import router from '@/router';
import config from '@/config';

const host = config.host;

// 创建一个axios实例
const instance = axios.create({
  baseURL: host,
  headers: {
    'Content-Type': 'application/json',
    'X-Csrf-Token': cookie.get('X-Csrf-Token') // 获取并设置CSRF Token
  }
});

// 请求拦截器
instance.interceptors.request.use(
  config => {
    const accessToken = cookie.get('ACCESS-TOKEN');
    if (accessToken) {
      config.headers.Authorization = `Bearer ${accessToken}`; // 添加Authorization Token
    }
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 响应拦截器
instance.interceptors.response.use(
  response => {
    return response.data; // 只返回响应数据部分
  },
  error => {
    if (error.response) {
      // 请求已发出，但服务器响应状态码不在2xx范围内
      const status = error.response.status;
      if (status === 401) {
        // 未授权，可能需要重新登录
        cookie.delete('ACCESS-TOKEN');
        message.error('您还未登录，请先登录', 1);
        router.push({ path: '/login' });
      } else if (status === 403) {
        // 禁止访问
        message.error('您没有权限进行此操作', 1);
      } else if (status === 404) {
        // 记录不存在
        message.error('记录不存在', 1);
      } else if (status === 500) {
        // 服务器错误
        message.error('服务器错误，请稍后再试', 1);
      } else {
        // 其他状态码错误
        message.error('请求错误，请稍后再试', 1);
      }
    } else {
      // 请求未发出，网络错误等
      message.error('网络错误，请稍后再试', 1);
    }
    return Promise.reject(error);
  }
);

// 封装get请求
export function getAction(url) {
  return instance.get(url);
}

// 封装post请求
export function postAction(url, data) {
  return instance.post(url, data);
}
