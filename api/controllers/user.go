package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rai-wtnb/accomplist-api/models"
	"github.com/rai-wtnb/accomplist-api/models/repository"
	"github.com/rai-wtnb/accomplist-api/utils/crypto"
	"github.com/rai-wtnb/accomplist-api/utils/mysession"
	"github.com/rai-wtnb/accomplist-api/utils/s3"
)

type UserController struct{}

// Signup : POST /users/signup
func (UserController) Signup(c *gin.Context) {
	var u repository.UserRepository
	r, err := u.CreateUser(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionID := mysession.NewSessionID()
	u.SaveSession(r.ID, sessionID)

	c.JSON(201, gin.H{
		"sessionID": sessionID,
		"userID":    r.ID,
	})

}

// Login : POST /users/login
func (UserController) Login(c *gin.Context) {
	var u repository.UserRepository
	var err error

	user, err := u.GetByEmail(c.PostForm("email"))
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbPassword := user.Password
	formPassword := c.PostForm("password")

	if err = crypto.Verify(dbPassword, formPassword); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	sessionID := mysession.NewSessionID()
	u.SaveSession(user.ID, sessionID)
	c.JSON(200, gin.H{
		"sessionID": sessionID,
		"userID":    user.ID,
	})
}

// Logout : POST /users/logout
func (UserController) Logout(c *gin.Context) {
	var u repository.UserRepository
	id := c.PostForm("id")

	if err := u.DeleteSession(id); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"logout error": err.Error()})
		return
	}
	c.String(204, "succeeded to logout")
}

// Index : GET /users
func (UserController) Index(c *gin.Context) {
	var u repository.UserRepository

	r, err := u.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, r)
}

// Show : GET /users/:id
func (UserController) Show(c *gin.Context) {
	id := c.Params.ByName("id")
	var err error
	var u repository.UserRepository
	var l repository.ListRepository
	var r repository.RelationRepository

	user, err := u.GetByID(id)
	lists, err := l.GetByUserID(id)
	follows, err := r.GetFollowID(id)
	followers, err := r.GetFollowerID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Lists = append(user.Lists, lists...)
	user.Count.FollowCount = len(follows)
	user.Count.FollowerCount = len(followers)

	c.JSON(200, user)
}

// Update : PUT /users/:id
func (UserController) Update(c *gin.Context) {
	var u repository.UserRepository
	var userAndSession models.UserAndSession
	var err error

	if err = c.BindJSON(&userAndSession); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Params.ByName("id")
	dbSessionID, _ := u.GetSession(id)

	// validation
	if userAndSession.SessionID == "" || dbSessionID != userAndSession.SessionID {
		log.Println("wrong sessionID")
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong sessionID"})
		return
	}

	r, err := u.UpdateByID(id, userAndSession)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, r)
}

// UpdateImg : PUT /users/:id/img
func (UserController) UpdateImg(c *gin.Context) {
	var u repository.UserRepository
	id := c.Params.ByName("id")
	var err error

	img, header, err := c.Request.FormFile("img")
	if err != nil {
		log.Println(img)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// upload to s3
	buf := bytes.NewBuffer(nil)
	_, err = io.Copy(buf, img)
	url, err := s3.Upload(buf.Bytes(), header.Filename)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save url in db
	_, err = u.SaveUrlByID(id, url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, err)
}

// Delete : DELETE /users/:id
func (UserController) Delete(c *gin.Context) {
	id := c.Params.ByName("id")
	var u repository.UserRepository

	err := u.DeleteByID(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(204, gin.H{"success": "deleted the user"})
}
