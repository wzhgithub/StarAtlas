/*
 * @Author: jacob
 * @Date: 2020-11-25 16:59:00
 * @LastEditTime: 2020-12-02 11:25:33
 * @LastEditors: jacob
 * @Description:请求工具类
 */
 import axios from 'axios'
 import jsonBig from 'json-bigint'
 // 在 JavaScript 模块中直接 import 获取容器即可
 // 这里得到的 store 和你在组件中的 this.$store 是一样一样的
 import store from '@/store'
 
 const request = axios.create({
//    baseURL: '/api/'
 })
 
 // transformResponse 是 axios 专门提供的一个 API
 // 它支持由用户来决定如何转换后端返回的数据
 request.defaults.transformResponse = [function (data) {
   try {
    //  正常的话它使用的 JSON.parse 对数据进行转换
     return JSON.parse(data)
 
     // 这里我们定制使用 json-bigint 这个第三方工具包来帮我们转换
     // 如果转换成功，就直接返回数据
     // 如果转换失败，就进入 catch 返回一个空对象给用户
    //  return jsonBig.parse(data)
   } catch (err) {
     return {}
   }
 }]
 
 // 请求拦截器
//  request.interceptors.request.use(function (config) {
//    // 统一设置 Token
//    const { user } = store.state
//    if (user) {
//      // 后端要求把 token 放到请求头中，使用名字 Authorization 指定
//      // config.headers 用来获取本次的请求头对象，这是 axios 的固定 API
//      // 注意，后端要求的 token 数据格式为：Bearer token数据，要注意 Bearer 后面有一个空格
//      config.headers.Authorization = `Bearer ${user}`
//    }
//    return config
//  }, function (error) {
//    // Do something with request error
//    return Promise.reject(error)
//  })
 
 // 响应拦截器
 request.interceptors.response.use(function (response) {
   // Any status code that lie within the range of 2xx cause this function to trigger
   // Do something with response data
   return response
 }, function (error) {
   // Any status codes that falls outside the range of 2xx cause this function to trigger
   // Do something with response error
   return Promise.reject(error)
 })
 
 export default request