package handlers

import (
    "net/http"
    "taskmanager/database"
    "taskmanager/models"
    "github.com/gin-gonic/gin"
)

// GET /tasks
func GetTasks(c *gin.Context) {
    var tasks []models.Task
    database.DB.Find(&tasks)
    c.JSON(http.StatusOK, tasks)
}

// POST /tasks
func CreateTask(c *gin.Context) {
    var task models.Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if task.Title == "" || task.Status == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Title and Status are required"})
        return
    }

    database.DB.Create(&task)
    c.JSON(http.StatusCreated, task)
}

// PUT /tasks/:id
func UpdateTask(c *gin.Context) {
    var task models.Task
    id := c.Param("id")

    if err := database.DB.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
        return
    }

    var input models.Task
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    task.Title = input.Title
    task.Description = input.Description
    task.Status = input.Status
    task.DueDate = input.DueDate

    database.DB.Save(&task)
    c.JSON(http.StatusOK, task)
}

// DELETE /tasks/:id
func DeleteTask(c *gin.Context) {
    id := c.Param("id")
    if err := database.DB.Delete(&models.Task{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
