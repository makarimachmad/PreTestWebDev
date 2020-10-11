import Vue from 'vue'
import Router from 'vue-router'

import Beranda from './components/Beranda.vue'
import Detail from './components/Detail.vue'
import Akun from './components/Akun.vue'
import Admin from './components/Admin.vue'
import Login from './components/Login.vue'
import Registrasi from './components/Registrasi.vue'

Vue.use(Router)
export default new Router ({
    mode: 'history',
    routes: [
        {
            path:'/',
            name:'beranda',
            component: Beranda
        },
        {
            path:'/detail',
            name:'detail',
            component: Detail
        },
        {
            path:'/akun',
            name:'akun',
            component: Akun
        },
        {
            path:'/admin',
            name:'admin',
            component: Admin
        },
        {
            path:'/registrasi',
            name:'registrasi',
            component: Registrasi
        },
        {
            path:'/login',
            name:'login',
            component: Login
        },
    ]
})