package routes

import (
    "github.com/gin-gonic/gin"
    "go-next/controllers"
)

func UserRoutes(router *gin.RouterGroup, uc controllers.UserController) {
    router.GET("/users", uc.GetUsers)
    router.POST("/users", uc.CreateUser)
    // เพิ่ม route สำหรับ PUT และ DELETE ที่นี่
}
