package models

import (
    "gorm.io/gorm"
)

type Employee struct {
    ID   uint   `json:"id" gorm:"primaryKey"`
    Name string `json:"name"`
    Email  string    `json:"email"`
}

func MigrateEmployees(db *gorm.DB) {
    db.AutoMigrate(&Employee{})
}
