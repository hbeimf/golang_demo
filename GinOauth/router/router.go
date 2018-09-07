package router

import (
	"log"
	"net/http"
	"os"
	// "time"

	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang_demo/GinOauth/router/middleware"

	"golang_demo/GinOauth/handler"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type User struct {
	UserName  string
	FirstName string
	LastName  string
	Uid       string
}

var identityKey = "id"

func Init() {
	port := os.Getenv("PORT")
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	// the jwt middleware
	authMiddleware, err := middleware.New()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)
	r.POST("/register", handler.RegisterHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/auth")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", handler.HelloHandler)
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
		// auth.GET("/hello", helloHandler)
	}

	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal(err)
	}
}
