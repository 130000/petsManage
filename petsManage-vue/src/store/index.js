import {createStore} from 'vuex'
import router from '../router'
import cookie from 'js-cookie'
const modules = import.meta.glob('../views/*/*.vue')
export default createStore({
    state: {
        isCollapse: true,
        currentMenu: null,
        tabsList: [
            {
                path: '/',
                name: 'home',
                label: '首页',
                icon: 'home'
            }
        ],
        menu: [],
        token: ''
    },
    mutations: {
        updateIsCollapse(state, payload){
            state.isCollapse = !state.isCollapse
        },
        selectMenu(state,val){
            if (val.name == 'home') {
                state.currentMenu = null
              } else {
                state.currentMenu = val
                let result = state.tabsList.findIndex(item => item.name === val.name)
                result == -1 ? state.tabsList.push(val) : ''
              }
        },
        closeTab(state,val){
            let res = state.tabsList.findIndex(item => item.name === val.name)
            state.tabsList.splice(res,1)
        },
        setMenu(state, val){
            state.menu = val
            localStorage.setItem('menu',JSON.stringify(val))
        },
        clearMenu(state){
            state.menu = []
            state.tabsList = []
            localStorage.removeItem('menu')
        },
        setToken(state,val){
            state.token= val
            let date = new Date()
            date.setMinutes(date.getMinutes() + 30)
            cookie.set('token',val,{expires: date})
        },
        clearToken(state){
            state.token = ''
            cookie.remove('token')
        },
        getToken(state){
            state.token = state.token || cookie.get('token')
        },
        addMenu(state){
            if(!localStorage.getItem('menu')){
                return
            }
            const menu = JSON.parse(localStorage.getItem('menu'))
            state.menu = menu
            const menuArray = []
            menu.forEach(item => {
                if(item.children){
                    item.children = item.children.map(item => {
                        item.component = modules[`../views/${item.url}.vue`]
                        return item
                    })
                    menuArray.push(...item.children)
                }
                else {
                    item.component = modules[`../views/${item.url}.vue`]
                    menuArray.push(item)
                }
            })
            menuArray.forEach(item => router.addRoute('home1',item))
        }
    }
})