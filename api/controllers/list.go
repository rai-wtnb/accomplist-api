package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type ListController struct{}

// Index : GET /lists
func (ListController) Index(c *gin.Context){
	fmt.Println("hello")
}
