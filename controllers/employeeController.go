package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "go-next/models"
    "gorm.io/gorm"
)

type EmployeeController struct {
    DB *gorm.DB
}

func (uc *EmployeeController) GetEmployees(c *gin.Context) {
    var employees []models.Employee
    uc.DB.Table("pond-next.employees").Find(&employees)
    c.JSON(http.StatusOK, employees)
}

func (uc *EmployeeController) CreateEmployee(c *gin.Context) {
    var newEmployee models.Employee
    if err := c.ShouldBindJSON(&newEmployee); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    uc.DB.Create(&newEmployee)
    c.JSON(http.StatusCreated, newEmployee)
}

// เพิ่มฟังก์ชันสำหรับ PUT และ DELETE ที่นี่
