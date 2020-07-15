package app

import (
	"github.com/gin-gonic/gin"
	"github.com/hoyeah69/bookstore_oauth-api/src/http"
	"github.com/hoyeah69/bookstore_oauth-api/src/repository/db"
	"github.com/hoyeah69/bookstore_oauth-api/src/repository/rest"
	"github.com/hoyeah69/bookstore_oauth-api/src/services/access_token"
)

var (
	router = gin.Default()
)

func StartApplication() {
	// session, dbErr := cassandra.GetSession()
	// if dbErr != nil {
	// 	panic(dbErr)
	// }
	// session.Close()

	atHandler := http.NewAccessTokenHandler(access_token.NewService(rest.NewRestUserRepository(), db.NewRepository()))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
