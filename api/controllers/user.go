package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rai-wtnb/accomplist-api/models/repository"
	"github.com/rai-wtnb/accomplist-api/utils/crypto"
	"github.com/rai-wtnb/accomplist-api/utils/mysession"
	"github.com/rai-wtnb/accomplist-api/models"
)

type UserController struct{}

// Signup : POST /users/signup
func (UserController) Signup(c *gin.Context) {
	var u repository.UserRepository
	r, err := u.CreateUser(c)
	if err != nil {
		// session.Set("loginUser", c.PostForm("id"))
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(201, r)
	}
}

// Login : POST /users/login
func (UserController) Login(c *gin.Context) {
	var u repository.UserRepository
	user, err := u.GetByEmail(c.PostForm("email"))
	if err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	dbPassword := user.Password
	formPassword := c.PostForm("password")

	if err := crypto.Verify(dbPassword, formPassword); err != nil {
		c.AbortWithStatus(400)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		sessionID := mysession.NewSessionID()
		u.SaveSession(user.ID, sessionID)
		c.JSON(200, gin.H{
			"sessionID": sessionID,
			"userID":    user.ID,
		})
	}
}

// Logout : POST /users/logout
func (UserController) Logout(c *gin.Context) {
	// var u repository.UserRepository
	userID, _ := c.Cookie("userID")
	log.Println("userID : ", userID)
	// user, err := u.GetByID(id)
	// u.DeleteSession()
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
	var u repository.UserRepository
	var userAndSession models.UserAndSession
	if err := c.BindJSON(&userAndSession); err != nil {
		c.AbortWithStatus(400)
	}

	id := c.Params.ByName("id")
	dbSessionID, err := u.GetSession(id)
	if err != nil {
		c.AbortWithStatus(400)
	}

	// validation
	if dbSessionID == userAndSession.SessionID {
		r, err := u.UpdateByID(id, userAndSession)
		if err != nil {
			log.Println("failed to update")
			c.AbortWithStatus(400)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(200, r)
		}
	} else {
		log.Println("wrong sessionID")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
