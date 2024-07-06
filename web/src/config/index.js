/*
    项目配置文件
*/

// 开发环境
// eslint-disable-next-line
const ENV_LOCAL={
    // 后端请求地址
    host:"http://localhost:714",
    // 文件服务器地址
    minio:"http://localhost:9000"
}

// 线上环境
// eslint-disable-next-line
const ENV_PRO={
    // 后端请求地址
    host:"http://203.57.227.253:8888",
    // 文件服务器地址
    minio:"http://203.57.227.253:9000"
}

const config=ENV_LOCAL

export default config