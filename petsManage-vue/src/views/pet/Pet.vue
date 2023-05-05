<template>
    <div class="user-header">
        <el-button type="success" @click="handleAdd">+新增</el-button>
        <el-form :inline="true" :model="formInline">
            <el-form-item>
                <el-input v-model="formInline.petName" placeholder="请输入昵称" />
            </el-form-item>
            <el-form-item>
                <el-input v-model="formInline.petFeature" placeholder="请输入特征"></el-input>
            </el-form-item>
            <el-form-item>
                <el-select v-model="formInline.petKind" @click="getKindList" clearable>
                    <el-option v-for="petKind in petKindList" :key="petKind" :value="petKind" :label="petKind">
                    </el-option>
                </el-select>
            </el-form-item>
            <el-form-item>
                <el-button type="success" @click="handleSearch">查询</el-button>
            </el-form-item>
        </el-form>
    </div>
    <div class="table">
        <el-table :data="list" style="width: 100%" height="500px">
            <!-- <el-table-column v-for="item in tableLabel" :key="item.prop" :label="item.label" :prop="item.prop"
                        :width="item.width ? item.width : 125" /> -->
            <el-table-column :label="tableLabel[0].label" :key="tableLabel[0].prop" :prop="tableLabel[0].prop"
                width="125px"></el-table-column>
            <el-table-column :label="tableLabel[1].label" :key="tableLabel[1].prop" :prop="tableLabel[1].prop"
                width="125px"></el-table-column>
            <el-table-column :label="tableLabel[2].label" :key="tableLabel[2].prop" :prop="tableLabel[2].prop">
                <template #default="scope">
                    <el-image style="width: 175px; height: 175px" :src="scope.row.petPhoto" />
                </template>
            </el-table-column>
            <el-table-column :label="tableLabel[3].label" :key="tableLabel[3].prop" :prop="tableLabel[3].prop"
                width="125px"></el-table-column>
            <el-table-column :label="tableLabel[4].label" :key="tableLabel[4].prop"
                :prop="tableLabel[4].prop"></el-table-column>
            <el-table-column :label="tableLabel[5].label" :key="tableLabel[5].prop" :prop="tableLabel[5].prop"
                width="125px"></el-table-column>
            <el-table-column :label="tableLabel[6].label" :key="tableLabel[6].prop" :prop="tableLabel[6].prop"
                width="125px"></el-table-column>
            <el-table-column fixed="right" label="操作" min-width="180">
                <template #default="scope">
                    <el-button type="success" @click="handleEdit(scope.row)">编辑</el-button>
                    <el-button type="danger" @click="handleDelete(scope.row.petId)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-pagination small background layout="prev, pager, next" :total="config.total" @current-change="changePage"
            class="pager mt-4" />
    </div>
    <el-dialog :draggable="true" :show-close="false" v-model="dialogVisible" :title="action == 'add' ? '新增宠物' : '编辑宠物'"
        width="44%">
        <el-form :inline="true" :model="formPet" class="demo-form-inline" ref="petForm">
            <el-row>
                <el-col :span="12">
                    <el-form-item label="昵称：" prop="petName" :rules="{ required: true, message: '昵称是必填项' }">
                        <el-input v-model="formPet.petName" />
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="年龄：" prop="petAge" :rules="[{ required: true, message: '年龄是必填项' },
                    { type: 'number', message: '年龄必须是数字' }]">
                        <el-input v-model.number="formPet.petAge" prop="petAge" />
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <el-form-item label="照片：" prop="petPhoto">
                        <el-upload class="avatar-uploader" action="" :show-file-list="false" :on-change="uploadImage"
                            :before-upload="beforeUpload">
                            <img v-if="petPhoto" :src="petPhoto" prop="petPhoto" class="avatar" />
                            <el-icon v-else class="avatar-uploader-icon">
                                <Plus />
                            </el-icon>
                        </el-upload>
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="特征：" prop="petFeature">
                        <el-input v-model.number="formPet.petFeature" prop="petFeature" />
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <el-form-item label="种类" prop="petKind" :rules="{ required: true, message: '宠物种类不能为空' }">
                        <el-select v-model="formPet.petKind" @click="getKindList">
                            <el-option v-for="petKind in petKindList" :key="petKind" :value="petKind" :label="petKind">
                            </el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="品种：" prop="petType">
                        <el-input v-model="formPet.petType" />
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row>
                <el-form-item>
                    <div class="tags">
                        <el-tag type="success" v-for="(e, index) in vaccineList" :key="index" closable
                            @close="handleTagClose(e, vaccineList)">{{ e.vaccineType.vaccineType }}</el-tag>
                    </div>
                </el-form-item>
            </el-row>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button type="success" v-if="action == 'edit'" @click="handleVaccineAdd">添加疫苗信息</el-button>
                <el-button type="danger" @click="handleCancel">取消</el-button>
                <el-button type="success" @click="onSumbit">
                    确定
                </el-button>
            </span>
        </template>
        <el-dialog label="新增疫苗信息" :show-close="false" v-model="dialog2Visible" width="40%" append-to-body>
            <el-form :inline="true" class="el-form2" :model="formVaccine" ref="vaccineForm">
                <el-form-item :label="table2Label[0]" prop="vaccineTypeId" :rules="{ required: true, message: '名称是必填项' }">
                    <el-select v-model="formVaccine.vaccineTypeId" placeholder="请选择名称" @click="getVaccineList">
                        <el-option v-for="vaccineType in vaccineTypeList" :key="vaccineType.vaccineTypeId"
                            :value="vaccineType.vaccineTypeId" :label="vaccineType.vaccineType">
                        </el-option>
                    </el-select>
                </el-form-item>
                <el-form-item :label="table2Label[1]" prop="inoculability" :rules="{ required: true, message: '日期是必填项' }">
                    <el-date-picker v-model="formVaccine.inoculability" type="datetime"
                        placeholder="请选择日期和时间"></el-date-picker>
                </el-form-item>
            </el-form>
            <template #footer>
                <span class="dialog-footer">
                    <el-button type="danger" @click="handleDialog2Cancel">取消</el-button>
                    <el-button type="success" @click="onDialog2Sumbit(formPet.petId)">
                        确定
                    </el-button>
                </span>
            </template>
        </el-dialog>
    </el-dialog>
