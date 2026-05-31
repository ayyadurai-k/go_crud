package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"crud/controllers"
	"crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

func main() {
	router := gin.Default()

	// Liveness/readiness probe for the reverse proxy and container healthcheck.
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Temporary diagnostic: reports the real database connection status/error in
	// the HTTP body so we can see WHY the container can't reach the DB.
	router.GET("/dbcheck", func(c *gin.Context) {
		if initializers.DB == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"db": "nil (ConnectDB failed at startup)"})
			return
		}
		sqlDB, err := initializers.DB.DB()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"stage": "DB()", "error": err.Error()})
			return
		}
		if err := sqlDB.Ping(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"stage": "ping", "error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"db": "connected"})
	})

	router.POST("/posts", controllers.PostCreate)
	router.GET("/posts", controllers.PostList)
	router.GET("/posts/:id", controllers.PostDetail)
	router.PUT("/posts/:id", controllers.PostUpdate)
	router.DELETE("/posts/:id", controllers.PostDelete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	// Serve in a goroutine so we can listen for shutdown signals below.
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server failed: %v", err)
		}
	}()
	log.Printf("listening on :%s", port)

	// Block until SIGINT/SIGTERM (e.g. `docker compose down`, deploys, k8s).
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down...")

	// Give in-flight requests up to 10s to finish.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("forced shutdown: %v", err)
	}
	log.Println("server exited cleanly")
}
