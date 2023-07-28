package handlers

import (
	"github.com/gin-gonic/gin"
	"goProject/student/server/models/psql"
	"net/http"
	"strconv"
)

func GetById(c *gin.Context) {

	sId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return
	}

	student, err := psql.SelectById(sId)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  1,
		"message": "查询成功",
		"student": student,
	})
}
