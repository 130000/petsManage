package main

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	notFound     = iota + 1 //没找到
	required                //必填
	wrongPwd                //密码错误
	uploadFail              //上传失败啊
	unique                  //用户已注册
	mustBeNum               //必须为数字
	unauthorized            //权限不足
	loggedIn                //用户已登录
	mustBeBool              //必须为布尔型
	unknown                 //未知
	OK           = 200
)

type Records interface {
	[]Pet | []User | []Response4 | []Response6 | []VaccineType
}
type G1 interface {
	Pet | User | BadRecord | ConditionQuery3 |
		ConditionQuery1 | ConditionQuery2 | ConditionQuery4 |
		ConditionQuery5 | VaccineType | ConditionQuery6
}
type Pet struct {
	Id       int         `json:"petId"`
	Name     string      `json:"petName"`
	Photo    string      `json:"petPhoto"`
	Age      int         `json:"petAge"`
	Feature  string      `json:"petFeature"`
	Kind     string      `json:"petKind"`
	Type     string      `json:"petType"`
	Vaccines []Response1 `json:"vaccine"`
}
type ConditionQuery1 struct {
	Kind    string `json:"petKind"`
	Name    string `json:"petName"`
	Feature string `json:"petFeature"`
	Current int    `json:"current"`
	Size    int    `json:"size"`
}
type ConditionQuery2 struct {
	Name    string `json:"name"`
	Current int    `json:"current"`
	Size    int    `json:"size"`
	Total   int    `json:"total"`
}
type ConditionQuery3 struct {
	Id      int `json:"vaccineTypeId"`
	PetId 	int `json:"petId"`
	Current int `json:"current"`
	Size    int `json:"size"`
	Total   int `json:"total                                                                                                 "`
}
type ConditionQuery4 struct {
	Size    int `json:"size"`
	Total   int `json:"total"`
	Current int `json:"current"`
}
type ConditionQuery5 struct {
	Size     int  `json:"size"`
	Total    int  `json:"total"`
	Current  int  `json:"current"`
	Examined bool `json:"examined"`
}
type ConditionQuery6 struct {
	VaccineId     int       `json:"vaccineId"`
	VaccineTypeId int       `json:"vaccineTypeId"`
	Inoculability time.Time `json:"inoculability"`
	PetId         int       `json:"petId"`
}
type MyCustomClaims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}
type User struct {
	Id        int        `json:"userId"`
	Name      string     `json:"name"`
	Age       int        `json:"userAge"`
	Account   string     `json:"username"`
	Password  string     `json:"userPassword"`
	Sex       string     `json:"userSex"`
	Role      string     `json:"userRole"`
	Photo     string     `json:"userPhoto"`
	LoginTime *time.Time `json:"loginTime"`
}
type VaccineType struct {
	Id   int    `json:"vaccineTypeId"`
	Type string `json:"vaccineType"`
}
type Vaccine struct {
	Id            int         `json:"vaccineId"`
	VaccineType   VaccineType `json:"vaccineType"`
	Inoculability time.Time   `json:"inoculability"`
	PetId         int         `json:"petId"`
}
type BadRecord struct {
	Id        int    `json:"badRecordId"`
	BadRecord string `json:"badRecord"`
	Photo     string `json:"badRecordPhoto"`
	UserId    int    `json:"userId"`
}
type PetProject struct {
	Id            int       `json:"petProjectId"`
	CreateTime    time.Time `json:"petProjectCreateTime"`
	ProjectTypeId int       `json:"projectTypeId"`
	PetId         int       `json:"petId"`
	Price         float32   `json:"petProjectPrice"`
	Detail        string    `json:"petProjectDetail"`
	Address       string    `json:"petProjectAddress"`
}
type Response1 struct {
	VaccineType   VaccineType `json:"vaccineType"`
	Inoculability time.Time   `json:"inoculability"`
}
type CommonResult struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
type Page[T Records] struct {
	Current int `json:"current"`
	Size    int `json:"size"`
	Total   int `json:"total"`
	Records T   `json:"records"`
}
type Response2 struct {
	Account string `json:"account"`
	Role    string `json:"userRole"`
}
type Response4 struct {
	Id            int         `json:"vaccineId"`
	VaccineType   VaccineType `json:"vaccineType"`
	Inoculability time.Time   `json:"inoculability"`
	Pet           Response5   `json:"pet"`
}
type Menu struct {
	Path  string `json:"path"`
	Name  string `json:"name"`
	Label string `json:"label"`
	Icon  string `json:"icon"`
	Url   string `json:"url"`
}
type Response3 struct {
	Menu  []Menu `json:"menu"`
	Token string `json:"token"`
}
type Response5 struct {
	Id   int    `json:"petId"`
	Name string `json:"petName"`
}
type Response6 struct {
	Id        int       `json:"badRecordId"`
	Photo     string    `json:"badRecordPhoto"`
	BadRecord string    `json:"badRecord"`
	User      Response7 `json:"user"`
}
type Response7 struct {
	Id   int    `json:"userId"`
	Name string `json:"username"`
}

//	func getPetDefault() Pet{
//		return Pet{Id: -1, Age: -1}
//	}
//
//	func getUserDefault() User{
//		return User{Id: -1, Age: -1}
//	}
func errHandle(err error) CommonResult {
	log.Println(err)
	return CommonResult{unknown, err.Error(), nil}
}
