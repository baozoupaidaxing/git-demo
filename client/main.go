package main

import (
	"github.com/gin-gonic/gin"
	"goProject/student/client/handlers"
)

func main() {
	r := gin.Default()
	r.GET("/students/client/:id", handlers.GetAgeById)
	r.Run(":8081")
}
