<template>
    <div class="report-div">
        <el-form ref="reportForm" :model="formBadRecords">
            <el-form-item label="举报信息：" prop="badRecord" :rows="6" :rules="{ required: true, message: '举报信息是必填项' }">
                <el-input v-model="formBadRecords.badRecord" type="textarea" placeholder="举报信息"></el-input>
            </el-form-item>
            <el-form-item label="照片：" prop="badRecordPhoto">
                <el-upload class="avatar-uploader" ref="upload" action="" :show-file-list="false" :on-change="uploadImage"
                    :before-upload="beforeUpload">
                    <img v-if="badRecordPhoto" :fit="contain" :src="badRecordPhoto" prop="badRecordPhoto" class="avatar" />
                    <el-icon v-else class="avatar-uploader-icon">
                        <Plus />
                    </el-icon>
                </el-upload>
            </el-form-item>
            <el-form-item label="选择用户：" prop="userId" :rules="{required: true,message: '至少选择一个用户'}">
                <el-select v-model="formBadRecords.userId" @click="getUsersList">
                    <el-option v-for="user in usersList" :key="user.userId" :label="user.username" :value="user.userId"
                    >
                    </el-option>
                </el-select>
            </el-form-item>
        </el-form>
        <el-button type="success" @click="handleSubmit">提交</el-button>
    </div>
</template>
<script setup>
import { ElMessage } from 'element-plus';
import {ref, reactive, getCurrentInstance} from 'vue'
const usersList = ref([])
const {proxy} = getCurrentInstance()
const badRecordPhoto = ref('')
const uploadImage = async file => {
    let fd = new FormData()
    fd.append('badRecordPhoto', file.raw)
    let res = await proxy.$api.uploadBadRecord(fd)
    formBadRecords.badRecordPhoto = res
    badRecordPhoto.value = res
}
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
        petPhoto = URL.createObjectURL(file)
        return false
    }
}
const formBadRecords = reactive({
    badRecord: '',
    badRecordPhoto: '',
    userId: 0
})
const handleSubmit = () => {
    proxy.$refs.reportForm.validate(async valid => {
        if(valid){
            console.log(formBadRecords.userId);
            await proxy.$api.addBadRecord(formBadRecords)
            proxy.$refs.reportForm.resetFields()
            proxy.$refs.upload.clearFiles()
            ElMessage.success("提交成功，请等待管理员审核")
        }
        else {
            ElMessage.error("请输入正确的内容")
        }
    })
}
const getUsersList = async () => {
    const res = await proxy.$api.getUsernameList()
    usersList.value = res
}
</script>
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
    width: 300px;
    height: 300px;
    text-align: center;
}
</style>
<style lang="less" scoped>
.avatar-uploader .avatar {
    width: 300px;
    height: 300px;
    display: block;
}
</style>