import request from "../utils/request"
export const login = (data)=>request.post('/users/login',data) // 登录
export const register = (data)=>request.post('/users/account',data) // 注册
export const segmentation = (data) => uploadFile('/segmentation',data) // 分割
export const detection = (data) => uploadFile('/detection',data) // 检测
export const track = (data) => uploadFile('/track',data) // 追踪
export const modifyAccount = (data) => request.put('/users/account',data) // 修改用户信息
export const getUserInfo = (data) => request.get('/users/account',data) // 获取用户信息
export const deleteRecord = (id) =>request.delete('/users/records/'+id) // 删除记录
export const getRecords = (data) => request.get('/users/records-all',data) // 获取所有记录
export const getRecordById = (id) => request.get('/users/records/'+id) // 获取记录
export const activationUser = (token) =>request.get('/users/account/activation/'+token) // 激活用户
export const getCsrf = () =>request.get('/csrf-token') // 获取csrf


  const uploadFile=(url,formData)=> {
  // 在这里，我们覆盖了默认的 Content-Type
  return request.post(url, formData, {
    headers: {
      'Content-Type': 'multipart/form-data' // 设置正确的 multipart/form-data Content-Type
    }
  })
    .then(response => {
      debugger
      // 处理响应
      return response
    })
    .catch(error => {
      // 处理错误
      throw error
    });
}


const getCookie = (url) =>{
  return request.get(url,{
    withCredentials:true
  })
    .then(response => {
      console.log(response)
      // 处理响应
      return response

    })
    .catch(error => {
      throw error
    });
}



