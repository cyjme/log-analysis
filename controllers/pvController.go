package controllers

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

type PvController struct {
}

func (ctrl *PvController) List(c *gin.Context) {
	query := c.Query("query")
	fmt.Println(query)

	pvs, err := PvRepo.List(query)
	if err != nil {
		c.Error(err)
	}

	c.JSON(200, gin.H{"data": pvs})
}
