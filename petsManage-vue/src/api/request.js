import axios from 'axios'
import { ElMessage } from 'element-plus'
import 'element-plus/es/components/message/style/css'
import cookie from 'js-cookie'
import config from '../config'
const service = axios.create({
    baseURL: config.baseApi
})
const NETWORK_ERROR = '网络请求异常, 请稍后重试'
//请求之前做的事情
service.interceptors.request.use(res => {
    if(res.url !== '/user/login'){
        res.headers['Authorization'] = cookie.get('token')
    }
    //可以自定义header
    //jwt
    return res
})
service.interceptors.response.use(res => {
    const {code,data,message} = res.data
    if(code === 200){
        return data
    }
    else{
        ElMessage.error(message || NETWORK_ERROR)
        return Promise.reject(message || NETWORK_ERROR)
    }
})
function request(option){
    option.method = option.method || 'get'
    if(option.method.toLowerCase() == 'get'){
        option.params = option.data
    }
    let isMock = config.mock
    //对mock的处理
    if(typeof option.mock !== 'undefined')
        isMock = option.mock
    service.defaults.baseURL = isMock ? config.mockApi : config.baseApi
    return service(option)
}
export default request