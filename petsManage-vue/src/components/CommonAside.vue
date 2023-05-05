<template>
    <el-aside :width="!$store.state.isCollapse ? '180px' : '64px'">
        <el-menu class="el-menu-vertical-demo" background-color="#545c64" text-color="#fff"
            :collapse="$store.state.isCollapse" :collapse-transition="false">
            <h3 v-show="!$store.state.isCollapse">社区宠物管理系统</h3>
            <el-menu-item :index="item.path" v-for="item in noChildren()" :key="item.path" @click="clickMenu(item)">
                <component class="icons" :is="item.icon"></component>
                <span>{{ item.label }}</span>
            </el-menu-item>
            <el-sub-menu :index="item.label" v-for="item in hasChildren()" :key="item.path">
                <template #title>
                    <component class="icons" :is="item.icon"></component>
                    <span>{{ item.label }}</span>
                </template>
                <el-menu-item-group title="Group One">
                    <el-menu-item :index="subItem.path" v-for="(subItem, subIndex) in item.children" :key="subIndex">
                        <component class="icons" :is="subItem.icon"></component>
                        <span>{{ subItem.label }}</span>
                    </el-menu-item>
                </el-menu-item-group>
            </el-sub-menu>
        </el-menu>
    </el-aside>

</template>
<script setup>
import { useRouter } from 'vue-router'
import { useStore } from 'vuex';
const router = useRouter()
const store = useStore()
const list = store.state.menu
const noChildren = () => list.filter(item => !item.children)
const hasChildren = () => list.filter(item => item.children)
const clickMenu = item => {
    router.push({
        name: item.name
    })
    store.commit('selectMenu',item)
}
</script>
<style scoped>
.icons {
    width: 18px;
    height: 18px;
}

.el-menu {
    border-right: none;
}

h3 {
    line-height: 48px;
    color: #fff;
    text-align: center;
}
</style>