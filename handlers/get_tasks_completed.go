package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (th taskHandler) GetTasksCompleted(c *gin.Context) {
	tasks, err := th.repo.GetAllTasks()
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, tasks)
}
