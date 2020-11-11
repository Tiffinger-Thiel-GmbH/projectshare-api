/* eslint-disable @typescript-eslint/no-explicit-any */
import axios from 'axios';

export const get = async (apiPath: string) => axios.get(apiPath);
