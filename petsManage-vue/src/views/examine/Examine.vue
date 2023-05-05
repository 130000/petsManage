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
            <el-table-column fixed="right" label="操作" min-width="180">
                <template #default="scope">
                    <el-button type="danger" @click="handleCancel(scope.row.badRecordId)">删除</el-button>
                    <el-button type="success" @click="handleSuccess(scope.row.badRecordId)">通过</el-button>
                </template>
            </el-table-column>
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
const handleSuccess = async badRecordId => {
    await proxy.$api.updateBadRecord(badRecordId,true)
    getBadRecordData(config)
}
const handleCancel = async badRecordId => {
    await proxy.$api.deleteBadRecord(badRecordId)
    getBadRecordData(config)
}
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
        label: '宠物昵称',
        prop: 'pet.petName'
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
    examined: false
})
const list = ref([])
const changePage = page => {
    config.current = page
    getBadRecordData(config)
}
const getBadRecordData =async config => {
    console.log(config);
    let res = await proxy.$api.getBadRecordList(config)
    config.total = res.total
    list.value = res.records
}
</script>
