package main

import (
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var err error

// pets
func petsAddorUpdateUrl(c *gin.Context) {
	var pet1 Pet
	err = unmarshalJSON(&pet1, c)
	if err != nil {
		return
	}
	if pet1.Id != 0 {
		err = updatePet(pet1)
		if err != nil {
			c.JSON(http.StatusOK, errHandle(err))
			return
		}
		c.JSON(http.StatusOK, CommonResult{Code: OK, Message: "更新成功"})

	} else {
		err = createPet(pet1)
		if err != nil {
			c.JSON(http.StatusOK, errHandle(err))
			return
		}
		c.JSON(http.StatusOK, CommonResult{OK, "插入成功", nil})
	}
}
func petsDeleteUrl(c *gin.Context) {
	id1 := c.Param("id")
	id2, err := strconv.Atoi(id1)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	err, deletedRows := deletePet(id2)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	if deletedRows <= 0 {
		c.JSON(http.StatusOK, CommonResult{Code: notFound, Message: "没有id=" + id1 + "的记录"})
	} else {
		c.JSON(http.StatusOK, CommonResult{Code: OK, Message: "删除成功"})
	}
}
func petsFind2Url(c *gin.Context) {
	var conditionQuery1 ConditionQuery1
	err = unmarshalJSON(&conditionQuery1, c)
	if err != nil {
		return
	}
	err, page := findPets(conditionQuery1)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, CommonResult{notFound, "无相关记录", nil})
			return
		}
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "查询成功", page})
}
func petsUploadUrl(c *gin.Context) {
	commonUpload("petPhoto", "pet/", c)
}
func getPetsPhotoUrl(c *gin.Context) {
	err, photos := findPetsPhoto()
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "图片展示成功", photos})
}
func getPetsKindUrl(c *gin.Context) {
	kindList := []string{"狗", "猫", "兔子"}
	c.JSON(http.StatusOK, CommonResult{OK, "查询成功", kindList})
}

