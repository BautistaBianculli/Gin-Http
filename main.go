package main

import (
	"GORUTINE/handlers"
	"GORUTINE/middleware"
	"os"

	// "fmt"
	"github.com/gin-gonic/gin"

	env "github.com/joho/godotenv"
)

func main() {
	err := env.Load(".env")
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	public := r.Group("/api")
	public.POST("/register",handlers.Register )
	public.POST("/login", handlers.Login)
	
	private := public.Group("/admin")
	private.Use(middleware.JwtAuthMiddleware())
	private.GET("/moovies", handlers.GetMoovies)




	
	port, ok := os.LookupEnv("PORT")
	if !ok {
		panic("PORT DONT SET ON ENV")
	}
	r.Run(port)
}

