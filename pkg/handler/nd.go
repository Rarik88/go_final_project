package handler

import (
	nextDate "github/Rarik88/go_final_project/pkg/data"
	"github/Rarik88/go_final_project/pkg/model"
	"log"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NextD(c *gin.Context) {
	var nd model.NextDate

	err := c.ShouldBindQuery(&nd)
	if err != nil {
		log.Println(err)
		c.String(400, err.Error())
		return
	}
	str, err := nextDate.NextDate(nd)
	if err != nil {
		log.Println(err)
		c.String(400, err.Error())
		return
	}
	c.String(200, str)
}
