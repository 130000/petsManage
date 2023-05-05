<template>
    <div class="tags">
        <el-tag :key="tag.name" v-for="(tag,index) in tags" :closable="tag.name !== 'home'" :disable-transitions="false"
         type="success"
            :effect="$route.name === tag.name ? 'dark' : 'light'" @click="changeMenu(tag)" @close="close(tag, index)">{{
                tag.label
            }}</el-tag>
    </div>
</template>
<script setup>
import { useRoute, useRouter } from 'vue-router';
import { useStore } from 'vuex';
const router = useRouter()
const store = useStore()
const route = useRoute()
const tags = store.state.tabsList
const changeMenu = item => router.push({ name: item.name })
const close = (tag, index) => {
    let length = tags.length - 1
    //处理vuex中的tabsList
    store.commit("closeTab", tag)
    if (tag.name !== route.name)
        return
    if (index === length)
        router.push({
            name: tags[index - 1].name,
        })
    else {
        router.push({
            name: tags[index].name,
        })
    }
}
</script>
<style lang="less">
.tags {
    padding: 20px;
    width: 100%;

    .el-tag {
        margin-right: 15px;
        cursor: pointer;
    }
}
</style>