// user
func userRegisterUrl(c *gin.Context) {
	var user1 User
	err = unmarshalJSON(&user1, c)
	err = userRegister2(&user1.Account)
	log.Println(user1.Account)
	if err == nil {
		c.JSON(http.StatusOK, CommonResult{unique, "用户已被注册", nil})
		return
	} else if err != sql.ErrNoRows {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	err = userRegister(user1)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "注册成功", nil})
}
func userLoginUrl(c *gin.Context) {
	var user1 User
	err = unmarshalJSON(&user1, c)
	if err != nil {
		return
	}
	if user1.Account == "" || user1.Password == "" {
		c.JSON(http.StatusOK, CommonResult{Code: required, Message: "用户名或密码不能为空"})
		return
	}
	var password string
	err = userLogin(user1.Account, &user1.Id, &user1.Role, &password)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, CommonResult{Code: notFound, Message: "没有此用户"})
			return
		}
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	h := sha1.New()
	io.WriteString(h, user1.Password)
	if hex.EncodeToString(h.Sum(nil)) != password {
		c.JSON(http.StatusOK, CommonResult{Code: wrongPwd, Message: "密码错误"})
		return
	}
	go func(username string,err1 *error){
		*err1 = setLoginTime(username)
	}(user1.Account,&err)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	token, err := generateToken(user1.Id)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	if user1.Role == "admin" {
		var menuList = []Menu{
			{
				"/home",
				"home",
				"首页",
				"house",
				"home/Home"},
			{
				"/pet",
				"pet",
				"宠物管理",
				"user",
				"pet/Pet"},
			{
				"/user",
				"user",
				"用户管理",
				"user",
				"User/User"},
			{
				"/vaccine",
				"vaccine",
				"疫苗管理",
				"user",
				"vaccine/Vaccine"},
			{
				"/vaccineType",
				"vaccineType",
				"疫苗种类管理",
				"user",
				"vaccineType/VaccineType"},
			{
				"/badRecord",
				"badRecord",
				"不良记录展示",
				"user",
				"badRecord/BadRecord"},
			{
				"/examine",
				"examine",
				"管理员审核",
				"user",
				"examine/Examine"}}
		c.JSON(http.StatusOK, CommonResult{OK, "登录成功", Response3{menuList, token}})
	} else {
		var menuList = []Menu{
			{
				"/home",
				"home",
				"首页",
				"house",
				"home/Home"},
			{
				"/pet",
				"pet",
				"宠物管理",
				"user",
				"pet/Pet"},
			{
				"/vaccine",
				"vaccine",
				"疫苗管理",
				"user",
				"vaccine/Vaccine"},
			{
				"/vaccineType",
				"vaccineType",
				"疫苗种类管理",
				"user",
				"vaccineType/VaccineType"},
			{
				"/badRecords",
				"badRecord",
				"不良行为展示",
				"user",
				"badRecord/BadRecord"},
			{
				"/report",
				"report",
				"不良行为举报",
				"user",
				"report/Report"}}
		c.JSON(http.StatusOK, CommonResult{OK, "登录成功", Response3{menuList, token}})
	}
}
func userEditUrl(c *gin.Context) {
	var user1 User
	err = unmarshalJSON(&user1, c)
	if err != nil {
		return
	}
	if user1.Id == 0 {
		c.JSON(http.StatusOK, CommonResult{Code: required, Message: "id为必填字段"})
		return
	}
	err2 := userEdit(user1)
	if err2 != nil {
		c.JSON(http.StatusOK, errHandle(err2))
		return
	}
	c.JSON(http.StatusOK, CommonResult{Code: OK, Message: "编辑成功"})
}
func userDeleteUrl(c *gin.Context) {
	id1 := c.Param("id")
	if id1 == "" {
		c.JSON(http.StatusOK, CommonResult{Code: required, Message: "id为必填字段"})
		return
	}
	id2, err := strconv.Atoi(id1)
	if err != nil {
		c.JSON(http.StatusOK, CommonResult{Code: mustBeNum, Message: "id必须为整数：" + err.Error()})
		return
	}
	token := c.GetHeader("Authorization")
	claims, _ := parseToken(token[7:])
	if claims.UserId == id2 {
		c.JSON(http.StatusOK, CommonResult{Code: loggedIn, Message: "当前登录用户无法删除"})
		return
	}
	err, deletedRows := userDelete(id2)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	if deletedRows <= 0 {
		c.JSON(http.StatusOK, CommonResult{Code: notFound, Message: "没有id=" + id1 + "的记录"})
	} else {
		c.JSON(http.StatusOK, CommonResult{Code: OK, Message: "删除成功"})
	}
}
func findUserUrl(c *gin.Context) {
	var conditionQuery2 ConditionQuery2
	err = unmarshalJSON(&conditionQuery2, c)
	if err != nil {
		return
	}
	err, page := findUser(conditionQuery2)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, CommonResult{notFound, "无相关记录", nil})
			return
		}
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "查询成功", page})
}
func findAllUserUrl(c *gin.Context) {
	err, users := findAllUser()
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
	} else {
		c.JSON(http.StatusOK, CommonResult{OK, "查询成功", users})
	}
}
func userUploadUrl(c *gin.Context) {
	commonUpload("userPhoto", "user/", c)
}

// vaccine
func addVaccineUrl(c *gin.Context) {
	var conditionQuery6 ConditionQuery6
	err = unmarshalJSON(&conditionQuery6, c)
	if err != nil {
		return
	}
	err = addVaccine(conditionQuery6)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "插入成功", nil})
}
func findVaccineByNameUrl(c *gin.Context) {
	var conditionQuery3 ConditionQuery3
	err = unmarshalJSON(&conditionQuery3, c)
	if err != nil {
		return
	}
	err, page := findVaccineByName(conditionQuery3)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "查询成功", page})
}
func findVaccineByPetIdUrl(c *gin.Context) {
	id1 := c.Param("petId")
	if id1 == "" {
		c.JSON(http.StatusOK, CommonResult{Code: required, Message: "id为必填字段"})
		return
	}
	id, err := strconv.Atoi(id1)
	if err != nil {
		c.JSON(http.StatusOK, CommonResult{mustBeNum, "id必须为整数", nil})
		return
	}
	err, vaccines := findVaccineById(id)
	if err != nil && err != sql.ErrNoRows {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "查询成功", vaccines})
}
func deleteVaccineUrl(c *gin.Context) {
	id1 := c.Param("id")
	if len(id1) == 0 {
		c.JSON(http.StatusOK, CommonResult{required, "id为必填字段", nil})
	}
	id2, err := strconv.Atoi(id1)
	if err != nil {
		c.JSON(http.StatusOK, CommonResult{mustBeNum, "id必须为整数", nil})
		return
	}
	err, deletedRows := deleteVaccine(id2)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	if deletedRows <= 0 {
		c.JSON(http.StatusOK, CommonResult{Code: notFound, Message: "没有id=" + id1 + "的记录"})
	} else {
		c.JSON(http.StatusOK, CommonResult{Code: OK, Message: "删除成功"})
	}
}

