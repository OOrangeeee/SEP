//封装本地存储存储数据与读取数据方法
//存储数据
export const SET_TOKEN = (token) => {
    localStorage.setItem('TOKEN', token)
}
//本地存储获取数据
export const GET_TOKEN = () => {
    return localStorage.getItem('TOKEN')
}
//本地存储删除数据方法
export const REMOVE_TOKEN = () => {
    localStorage.removeItem('TOKEN')
}


export const SET_COOKIE =  (cookie) =>{
    localStorage.setItem('COOKIE', cookie)
}

export const GET_COOKIE =  () =>{
   return localStorage.getItem('COOKIE')
}

