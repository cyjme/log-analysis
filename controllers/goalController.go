package controllers

import (
	"github.com/gin-gonic/gin"
	"ideaparLog/model"
	"strconv"
)

type GoalController struct {
}

func (ctrl *GoalController) Index(c *gin.Context) {
	goals := GoalRepo.List()

	c.JSON(200, gin.H{"data": goals})
}

func (ctrl *GoalController) Create(c *gin.Context) {
	goal := model.Goal{
		Id:      0,
		Name:    c.PostForm("name"),
		Desc:    c.PostForm("desc"),
		DataUrl: c.PostForm("dataUrl"),
	}

	GoalRepo.Create(&goal)

	c.Abort()
}

func (ctrl *GoalController) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	goal := model.Goal{
		Id:      id,
		Name:    c.PostForm("name"),
		Desc:    c.PostForm("desc"),
		DataUrl: c.PostForm("dataUrl"),
	}

	GoalRepo.Update(&goal)

	c.Abort()
}

func (ctrl *GoalController) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	GoalRepo.Delete(id)

	c.Abort()
}
