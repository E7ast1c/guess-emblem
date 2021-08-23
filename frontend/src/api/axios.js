import axios from 'axios'

axios.defaults.timeout = 10000
axios.defaults.baseURL = 'http://localhost:3000/'

axios.interceptors.request.use(config => {
    return config
}, error => {
    console.log('req error')
    return Promise.reject(error)
})

axios.interceptors.response.use(data => {
    return data
}, error => {
    console.log('res error...')
    return Promise.reject(error)
})

export default axios