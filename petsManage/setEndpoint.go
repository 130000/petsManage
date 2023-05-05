package main

import "github.com/gin-gonic/gin"

func setEndpoint(r *gin.Engine) {
	g1 := r.Group("/pet")
	{
		g1.Use(AuthMiddleWare())
		g1.POST("/add", petsAddorUpdateUrl)
		g1.GET("/delete/:id", petsDeleteUrl)
		g1.POST("/update", petsAddorUpdateUrl)
		g1.POST("/find", petsFind2Url)
		g1.POST("/upload", petsUploadUrl)
		g1.GET("/photo", getPetsPhotoUrl)
		g1.GET("/kind", getPetsKindUrl)
	}
	g2 := r.Group("/user")
	{
		g2.POST("/add", AuthMiddleWare(), userRegisterUrl)
		g2.POST("/login", userLoginUrl)
		g2.GET("/delete/:id", AuthMiddleWare(), userDeleteUrl)
		g2.POST("/update", AuthMiddleWare(), userEditUrl)
		g2.POST("/find", AuthMiddleWare(), findUserUrl)
		g2.POST("/upload", AuthMiddleWare(), userUploadUrl)
		g2.GET("/find/name", AuthMiddleWare(), findAllUserUrl)
	}
	g3 := r.Group("/vaccine")
	{
		g3.Use(AuthMiddleWare())
		g3.POST("/add", addVaccineUrl)
		g3.POST("/find", findVaccineByNameUrl)
		g3.GET("/find/:petId", findVaccineByPetIdUrl)
		g3.GET("/delete/:id", deleteVaccineUrl)
	}
	g4 := r.Group("/badRecord")
	{
		g4.Use(AuthMiddleWare())
		g4.POST("/add", addBadRecordsUrl)
		g4.POST("/find", findBadRecordUrl)
		g4.GET("/delete/:badRecordId", deleteBadRecordsUrl)
		g4.POST("/upload", badRecordUploadUrl)
		g4.GET("/update", updateBadRecordUrl)
	}
	g5 := r.Group("/project")
	{
		g5.Use(AuthMiddleWare())
	}
	g6 := r.Group("/vaccineType")
	{
		g6.Use(AuthMiddleWare())
		g6.POST("/find",findVaccineTypeUrl)
		g6.POST("/add",addVaccineTypeUrl)
		g6.GET("/delete/:id",deleteVaccineTypeUrl)
		g6.GET("/find",getVaccineListUrl)
	}
}
