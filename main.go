package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Place your routes here
  
  // Needed for tailwindcss installation
	r.Static("/static", "./static")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server %s", err.Error())
	}
}
