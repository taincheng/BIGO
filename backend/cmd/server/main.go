package main

import (
	"BIGO/backend/internal/config"
	"BIGO/backend/internal/handlers"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	if err := config.Init(); err != nil {
		panic(err)
	}
	router := setupRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", config.Cfg.Server.Port),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/ping", handlers.Ping)
	return router
}
