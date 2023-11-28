import { createRouter, createWebHistory } from 'vue-router'

const routes = [{
  path: '/my-admin',
  redirect: '/login'
},
{
  path: '/init',
  name: 'Init',
  component: () => import('@/view/init/index.vue')
},
{
  path: '/login',
  name: 'Login',
  component: () => import('@/view/login/index.vue')
},
{
  path: '/:catchAll(.*)',
  meta: {
    closeTab: true,
  },
  component: () => import('@/view/error/index.vue')
},
{
  path: '/layout-et1/:companyPath',
  name: 'LayoutEt1',
  meta: {
    closeTab: true,
  },
  component: () => import('@/components/external/LayoutOne.vue')
}
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
