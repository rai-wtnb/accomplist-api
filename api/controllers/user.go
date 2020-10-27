package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rai-wtnb/accomplist-api/crypto"
	"github.com/rai-wtnb/accomplist-api/models/repository"
)

type UserController struct{}

// Index : GET /users
func (UserController) Index(c *gin.Context) {
	var u repository.UserRepository
	r, err := u.GetAll()
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, r)
	}
}

// Signup : POST /users/signup
func (UserController) Signup(c *gin.Context) {
	var u repository.UserRepository
	r, err := u.CreateUser(c)
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, r)
	}
}

// Login : POST /users/login
func (UserController) Login(c *gin.Context) {
	var u repository.UserRepository
	user := u.GetByEmail(c.PostForm("email"))

	dbPassword := user.Password
	formPassword := c.PostForm("password")

	if err := crypto.Verify(dbPassword, formPassword); err != nil {
		c.AbortWithStatus(400)
		fmt.Println("ログイン失敗")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		fmt.Println("ログイン成功")
		c.JSON(200, user)
	}
}

// Show : GET /users/:id
func (UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository
	idInt, _ := strconv.Atoi(id)
	user, err := u.GetByID(idInt)

	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, user)
	}
}

// Update : PUT /users/:id
func (UserController) Update(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository
	idInt, _ := strconv.Atoi(id)
	r, err := u.UpdateByID(idInt, c)

	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, r)
	}
}

// Delete : DELETE /users/:id
func (UserController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository
	idInt, _ := strconv.Atoi(id)
	if err := u.DeleteByID(idInt); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "ユーザーを削除しました"})
	return
}
