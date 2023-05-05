import {createRouter, createWebHashHistory} from 'vue-router'
const routes = [
    {
        path: '/',
        name: 'home1',
        component: () => import('../views/Main.vue'),
        redirect: '/home',
        children: []
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('../views/login.vue')
    }
]
const router = createRouter({
    history: createWebHashHistory(),
    routes
})
export default router