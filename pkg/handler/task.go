package handler

import (
	"github/Rarik88/go_final_project/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) Task(c *gin.Context) {
	id := c.Query("id")
	logrus.Println("Получен запрос на получение задачи с ID:", id)
	task, err := h.api.TaskByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

func (h *Handler) AddTask(c *gin.Context) {
	var task model.Task
	err := c.ShouldBindJSON(&task)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logrus.Printf("in %v %v %v %v", task.Date, task.Title, task.Comment, task.Repeat)

	id, err := h.api.AddTask(task)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) Tasks(c *gin.Context) {
	list, err := h.api.Tasks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

func (h *Handler) UpdateTask(c *gin.Context) {
	var task model.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	logrus.Printf("in %v %v %v %v", task.ID, task.Date, task.Title, task.Comment, task.Repeat)

	_, err = h.api.TaskByID(task.ID)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.api.UpdateTask(task)
	if err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) TaskDone(c *gin.Context) {
	id := c.Query("id")
	logrus.Println("Получен запрос на отметку задачи как выполненной с ID:", id)

	err := h.api.TaskDone(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func (h *Handler) TaskDelete(c *gin.Context) {
	id, _ := c.GetQuery("id")
	err := h.api.TaskDelete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
