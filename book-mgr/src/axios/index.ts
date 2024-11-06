import axios from 'axios';
// import useUserInfoStore from "@/stores/useUserInfoStore";

// axios 实例
const instance = axios.create( {
    baseURL: 'http://localhost:7001',
    // baseURL: 'http://gateway.orb.local:8081',
    timeout: 10000 // 设置超时时间
});

// // 请求拦截器
// instance.interceptors.request.use(config => {
//     const token = sessionStorage.getItem('token');
//     console.log(token)
//     if (token) {
//         config.headers.Authorization = `Bearer ${token}`;
//     }
//     return config;
// }, error => {
//     return Promise.reject(error);
// });
//
// // 响应拦截器
// instance.interceptors.response.use(response => {
//     return response;
// }, error => {
//     if (error.response && error.response.status === 401) {
//         let userInfoStore = useUserInfoStore();
//         // token 过期
//         console.error('Token expired or invalid. Please log in again.');
//         // 登出操作
//         userInfoStore.logout()
//
//     }
//     return Promise.reject(error);
// });

export default instance;
