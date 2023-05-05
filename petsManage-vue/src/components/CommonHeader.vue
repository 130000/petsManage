<script setup>
import { computed } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex'
const store = useStore()
const getImg = () => new URL("../assets/R-C.jfif", import.meta.url).href;
let handleCollapse = () => store.commit("updateIsCollapse")
const current = computed(() => store.state.currentMenu)
const handleLogout = () => {
    store.commit('clearMenu')
    store.commit('clearToken')
    router.push({
        name: 'login'
    })
}
const router = useRouter()
</script>

<template>
    <el-header>
        <div class="l-content">
            <!-- 图标的展示 -->
            <el-button size="small" plain @click="handleCollapse">
                <el-icon :size="20">
                    <Menu />
                </el-icon>
            </el-button>
            <el-breadcrumb separator="/" class="bread">
                <!-- 首页是一定存在的所以直接写死 -->
                <el-breadcrumb-item :to="{ path: '/' }">首页</el-breadcrumb-item>
                <el-breadcrumb-item :to="current.path" v-if="current">{{
                    current.label
                }}</el-breadcrumb-item>
            </el-breadcrumb>
        </div>
        <div class="r-content">
            <el-dropdown>
                <span class="el-dropdown-link">
                    <img class="user" :src="getImg()" alt="" />
                </span>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item>个人中心</el-dropdown-item>
                        <el-dropdown-item @click="handleLogout">退出</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
        </div>
    </el-header>
</template>

<style lang="less" scoped>
header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    width: 100%;
    background-color: black;
}

.r-content .user {
    width: 40px;
    height: 40px;
    border-radius: 50%;
}

.l-content {
    display: flex;
    align-items: center;

    .el-button {
        margin-right: 20px;
    }

    h3 {
        color: #fff;
    }
}

:deep(.bread span) {
    color: #fff !important;
    cursor: pointer;
}
</style>
