package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rai-wtnb/accomplist-api/utils/crypto"
	"github.com/rai-wtnb/accomplist-api/utils/mysession"
	"github.com/rai-wtnb/accomplist-api/models/repository"
)

type UserController struct{}

// Signup : POST /users/signup
func (UserController) Signup(c *gin.Context) {
	var u repository.UserRepository
	r, err := u.CreateUser(c)
	if err != nil {
		session := sessions.Default(c)
		session.Set("loginUser", c.PostForm("id"))
		session.Save()
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, r)
	}
}

// Login : POST /users/login
func (UserController) Login(c *gin.Context) {
	var u repository.UserRepository
	user, err := u.GetByEmail(c.PostForm("email"));
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	dbPassword := user.Password
	formPassword := c.PostForm("password")

	if err := crypto.Verify(dbPassword, formPassword); err != nil {
		c.AbortWithStatus(400)
		log.Println("ログイン失敗")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		sessionID := mysession.NewSessionID()
		session := sessions.Default(c)
		session.Set("sessionID", sessionID)
		session.Set("loginUser", user.ID)
		session.Save()
		// todo
		r := make(map[string]string, 2)
		r["sessionID"] = sessionID
		r["userID"] = user.ID
		c.JSON(http.StatusOK, r)
	}
}

// Logout : POST /users/logout
func (UserController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	// todo
	c.String(http.StatusOK, "ログアウト完了")
}

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
	err := u.DeleteByID(id)
	if err != nil {
		c.AbortWithStatus(403)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "ユーザーを削除しました"})
	return
}
