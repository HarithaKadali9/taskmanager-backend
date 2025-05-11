package database

import (
    "gorm.io/gorm"
    "github.com/glebarez/sqlite"  // ✅ This is the pure Go driver
    "taskmanager/models"          // ✅ Ensure this import matches your folder structure
)

var DB *gorm.DB

func Connect() {
    db, err := gorm.Open(sqlite.Open("tasks.db"), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    db.AutoMigrate(&models.Task{})
    DB = db
}
