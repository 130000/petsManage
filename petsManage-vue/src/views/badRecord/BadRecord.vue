<template>
    <div class="table">
        <el-table :data="list" style="width: 100%" height="500px">
            <el-table-column :key="tableLabel[0].prop" :label="tableLabel[0].label" :prop="tableLabel[0].prop" width="300px" />
            <el-table-column :label="tableLabel[1].label" :key="tableLabel[1].prop" :prop="tableLabel[1].prop">
                <template #default="scope">
                    <el-image style="width: 175px; height: 175px" :src="scope.row.badRecordPhoto" />
                </template>
            </el-table-column>
            <el-table-column :label="tableLabel[2].label" :key="tableLabel[2].prop" :prop="tableLabel[2].prop">
            </el-table-column>
            <el-table-column :key="tableLabel[3].prop" :label="tableLabel[3].label" :prop="tableLabel[3].prop" width="0px"/>
        </el-table>
        <el-pagination background layout="prev, pager, next" :total="config.total" @current-change="changePage"
            class="pager mt-4" />
    </div>
</template>
<script setup>
import { ref, reactive, onMounted, getCurrentInstance} from 'vue';
const {proxy } = getCurrentInstance()
onMounted(() => {
    getBadRecordData(config)
})
const tableLabel = [
    {
        prop: 'badRecord',
        label: '描述'
    },
    {
        prop: 'badRecordPhoto',
        label: '照片'
    },
    {
        label: '用户昵称',
        prop: 'user.username'
    },
    {
        label:  'id',
        prop: 'badRecordId'
    }
]
const config = reactive({
    total: 0,
    current: 1,
    size: 10,
    examined: true
})
const list = ref([])
const changePage = page => {
    config.current = page
    getBadRecordData(config)
}
const getBadRecordData =async config => {
    console.log(config);
    let res = await proxy.$api.getBadRecordList(config)
    console.log(res);
    config.total = res.total
    list.value = res.records
}
</script>
