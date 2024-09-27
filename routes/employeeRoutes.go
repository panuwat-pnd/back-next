package routes

import (
    "github.com/gin-gonic/gin"
    "go-next/controllers"
)

func EmployeeRoutes(router *gin.RouterGroup, uc controllers.EmployeeController) {
    router.GET("/employees", uc.GetEmployees)
    router.POST("/employees", uc.CreateEmployee)
    // เพิ่ม route สำหรับ PUT และ DELETE ที่นี่
}
