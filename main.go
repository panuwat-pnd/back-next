package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "go-next/controllers" // เปลี่ยนให้ตรงกับ path ของโปรเจกต์
    "go-next/models"
)

var db *gorm.DB

func initDB() {
    // ตั้งค่าการเชื่อมต่อฐานข้อมูล
    dsn := "host=localhost user=postgres password= dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Bangkok search_path=pond-next"
    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    // ทำการ Migrate ตาราง User หากยังไม่มี
    db.AutoMigrate(&models.User{})
    db.AutoMigrate(&models.Employee{})
}

func main() {
    initDB()

    router := gin.Default()

    // สร้าง instance ของ UserController
    userController := controllers.UserController{DB: db}
    employeeController := controllers.EmployeeController{DB: db}

    // กำหนดเส้นทางสำหรับการจัดการผู้ใช้
    router.GET("/users", userController.GetUsers)
    router.POST("/users", userController.CreateUser)
    router.GET("/employees", employeeController.GetEmployees)
    router.POST("/employee", employeeController.CreateEmployee)
    // router.PUT("/users/:id", userController.UpdateUser)
    // router.DELETE("/users/:id", userController.DeleteUser)

    router.Run(":8080")
}
