package controllers

import (
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
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

// IndexID : GET /users/ids
func IndexID(c *gin.Context) {
	var u repository.UserRepository
	r, err := u.GetAll()
	ids := make([]string, 0)
	for _, s := range r {
		ids = append(ids, s.ID)
	}
	if err != nil {
		c.AbortWithStatus(404)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, ids)
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
		log.Println("ログイン失敗")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		session := sessions.Default(c)
		session.Set("loginUser", user.ID)
		session.Save()
		c.String(http.StatusOK, "ログイン完了")
	}
}

// Logout : POST /users/logout
func (UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.String(http.StatusOK, "ログアウト完了")
}

// Show : GET /users/:id
func (UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository
	user, err := u.GetByID(id)

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
	r, err := u.UpdateByID(id, c)

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
	if err := u.DeleteByID(id); err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "ユーザーを削除しました"})
	return
}
