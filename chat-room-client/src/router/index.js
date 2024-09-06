import { createRouter, createWebHistory } from 'vue-router';
import HomeView from '../views/HomeView.vue';
import { useUsersStore } from '@/stores/users';
import { computed } from 'vue';

const isAuthenticated = () => {
  const store = useUsersStore();
  const token = store.token;
  return !!token;
};

const isTokenAlready = () => {
  const store = useUsersStore();
  const token = store.token;
  return !!token;
};

// Fungsi untuk memeriksa apakah pengguna adalah admin
const isAdmin = () => {
  // Implementasikan logika pemeriksaan peran pengguna di sini
  return localStorage.getItem('userRole') === 'admin'; // Contoh sederhana
};

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/login',
      name: 'login',
      component: () => import('../views/LoginView.vue'),
      meta: { tokenAlready: true, requiresAuth: false, requiresAdmin: false }, // Meta data untuk route ini
    },
    {
      path: '/room/:roomId/chat',
      name: 'chat',
      component: () => import('../views/ChatView.vue'),
      meta: { tokenAlready: false, requiresAuth: true, requiresAdmin: false }, // Meta data untuk route ini
    },
    {
      path: '/room',
      name: 'room',
      component: () => import('../views/RoomView.vue'),
      meta: { tokenAlready: false, requiresAuth: true, requiresAdmin: false }, // Meta data untuk route ini
    },
    {
      path: '/:pathMatch(.*)*',
      component: () => import('../views/NotFound.vue'),
    }, // Rute 404
  ],
});

// Route guard sebelum setiap navigasi
router.beforeEach((to, from, next) => {
  const requiresAuth = to.meta.requiresAuth;
  const requiresAdmin = to.meta.requiresAdmin;
  const tokenAlready = to.meta.tokenAlready;

  if (requiresAuth && !isAuthenticated()) {
    // Jika route memerlukan otentikasi dan pengguna belum otentikasi, redirect ke halaman login
    next({ name: 'login' });
  } else if (requiresAdmin && !isAdmin()) {
    // Jika route memerlukan admin dan pengguna bukan admin, redirect ke halaman home atau halaman lain
    next({ name: 'home' });
  } else if (tokenAlready && isTokenAlready()) {
    next({ name: 'home' });
  } else {
    // Izinkan navigasi ke route yang diminta
    next();
  }
});

export default router;
