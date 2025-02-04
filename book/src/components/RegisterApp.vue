<template>
    <div class="login">
      <h2>Đăng Ký</h2>
      <form @submit.prevent="handleregister">
        <div>
          <label for="Name">Name</label>
          <input type="Text" v-model="Ten" id="Ten" required />
        </div>
        <div>
          <label for="Email">Email</label>
          <input type="email" v-model="email" id="Email" required />
        </div>
        <div>
          <label for="password">Mật khẩu</label>
          <input type="password" v-model="password" id="password" required />
        </div>
        <button type="submit">Đăng Kí</button>
      </form>
    </div>
  </template>
  <script>
  import axios from 'axios';  // Import axios
  
  export default {
    data() {
      return {
        Ten: '',
        email: '',
        password: '',
      };
    },
    methods: {
      async handleregister() {
        try {
          // Gửi yêu cầu POST tới API đăng nhập
          const response = await axios.post('http://localhost:8081/api/v1/book/register', {
            headers: {
              "Content-Type": "application/json", // Ensuring the request content type is set to JSON
            },
          }, {
            Ten: this.Ten,
            email: this.email,
            password: this.password,
          });
          // Kiểm tra phản hồi từ server
          if (response.data.code == 200) {
            console.log('Đăng Ký thành công');
            // Chuyển hướng sang trang danh sách sách nếu đăng nhập thành công
            this.$router.push('/');
          } else {
            console.log('Đăng Ký thất bại:');
          }
        } catch (error) {
          console.error('Lỗi khi đăng Ký:', error);
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