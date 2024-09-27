package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-next/models"
    "gorm.io/gorm"
)

type UserController struct {
    DB *gorm.DB
}

func (uc *UserController) GetUsers(c *gin.Context) {
    var users []models.User
    uc.DB.Table("pond-next.users").Find(&users)
    c.JSON(http.StatusOK, users)
}

func (uc *UserController) CreateUser(c *gin.Context) {
    var newUser models.User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    uc.DB.Create(&newUser)
    c.JSON(http.StatusCreated, newUser)
}

// เพิ่มฟังก์ชันสำหรับ PUT และ DELETE ที่นี่
