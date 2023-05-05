package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	createDB()
	r := gin.Default()
	r.Use(Cors())
	setEndpoint(r)
	r.MaxMultipartMemory = 2 << 21 //文件大小设置为10MB
	err2 := r.Run()
	if err2 != nil {
		panic(err2)
	}
}
