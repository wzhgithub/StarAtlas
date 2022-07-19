/**
 * axios请求模块
 */
import request from '@/utils/request'

/**
 * 一言请求模块
 * hitokoto(一言)、en(中英文)、social(社会语录)、soup(毒鸡汤)、fart(彩虹屁)、zha(渣男语录)
 */
export const oneSay = () => {
  return request({
    method: 'GET',
    url: 'https://api.uixsj.cn/hitokoto/get?type=fart&code=json'
  })
}
/**
 * 本地请求模块
 * /name请求名字  /all 请求所有数据
 */
export const nameOrAll = () => {
  return request({
    method: 'GET',
    url: 'http://127.0.0.1:8080/'
  })
}

export const getVMCData = (vmc_id) => {
  return request({
    method: 'GET',
    url: '/api/vmcdata?vmc_id=' + vmc_id
  })
}

export const getDeviceData = (vmc_id, device_type) => {
  return request({
    method: 'GET',
    url: `/api/devicedata?vmc_id=${vmc_id}&device_type=${device_type}`
  })
}

export const getAppInfo = (vmc_id) => {
  return request({
    method: 'GET',
    url: `/api/appinfo?vmc_id=${vmc_id}`
  })
}