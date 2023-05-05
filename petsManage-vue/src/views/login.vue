<template>
    <div class="body">
        <el-form :model="loginForm" class="login-container">
            <h3>社区宠物可视化管理系统登录</h3>
            <el-form-item>
                <el-input type="input" placeholder="请输入账号" v-model="loginForm.username"></el-input>
            </el-form-item>
            <el-form-item>
                <el-input type="password" placeholder="请输入密码" v-model="loginForm.userPassword"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="success" @click="login">登录</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>
<script setup>
import { reactive, getCurrentInstance } from 'vue';
import { useRouter } from 'vue-router';
import { useStore } from 'vuex';
const { proxy } = getCurrentInstance()
const loginForm = reactive({
    username: 'user2',
    userPassword: ''
})
const store = useStore()
const router = useRouter()
const login = async () => {
    const res = await proxy.$api.login(loginForm)
    console.log(res);
    store.commit('setMenu', res.menu)
    store.commit('addMenu', router)
    store.commit('setToken', res.token)
    router.push({
        name: 'home'
    })
}
</script>
<style lang="less" scoped>
.body {
    width: 100%;
    height: 100%;
    background-color: #d1edc4;
    position: fixed;
}

.login-container {
    width: 350px;
    background-color: #f0f9eb;
    border-radius: 15px;
    padding: 35px 35px 15px 35px;
    margin: 180px auto;

    h3 {
        text-align: center;
        margin-bottom: 20px;
        color: #505450;
    }

    :deep(.el-form-item__content) {
        justify-content: center;
    }
}
</style>