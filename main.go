package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)
// If you are using a database, uncomment the below comment block and the block at the 24th line.
/*
func createConn(ctx *context.Context) (*pgx.Conn, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return pgx.Connect(*ctx, os.Getenv("DATABASE_URL"))
}
*/

func main() {
	r := gin.Default()
	ctx := context.Background()
  /*
	conn, err := createConn(&ctx)
  if err != nil {
    log.Fatal(err)
    return
  }
  defer conn.Close()
  queries := database.New(conn)
  */

	// Place your routes here

	// Needed for tailwindcss installation
	r.Static("/static", "./static")

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server %s", err.Error())
	}
}