</template>
<script setup>
import { ElMessage } from 'element-plus';
import { onMounted, getCurrentInstance, ref, reactive } from 'vue';
const { proxy } = getCurrentInstance()
const list = ref([])
const config = reactive({
    total: 0,
    current: 1,
    size: 10,
    petName: '',
    petFeature: '',
    petType: ''
})
onMounted(() => {
    getPetData(config)
})
const vaccineTypeList = ref([])
let vaccineList = ref([])
const petKindList = ref([])
const getVaccineList = async () => {
    vaccineTypeList.value = await proxy.$api.getVaccineList()
}
const getKindList = async () => petKindList.value = await proxy.$api.getPetKindList()
function beforeUpload(file) {
    let extensionList = [
        "png",
        "PNG",
        "jpg",
        "JPG",
        "jpeg",
        "JPEG",
        "bmp",
        "BMP",
        "jfif",
        "webp"
    ]
    let index = file.name.lastIndexOf('.')
    let extension = file.name.subStr(index + 1)
    if (extensionList.indexOf(extension) < 0) {

        ElMessage.error('当前格式文件不支持')
        return false
    }
    else if (file.size / 1024 / 1024 < 10) {
        ElMessage.error('图片不能超过10MB')
        return false
    } else {
        petPhoto.value = URL.createObjectURL(file)
        return false
    }
}
async function uploadImage(file) {
    let fd = new FormData()
    fd.append('petPhoto', file.raw)
    let res = await proxy.$api.petPhotoUpload(fd)
    formPet.petPhoto = res
    petPhoto.value = res
}
const petPhoto = ref('')
const table2Label = [
    "疫苗名称",
    "接种时间"
]
const tableLabel = [
    {
        prop: "petId",
        label: "id"
    },
    {
        prop: "petName",
        label: "昵称",
    },
    {
        prop: "petPhoto",
        label: "照片",
    },
    {
        prop: "petAge",
        label: "年龄",
    },
    {
        prop: "petFeature",
        label: "特征",
        width: 200,
    },
    {
        prop: 'petKind',
        label: '种类'
    },
    {
        prop: "petType",
        label: "品种"
    }
]

