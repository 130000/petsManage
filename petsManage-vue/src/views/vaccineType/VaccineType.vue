<template>
    <div class="user-header">
        <el-button type="success" @click="handleAdd">+新增</el-button>
    </div>
    <div class="table">
        <el-table :data="vaccineTypeList" style="width: 100%" height="500px">
            <el-table-column v-for="item in tableLabel" :key="item.prop" :label="item.label" :prop="item.prop" />
            <el-table-column fixed="right" label="操作" min-width="180">
                <template #default="scope">
                    <el-button type="danger" @click="handleDelete(scope.row.vaccineTypeId)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-pagination background layout="prev, pager, next" :total="config.total" @current-change="changePage"
            class="pager mt-4" /> 
    </div>
    <el-dialog :draggable="true" :show-close="false" v-model="dialogVisible" title="新增疫苗种类" width="44%">
        <el-form :inline="true" :model="formVaccineType" class="demo-form-inline" ref="vaccineTypeForm">
            <el-form-item label="疫苗种类：" prop="vaccineType" :rules="{ required: true, message: '种类是必填项' }">
                <el-input v-model="formVaccineType.vaccineType" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button type="danger" @click="handleCancel">取消</el-button>
                <el-button type="success" @click="onSumbit">
                    确定
                </el-button>
            </span>
        </template>
    </el-dialog>
</template>
<script setup>
import { onMounted, getCurrentInstance, ref, reactive } from 'vue';
const { proxy } = getCurrentInstance()
const config = reactive({
    total: 0,
    page: 1
})
const tableLabel = [
    {
        prop: 'vaccineTypeId',
        label: '疫苗种类id'
    },
    {
        prop: 'vaccineType',
        label: '疫苗种类'
    }
]
onMounted(() => {
    getVaccineTypeData(config)
})
const vaccineTypeList = ref([])
const getVaccineTypeData = async config => {
    let res = await proxy.$api.getVaccineType(config)
    vaccineTypeList.value = res.records
}
const handleCancel = () => {
    dialogVisible.value = false
}
const onSumbit = () => {
    proxy.$refs.vaccineTypeForm.validate(async valid => {
        if (valid) {
            await proxy.$api.addVaccineType(formVaccineType)
            dialogVisible.value = false
            proxy.$nextTick(() => getVaccineTypeData(config))
            proxy.$refs.userForm.resetFields()
        }
        else {
            ElMessage.error('请输入正确的内容')
        }
    })
}
const dialogVisible = ref(false)
const handleAdd = () => {
    dialogVisible.value = true
}
const handleDelete = async id => {
    const confirm = window.confirm("你确定删除吗？")
    if (confirm) {
        await proxy.$api.deleteVaccineType(id)
        getVaccineTypeData(config)
    }
}
const formVaccineType = reactive({
    vaccineType: ''
})
</script>
<style lang="less" scoped>
.table {

    position: relative;
    height: 520px;

    .pager {
        position: absolute;
        right: 0;
        bottom: -20px;
    }
}

.user-header {
    display: flex;
    justify-content: space-between;
    margin-bottom:20px;
}
</style>