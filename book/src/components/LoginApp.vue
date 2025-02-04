<template>
  <div class="login">
    <h2>Đăng Nhập</h2>
    <form @submit.prevent="handleLogin">
      <div>
        <label for="Email">Email</label>
        <input type="email" v-model="email" id="Email" required />
      </div>
      <div>
        <label for="password">Mật khẩu</label>
        <input type="password" v-model="password" id="password" required />
      </div>
      <button type="submit">Đăng nhập</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';  // Import axios
import { jwtDecode } from 'jwt-decode';
export default {
  data() {
    return {
      email: '',
      password: '',
      decoded: ''
    };
  },
  methods: {
    async handleLogin() {
      try {
        // Gửi yêu cầu POST tới API đăng nhập
        const response = await axios.post('http://localhost:8081/api/v1/book/Login', {
          email: this.email,
          password: this.password,
        }, {
          headers: {
            "Content-Type": "application/json",
          },
        });

        // Kiểm tra phản hồi từ server
        if (response.data.code == 200) {
          console.log('Đăng nhập thành công');
          this.decoded = response.data.data
          // Giải mã token JWT từ phản hồi
          const decodedToken = jwtDecode(response.data.data);
          // Lưu token vào localStorage
          localStorage.setItem('idkhachhang', decodedToken.user_id);

          // Chuyển hướng sang trang danh sách sách
          this.$router.push('/order');
        } else {
          console.log('Đăng nhập thất bại:', response.data.message);
        }
      } catch (error) {
        console.error('Lỗi khi đăng nhập:', error);
      }
    },
  },
};
</script>

<style scoped>
/* Thêm CSS cho trang đăng nhập */
.login {
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  background: #f4f4f4;
  border-radius: 8px;
}

form {
  display: flex;
  flex-direction: column;
}

label {
  margin-bottom: 8px;
}

input {
  margin-bottom: 12px;
  padding: 8px;
  border-radius: 4px;
}

button {
  padding: 10px;
  background-color: #4CAF50;
  color: white;
  border: none;
  cursor: pointer;
  border-radius: 4px;
}
</style>
