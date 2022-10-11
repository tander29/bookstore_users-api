package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)

var (
	router = gin.Default()
)

func StartApplication() {
	fmt.Println("started")
	mapUrls()
	err := router.Run(":8088")
	if err != nil {
		fmt.Println("Error starting router:", err)
		os.Exit(1)
	}
}
