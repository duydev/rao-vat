import axios from 'axios';

const request = axios.create({
  baseURL: 'https://rao-vat-tumhpigaia-as.a.run.app',
  timeout: 3000
});

export default request;
