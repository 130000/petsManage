<template>
    <div class="user-header">
        <el-button type="success" @click="handleAdd">+新增</el-button>
        <el-form :inline="true" :model="formInline">
            <el-form-item label="姓名：">
                <el-input v-model="formInline.keyword" placeholder="请输入姓名" />
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
            <!-- <el-table-column v-for="item in tableLabel" :key="item.prop" :label="item.label" :prop="item.prop"
                :width="item.width ? item.width : 125" /> -->
            <el-table-column :label="tableLabel[0].label" :key="tableLabel[0].prop" :prop="tableLabel[0].prop"
                width="125px"></el-table-column>
            <el-table-column :label="tableLabel[1].label" :key="tableLabel[1].prop" :prop="tableLabel[1].prop"
                width="125px"></el-table-column>
            <el-table-column :label="tableLabel[2].label" :key="tableLabel[2].prop" :prop="tableLabel[2].prop"
                width="125px"></el-table-column>
            <el-table-column :label="tableLabel[3].label" :key="tableLabel[3].prop"
                :prop="tableLabel[3].prop"></el-table-column>
            <el-table-column :label="tableLabel[4].label" :key="tableLabel[4].prop" :prop="tableLabel[4].prop">
                <template #default="scope">
                    <el-image style="width: 175px; height: 175px" :src="scope.row.userPhoto" />
                </template>
            </el-table-column>
            <el-table-column :label="tableLabel[5].label" :key="tableLabel[5].prop"
                :prop="tableLabel[5].prop"></el-table-column>
            <el-table-column :label="tableLabel[6].label" :key="tableLabel[6].prop"
                :prop="tableLabel[6].prop"></el-table-column>
            <el-table-column fixed="right" label="操作" min-width="180">
                <template #default="scope">
                    <el-button type="success" @click="handleEdit(scope.row)">编辑</el-button>
                    <el-button type="danger" @click="handleDelete(scope.row)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
        <el-pagination small background layout="prev, pager, next" :total="config.total" @current-change="changePage"
            class="pager mt-4" />
    </div>
    <el-dialog v-model="dialogVisible" :title="action == 'add' ? '新增用户' : '编辑用户'" width="44%">
        <el-form :inline="true" :model="formUser" class="demo-form-inline" ref="userForm">
            <el-row>
                <el-col :span="12">
                    <el-form-item label="姓名：" prop="name" :rules="[{ required: true, message: '姓名是必填项' }]">
                        <el-input v-model="formUser.name" />
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="年龄：" prop="userAge" :rules="[{ required: true, message: '年龄是必填项' },
                    { type: 'number', message: '年龄必须是数字' }]">
                        <el-input v-model.number="formUser.userAge" prop="userAge" />
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <el-form-item label="性别：" prop="userSex" :rules="[{ required: true, message: '性别是必填项' }]">
                        <el-select v-model="formUser.userSex">
                            <el-option label="男" value="男" />
                            <el-option label="女" value="女"></el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
                <el-col :span="12">
                    <el-form-item label="账号：" prop="username" :rules="[{ required: true, message: '账号是必填项' }]">
                        <el-input v-model="formUser.username" />
                    </el-form-item>
                </el-col>
            </el-row>
            <el-row>
                <el-form-item>
                    <el-form-item label="密码：" prop="userPassword">
                        <el-input v-model="formUser.userPassword" type="password" show-password />
                    </el-form-item>
                </el-form-item>
            </el-row>
            <el-row>
                <el-col :span="12">
                    <el-form-item label="照片：" prop="userPhoto">
                        <el-upload class="avatar-uploader" action="" :show-file-list="false" :on-change="uploadImage"
                            :before-upload="beforeUpload">
                            <img v-if="userPhoto" :src="userPhoto" prop="userPhoto" class="avatar" />
                            <el-icon v-else class="avatar-uploader-icon">
                                <Plus />
                            </el-icon>
                        </el-upload>
                    </el-form-item>
                </el-col>
            </el-row>
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
import { ElMessage } from 'element-plus'
import 'element-plus/es/components/message/style/css'
import { onMounted, getCurrentInstance, ref, reactive } from 'vue';
const { proxy } = getCurrentInstance()
const list = ref([])
const config = reactive({
    total: 0,
    page: 1,
    name: ''
})
const tableLabel = reactive([
    {
        prop: "userId",
        label: "id",
        width: 0
    },
    {
        prop: "name",
        label: "姓名",
    },
    {
        prop: "userAge",
        label: "年龄",
    },
    {
        prop: "userSex",
        label: "性别",
    },
    {
        prop: "userPhoto",
        label: "照片"
    },
    {
        prop: "username",
        label: "账号",
        width: 200,
    },
    {
        prop: "loginTime",
        label: "上次登录时间",
        width: 320,
    },
])
onMounted(() => {
    getUserData(config)
})
const beforeUpload = file => {
    let extensionList = [
        "png",
        "PNG",
        "jpg",
        "JPG",
        "jpeg",
        "JPEG",
        "bmp",
        "BMP",
        "jfif"
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
        userPhoto.value = URL.createObjectURL(file)
        return false
    }
}
const userPhoto = ref('')
const getUserData = async config => {
    let res = await proxy.$api.getUserList(config)
    console.log(res);
    config.total = res.total
    list.value = res.records
}
const changePage = page => {
    config.page = page
    getUserData(config)
}
const formInline = reactive({
    keyword: ''
})
const handleSearch = () => {
    config.name = formInline.keyword
    getUserData(config)
}
const uploadImage = async file => {
    let fd = new FormData()
    fd.append('userPhoto', file.raw)
    let res = await proxy.$api.userUpload(fd)
    formUser.userPhoto = res
    userPhoto.value = res
}
const dialogVisible = ref(false)
const formUser = reactive({
    username: '',
    userAge: 0,
    userSex: '',
    name: '',
    userPassword: '',
    userPhoto: ''
})
const onSumbit = () => {
    proxy.$refs.userForm.validate(async valid => {
        if (valid) {
            if (action.value == 'add') {
                await proxy.$api.addUser(formUser)
                dialogVisible.value = false
                proxy.$nextTick(()=>getUserData(config))
                userPhoto.value = ''
                proxy.$refs.userForm.resetFields()
            } else {
                await proxy.$api.updateUser(formUser)
                dialogVisible.value = false
                proxy.$nextTick(() => getUserData(config))
                userPhoto.value = ''
                proxy.$refs.userForm.resetFields()
            }
        }
        else {
            ElMessage.error('请输入正确的内容')
        }
    })
}
const handleCancel = () => {
    dialogVisible.value = false
    setTimeout(() => {
        userPhoto.value = ''
        proxy.$refs.userForm.resetFields()
    }, 100);
  
}
const action = ref('add')
const handleEdit = row => {
    action.value = 'edit'
    userPhoto.value = row.userPhoto
    dialogVisible.value = true
    proxy.$nextTick(() => Object.assign(formUser, row))
}
const handleAdd = () => {
    action.value = 'add'
    dialogVisible.value = true
}
const handleDelete = async row => {
    const flag = confirm("你确认删除吗？")
    if (flag) {
        await proxy.$api.deleteUser(row.userId)
        getUserData(config)
    }
};

</script>
<style lang="less" scoped>
.avatar-uploader .avatar {
    width: 178px;
    height: 178px;
    display: block;
}

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

.el-icon .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 178px;
    height: 178px;
    text-align: center;
}
</style>