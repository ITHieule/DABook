import { createRouter, createWebHistory } from 'vue-router';
import LoginApp from '@/components/LoginApp.vue';
import registerApp from '@/components/RegisterApp.vue';
import Bookstor from '@/components/Bookstor.vue';
import OderApp from '@/components/OderApp.vue';

const routes = [
  {
    path: '/',
    name: 'loginApp',
    component: LoginApp,
  },
  {
    path: '/order',
    name: 'OderApp',
    component: OderApp,
  },
  {
    path: '/register',
    name: 'registerApp',
    component: registerApp,
  },
  {
    path: '/listbook',
    name: 'listbook',
    component: Bookstor,
  },
  // Có thể thêm các route khác ở đây
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
