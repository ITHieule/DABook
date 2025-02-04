import { createApp } from 'vue'
import App from './App.vue'
import router from './router';

const app = createApp(App);
app.use(router);  // Đăng ký router
app.mount('#app');  // Mount ứng dụng Vue vào phần tử có id="app"