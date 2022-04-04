package api_web

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/smartschool/repository"
// )

// func GetTeacherCourse(c *gin.Context) {
// 	params := c.Request.URL.Query()

// 	result, err := repository.SearchAttendance(params)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Cannot get data"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, result)
// }
