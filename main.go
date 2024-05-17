package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/samber/lo"
)

// main
//
//	@Description:
func main() {
	if err := runServer(); err != nil {
		slog.Error("Failed to start server!", "details", err.Error())
		os.Exit(1)
	}
}

// runServer
//
//	@Description:
//	@return error
func runServer() error {

	if err := godotenv.Load(); err != nil {
		return err
	}

	port, err := strconv.Atoi(lo.Ternary(os.Getenv("SERVER_PORT") == "",
		"9000", os.Getenv("SERVER_PORT")))
	if err != nil {
		return err
	}

	router := gin.Default()
	router.HTMLRender = &TemplRender{}

	router.Static("/static", "./static")

	loadHandlers(router)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      router,
	}

	slog.Info("Starting server...", "port", port)

	return server.ListenAndServe()
}
