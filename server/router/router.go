package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goProject/student/server/handlers"
	"net/http"
)

func InitEngine(mode string) (*gin.Engine, error) {
	//gin.SetMode(mode)

	engine := gin.Default()
	//规则
	//engine.Use(cors.New(cors.Config{
	//	AllowAllOrigins:  true,
	//AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "OPTIONS"},
	//AllowHeaders:     []string{"x-xq5-jwt", "Content-Type", "Origin", "Content-Length"},
	//ExposeHeaders:    []string{"x-xq5-jwt", "Download-Status"},
	//AllowCredentials: true,
	//MaxAge:           12 * time.Hour,
	//}))

	engine.Use(gin.Recovery())

	Client(engine)
	engine.NoRoute(func(c *gin.Context) {
		fmt.Printf("req ip=%s, url:%s", c.ClientIP(), c.Request.URL)
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"error":  "page not found",
		})
	})
	return engine, nil

}

func Client(r *gin.Engine) {
	student := r.Group("/students")

	{
		student.GET("/server/:id", handlers.GetById)
	}

}
