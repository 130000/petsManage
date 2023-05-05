<template>
    <div class="user-header">
        <el-form :inline="true">
            <el-form-item>
                <el-select v-model="config.vaccineTypeId" @click="handleSearch" @clear="clearSelect" clearable>
                            <el-option v-for="vaccine in vaccineList" :key="vaccine.vaccineTypeId" :value="vaccine.vaccineTypeId" :label="vaccine.vaccineType">
                            </el-option>
                        </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="success" @click="getVaccineData">查询</el-button>
            </el-form-item>
        </el-form>
    </div>
    <div class="table">
        <el-table :data="list" style="width: 100%" height="500px">
            <el-table-column v-for="item in tableLabel" :key="item.prop" :label="item.label" :prop="item.prop" />
        </el-table>
        <el-pagination background layout="prev, pager, next" :total="config.total" @current-change="changePage"
            class="pager mt-4" />
    </div>
</template>
<script setup>
import { onMounted, getCurrentInstance, ref, reactive } from 'vue';
const { proxy } = getCurrentInstance()
const list = ref([])
const config = reactive({
    total: 0,
    page: 1,
    vaccineTypeId: 0
})
const tableLabel = [
    {
        prop: 'vaccineType.vaccineType',
        label: '疫苗名称'
    },
    {
        prop: 'inoculability',
        label: '接种时间'
    },
    {
        label: '所属宠物',
        prop: 'pet.petName'
    }
]
onMounted(() => {
    getVaccineData(config)
})
const vaccineList = ref([])
const getVaccineData = async () => {
    console.log(typeof config.vaccineTypeId);
    let res = await proxy.$api.findVaccineByName(config)
    list.value = res.records
    config.total = res.total
}
const changePage = page => {
    config.current = page
    getVaccineData(config)
}
const handleSearch = async () => {
    vaccineList.value = await proxy.$api.getVaccineList()
}
const clearSelect = () => {
    config.vaccineTypeId = 0
}
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
}
</style>