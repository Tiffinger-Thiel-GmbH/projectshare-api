/* eslint-disable @typescript-eslint/no-explicit-any */
import axios, { AxiosRequestConfig } from 'axios';

export const get = async (apiPath: string, axiosConfig?: AxiosRequestConfig) => axios.get(apiPath, axiosConfig);
