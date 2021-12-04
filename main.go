package main

import (
	"context"
	"fmt"

	"github.com/alen/project-sabar/routes"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

func main() {

	conn, err := connectDB()
	if err != nil {
		return
	}
	router := gin.Default()

	router.Use(dbMiddleware(*conn))
	usersGroup := router.Group("users")
	{
		usersGroup.POST("/register", routes.UsersRegister)
		usersGroup.POST("/login", routes.UsersLogin)
	}
	router.Run(":8000")

}

func connectDB() (c *pgx.Conn, err error) {
	databaseUrl := "postgres://postgres:alen@localhost:5432/offersapp"
	conn, err := pgx.Connect(context.Background(), databaseUrl)
	if err != nil || conn == nil {
		fmt.Println("Error connecting to DB")
		fmt.Println(err.Error())
	}
	fmt.Println("database Connect Successfull")
	_ = conn.Ping(context.Background())
	return conn, err
}

func dbMiddleware(conn pgx.Conn) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db", conn)
		c.Next()
	}
}
