import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse } from 'axios';
import { showMessage, StatusType } from './status';
import { useStorage } from '@/utils/use-storage';
import { onWarning, onError } from '@/utils/messages';

// 如果请求花费了超过 `timeout` 的时间，请求将被中断
axios.defaults.timeout = 5000;
// 表示跨域请求时是否需要使用凭证
axios.defaults.withCredentials = false;
// 允许跨域
axios.defaults.headers.post['Access-Control-Allow-Origin-Type'] = '*';

const axiosInstance: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_URL + '/', // Proxy 本地开发使用
  // baseURL: 'https://gcloud.aoau.top', // 生产环境使用
  headers: {
    Accept: 'application/json',
    'Content-Type': 'application/json;charset=UTF-8',
  },
});

// axios实例拦截响应
axiosInstance.interceptors.response.use(
  (response: AxiosResponse) => {
    if (response.headers.authorization) {
      localStorage.setItem('token', response.headers.authorization);
      useStorage('token', response.headers.authorization);
    } else if (response.data && response.data.token) {
      localStorage.setItem('token', response.data.token);
      useStorage('token', response.data.token);
    }

    if (!response.status) {
      onError(showMessage(506 as StatusType));
      return Promise.reject(response);
    }

    if (response.status === 200) {
      return Promise.resolve(response);
    }
    onError(showMessage(response.status as StatusType));
    return Promise.reject(response);
  },
  (error: any) => {
    const { response } = error;
    return Promise.reject(showMessage(response.status as StatusType));
  }
);

axiosInstance.interceptors.request.use(
  (config: AxiosRequestConfig) => {
    // config.headers!["Content-Encoding"] = "gzip";
    const token = useStorage('token');
    if (token !== '') {
      config.headers!.Authorization = `${token}`;
    }
    if (config.url === '/refresh/authorization') {
      const refresh_token = useStorage('refresh_token');
      config.headers!.Authorization = `${refresh_token}`;
    }
    return config;
  },
  (error: any) => {
    return Promise.reject(error);
  }
);

export default axiosInstance;