const getPetData = async config => {
    let res = await proxy.$api.getPetList(config)
    config.total = res.total
    list.value = res.records
}
const changePage = page => {
    config.current = page
    getPetData(config)
}
const formInline = reactive({
    petName: '',
    petKind: '',
    petFeature: ''
})
const handleSearch = () => {
    config.petFeature = formInline.petFeature
    config.petName = formInline.petName
    config.petKind = formInline.petKind
    getPetData(config)
}
const handleVaccineAdd = () => {
    dialog2Visible.value = true
}
const dialogVisible = ref(false)
const dialog2Visible = ref(false)
const formPet = reactive({
    petId: 0,
    petName: '',
    petPhoto: '',
    petAge: -1,
    petFeature: '',
    petKind: '',
    petType: ''
})
const formVaccine = reactive({
    vaccineId: 0,
    vaccineTypeId: 0,
    inoculability: '',
    petId: -1
})
const onDialog2Sumbit = petId =>
    proxy.$refs.vaccineForm.validate(async valid => {
        if (valid) {
            formVaccine.petId = petId
            await proxy.$api.addVaccine(formVaccine)
            dialog2Visible.value = false
            vaccineList.value = await proxy.$api.findVaccineById(petId)
            proxy.$nextTick(() => {
                getPetData(config)
            })
            setTimeout(() => {
                proxy.$refs.vaccineForm.resetFields()
            }, 300);

        }
        else {
            ElMessage.error('请输入正确的内容')
        }
    })
const onSumbit = () => {
    proxy.$refs.petForm.validate(async valid => {
        if (valid) {
            if (action.value == 'add') {
                await proxy.$api.addPet(formPet)
            }
            else {
                await proxy.$api.updatePet(formPet)
            }
            dialogVisible.value = false
            proxy.$nextTick(() => getPetData(config))
            setTimeout(() => {
            petPhoto.value = ''
              proxy.$refs.petForm.resetFields()  
            }, 300);
            
        }
        else {
            ElMessage.error('请输入正确的内容')
        }
    })
}
const handleCancel = () => {
    dialogVisible.value = false
    setTimeout(() => {
        petPhoto.value = ''
        proxy.$refs.petForm.resetFields()
    }, 300);

}
const handleDialog2Cancel = () => {
    dialog2Visible.value = false
    proxy.$refs.vaccineForm.resetFields()
}
const action = ref('add')
const handleEdit = async row => {
    action.value = 'edit'
    petPhoto.value = row.petPhoto
    let res = await proxy.$api.findVaccineById(row.petId)
    console.log(res);
    vaccineList.value = res
    dialogVisible.value = true
    //object.assign从一个或多个源对象复制到目标对象，返回修改后的对象
    proxy.$nextTick(() => Object.assign(formPet, row))
}
const handleAdd = () => {
    action.value = 'add'
    dialogVisible.value = true
}
// async function handleDelete(row) {
//     const flag = window.confirm("确认删除?")
//     if (flag) {
//         await proxy.$api.deletePet(row.petId) 
//         getPetData(config)
//     } else {
//         return -1
//     }
// };
const handleTagClose = async (e, vaccineList) => {
    for (let i = 0; i < vaccineList.length; i++) {
        if (vaccineList[i] == e) {
            await proxy.$api.deleteVaccine(e.vaccineId)
            vaccineList.splice(i, 1)
            break
        }
    }

}
const handleDelete = async id => {
    const flag = confirm("确认删除吗？")
    if (flag) {
        await proxy.$api.deletePet(id)
        getPetData(config)
    }
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

.cell-item {
    display: flex;
    align-items: center;
}

.user-header {
    display: flex;
    justify-content: space-between;
}

.margin-top {
    margin-top: 20px;
}

.avatar-uploader .avatar {
    width: 178px;
    height: 178px;
    display: block;
}

.dialog-footer {
    display: flex;
    justify-content: flex-end;
}
</style>
<style lang="less">
.avatar-uploader .el-upload {
    border: 1px dashed var(--el-border-color);
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
    transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
    border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
}
</style>
