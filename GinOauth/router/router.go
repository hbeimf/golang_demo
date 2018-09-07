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

func Init() {
	port := os.Getenv("PORT")
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	if port == "" {
		port = "8000"
	}

	// the jwt middleware
	authMiddleware, err := middleware.New()

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	router.POST("/login", authMiddleware.LoginHandler)
	router.POST("/register", handler.RegisterHandler)

	router.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	adminGroup := router.Group("/api/admin")
	adminGroup.Use(authMiddleware.MiddlewareFunc())
	{
		adminGroup.GET("/hello", handler.HelloHandler)
		adminGroup.GET("/refresh_token", authMiddleware.RefreshHandler)
		// adminGroup.GET("/hello", helloHandler)
	}

	clientGroup := router.Group("/api/client")
	clientGroup.Use(authMiddleware.MiddlewareFunc())
	{
		clientGroup.GET("/hello", handler.HelloHandler)
		clientGroup.GET("/refresh_token", authMiddleware.RefreshHandler)
		// adminGroup.GET("/hello", helloHandler)
	}

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
