import axios from 'axios';
import Cookies from 'js-cookie';

export const login = async (email, password) => {
  console.log(email);
  const response = await axios.post('http://localhost:9000/api/v1/auth/login', {
    email,
    password,
  });
  console.log(response.data);
  return response.data.data.token;
};
