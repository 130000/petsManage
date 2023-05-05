
import request from "./request"
export default {

    petPhotoUpload(data){
        return request({
            url: '/pet/upload',
            method: 'post',
            headers: {'Content-Type': 'multipart/form-data'},
            data: data,
        }) 
    },
    addPet(data){
        return request({
            url: '/pet/add',
            method: 'post',
            data,
        })
    },
    updatePet(data){
        return request({
            url: '/pet/update',
            method: 'post',
            data,
        })
    },
    deletePet(id){
        return request({
            url: `/pet/delete/${id}`,
            method: 'get',
        })
    },
    getPetList(data){
        return request({
            url:'/pet/find',
            method:'post',
            data
        })
    },
    addVaccine(data){
        return request({
            url: '/vaccine/add',
            method: 'post',
            data
        })
    },
    findVaccineByName(data){
        return request({
            url: '/vaccine/find',
            method: 'post',
            data
        })
    },
    findVaccineById(id){
        return request({
            url: `/vaccine/find/${id}`,
            method: 'get',
        })
    },
    deleteVaccine(id){
        return request({
            url: `/vaccine/delete/${id}`,
            method: 'get',
        })
    },
    login(data){
        return request({
            url: '/user/login',
            method: 'post',
            data,
        })
    },
    getUserList(data){
        return request({
            url: '/user/find',
            method: 'post',
            data
        })
    },
    getUsernameList(){
        return request({
            url: '/user/find/name',
            method: 'get'
        })
    },
    addUser(data){
        return request({
            url: '/user/add',
            method: 'post',
            data
        })
    },
    updateUser(data){
        return request({
            url: '/user/update',
            method: 'post',
            data
        })
    },
    userUpload(data){
        return request({
            url: '/user/upload',
            method: 'post',
            headers: {'Content-Type': 'multipart/form-data'},
            data: data,
        }) 
    },
    deleteUser(id){
        return request({
            url: `/user/delete/${id}`,
            method: 'get'
        })
    },
    getPetsPhoto(){
        return request({
            url: '/pet/photo',
            method: 'get'
        })
    },
    getPetsList(){
        return request({
            url: '/pet/find',
            method: 'get'
        }) 
    },
    addBadRecord(data){
        return request({
            url: '/badRecord/add',
            method: 'post',
            data
        })
    },
    uploadBadRecord(file){
        return request({
            url: '/badRecord/upload',
            method: 'post',
            headers: {'Content-Type': 'multipart/form-data'},
            data: file,
        })
    },
    getBadRecordList(data){
        return request({
            url: '/badRecord/find',
            method: 'post',
            data
        })
    },
    updateBadRecord(id,examined){
        return request({
            url: `/badRecord/update?badRecordId=${id}&examined=${examined}`,
            method: 'get'
        })
    },
    deleteBadRecord(id){
        return request({
            url: `/badRecord/delete/${id}`,
            method: 'get'
        })
    },
    getPetKindList(){
        return request({
            url: '/pet/kind',
            method: 'get'
        })
    },
    getVaccineList(){
        return request({
            url: '/vaccineType/find',
            method: 'get'
        })
    },
    getVaccineType(data){
        return request({
            url: '/vaccineType/find',
            method: 'post',
            data
        })
    },
    deleteVaccineType(id){
        return request({
            url: `/vaccineType/delete/${id}`,
            method: 'get'
        })
    },
    addVaccineType(data){
        return request({
            url: '/vaccineType/add',
            method: 'post',
            data
        })
    }
}