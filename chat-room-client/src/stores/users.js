import { ref, computed } from 'vue';
import { defineStore } from 'pinia';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode';
import Cookies from 'js-cookie';

export const useUsersStore = defineStore('counter', () => {
  const userId = ref('');
  const userName = ref('');
  const token = ref('');
  const loginUser = async (email, password) => {
    try {
      const response = await axios.post(
        'http://localhost:9000/api/v1/auth/login',
        { email, password }
      );
      token.value = response.data.data.access_token;
      const decode = jwtDecode(token.value);
      userId.value = decode.id;
      userName.value = decode.username;
      // localStorage.setItem('token', token.value);
      Cookies.set('token', token.value, {
        expires: new Date(decode.exp * 1000),
      });
    } catch (error) {
      console.error('Login failed:', error);
      throw error; // Lempar error untuk ditangani di komponen
    }
  };
  const setToken = () => {
    // const result = localStorage.getItem('token');
    const result = Cookies.get('token');
    if (result === undefined || result === null) {
      console.log('token not found');
    } else {
      token.value = result;
      const decode = jwtDecode(token.value);
      userId.value = decode.id;
      userName.value = decode.username;
    }
  };

  return { userId, userName, token, loginUser, setToken };
});
