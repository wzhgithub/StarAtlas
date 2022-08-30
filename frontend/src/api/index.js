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

// 根据vmc_id获取vmc具体信息 getVMCData(1)
export const getVMCData = (vmc_id) => {
  return request({
    method: 'GET',
    url: '/api/vmcdata?vmc_id=' + vmc_id
  })
}

// 根据vmc_id和设备的类型获取vmc具体信息 getDeviceData(1, 'cpu')
export const getDeviceData = (vmc_id, device_type) => {
  return request({
    method: 'GET',
    url: `/api/devicedata?vmc_id=${vmc_id}&device_type=${device_type}`
  })
}

// 根据vmc_id获取vmc分区任务 getAppInfo(1)
export const getAppInfo = (vmc_id) => {
  return request({
    method: 'GET',
    url: `/api/appinfo?vmc_id=${vmc_id}`
  })
}

// 根据vmc_id或者折线图数据 getVMCDataSeq(1)
export const getVMCDataSeq = (vmc_id) => {
  return request({
    method: 'GET',
    url: `/api/vmcdata/sequences?vmc_id=${vmc_id}`
  })
}

// 获取topo图的所有节点以及连接方式
export const getTopoShow = () => {
  return request({
    method: 'GET',
    url: `/api/topo/show`
  })
}

// 插入一个node
// var node_json = { "name": "cpu_new", "device_type": "cpu", "parent_id": 2, "upstream_id": 0 }
export const insertNode = (node_json) => {
  return request({
    method: 'POST',
    url: '/api/topo/insert',
    data: JSON.stringify(node_json)
  })
}

// 删除一个node
// var id_json = { "id": 1 }
export const deleteNode = (id_json) => {
  return request({
    method: 'POST',
    url: '/api/topo/delete',
    data: JSON.stringify(id_json)
  })
}

// 任务迁移
// migrate_json = { "from": {"vimd_id": "1"}, "to": {"vimd_id": "2"} }
export const failureOver = (migrate_json) => {
  return request({
    method: 'POST',
    url: '/api/vmc/failure_over',
    data: JSON.stringify(migrate_json)
  })
}
//  处理name非法字符的方法
export const filterName = (name) => {
  let arr = name.split('\u0000')
  return arr[0]
}
// 故障模拟
// migrate_json = { "transType": '', "fromVmcId": 80,"isFault":0,"deviceId":111,"taskName":'',"taskType":'',"appId":100 }
export const doFailureOver = (migrate_json) => {
  return request({
    method: 'POST',
    url: '/api/vmc/do_failure_over',
    data: JSON.stringify(migrate_json)
  })
}
// 轮训结果
// migrate_json = { "from": {"vimd_id": "1"}, "to": {"vimd_id": "2"} }
export const getFailureResult = (migrate_json) => {
  return request({
    method: 'POST',
    url: '/api/vmc/failure_over_result',
    data: JSON.stringify(migrate_json)
  })
}