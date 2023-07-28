package handlers

import (
	"github.com/gin-gonic/gin"
	"goProject/student/client/grpc"
	"net/http"
	"strconv"
)

func GetAgeById(c *gin.Context) {
	sId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "查询失败，输入数据错误！",
		})
		return
	}
	sAge, err := grpc.GetAgeById(sId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  0,
			"message": "查询失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"message": "查询成功",
		"年龄":      sAge,
	})
}