// badRecord
func addBadRecordsUrl(c *gin.Context) {
	var badRecord BadRecord
	err = unmarshalJSON(&badRecord, c)
	if err != nil {
		return
	}
	err = addBadRecords(badRecord)
	if err != nil {
		c.JSON(OK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "插入成功", nil})
}
func findBadRecordUrl(c *gin.Context) {
	var conditionQuery5 ConditionQuery5
	err = unmarshalJSON(&conditionQuery5, c)
	if err != nil {
		return
	}
	err, page := findBadRecords(conditionQuery5)
	if err != nil && err == sql.ErrNoRows {
		c.JSON(http.StatusOK, CommonResult{notFound, "无相关记录", nil})
		return
	}
	if err != nil {
		c.JSON(OK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "查询成功", page})
}
func badRecordUploadUrl(c *gin.Context) {
	commonUpload("badRecordPhoto", "badRecord/", c)
}
func updateBadRecordUrl(c *gin.Context) {
	examined, err := strconv.ParseBool(c.Query("examined"))
	if err != nil {
		c.JSON(http.StatusOK, CommonResult{mustBeBool, "非布尔型：" + err.Error(), nil})
		return
	}
	id, err := strconv.Atoi(c.Query("badRecordId"))
	if err != nil {
		c.JSON(http.StatusOK, CommonResult{mustBeNum, "id必须为整数：" + err.Error(), nil})
		return
	}
	err = updateBadRecords(id, examined)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "更新成功", nil})
}
func deleteBadRecordsUrl(c *gin.Context) {
	id1 := c.Param("badRecordId")
	if len(id1) == 0 {
		c.JSON(http.StatusOK, CommonResult{required, "id为必填字段", nil})
	}
	id2, err := strconv.Atoi(id1)
	if err != nil {
		c.JSON(http.StatusOK, CommonResult{mustBeNum, "id必须为整数", nil})
		return
	}
	err, deletedRows := deleteBadRecords(id2)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	if deletedRows <= 0 {
		c.JSON(http.StatusOK, CommonResult{Code: notFound, Message: "没有id=" + id1 + "的记录"})
	} else {
		c.JSON(http.StatusOK, CommonResult{Code: OK, Message: "删除成功"})
	}
}

// vaccineType
func getVaccineListUrl(c *gin.Context) {
	err, vaccineTypeList := findAllVaccineType()
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "查询成功", vaccineTypeList})
}
func addVaccineTypeUrl(c *gin.Context) {
	var vaccineType VaccineType
	err = unmarshalJSON(&vaccineType, c)
	if err != nil {
		return
	}
	err = addVaccineType(vaccineType)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "插入成功", nil})
}
func findVaccineTypeUrl(c *gin.Context) {
	var conditionQuery4 ConditionQuery4
	err = unmarshalJSON(&conditionQuery4, c)
	if err != nil {
		return
	}
	err, page := findVaccineType(conditionQuery4)
	if err != nil && err == sql.ErrNoRows {
		c.JSON(http.StatusOK, CommonResult{notFound, "无相关记录", nil})
		return
	}
	if err != nil {
		c.JSON(OK, errHandle(err))
		return
	}
	c.JSON(http.StatusOK, CommonResult{OK, "查询成功", page})
}
func deleteVaccineTypeUrl(c *gin.Context) {
	id1 := c.Param("id")
	if len(id1) == 0 {
		c.JSON(http.StatusOK, CommonResult{required, "id为必填字段", nil})
	}
	id2, err := strconv.Atoi(id1)
	if err != nil {
		c.JSON(http.StatusOK, CommonResult{mustBeNum, "id必须为整数", nil})
		return
	}
	err, deletedRows := deleteVaccineType(id2)
	if err != nil {
		c.JSON(http.StatusOK, errHandle(err))
		return
	}
	if deletedRows <= 0 {
		c.JSON(http.StatusOK, CommonResult{Code: notFound, Message: "没有id=" + id1 + "的记录"})
	} else {
		c.JSON(http.StatusOK, CommonResult{Code: OK, Message: "删除成功"})
	}
}
