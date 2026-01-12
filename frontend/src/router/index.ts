import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'
import Home from '@/views/Home.vue'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: Home,
  },
  {
    path: '/chart',
    name: 'Chart',
    component: () => import('@/views/Chart.vue'),
  },
  {
    path: '/editor',
    name: 'Editor',
    component: () => import('@/views/Editor.vue'),
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
