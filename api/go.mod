module github.com/rai-wtnb/accomplist-api

go 1.15

require (
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/sessions v0.0.3
	github.com/gin-gonic/gin v1.6.3
	github.com/jinzhu/gorm v1.9.16
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897
)

replace (
	github.com/rai-wtnb/accomplist-api/controllers => ./controllers
	github.com/rai-wtnb/accomplist-api/utils/mysession => ./utils/mysession
	github.com/rai-wtnb/accomplist-api/utils/crypto => ./utils/crypto
	github.com/rai-wtnb/accomplist-api/db => ./db
	github.com/rai-wtnb/accomplist-api/models => ./models
	github.com/rai-wtnb/accomplist-api/models/repository => ./models/repository
	github.com/rai-wtnb/accomplist-api/server => ./server
)
