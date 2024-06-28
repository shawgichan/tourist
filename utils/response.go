package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func SuccessResponse(data interface{}) gin.H {
	return gin.H{"Message": "success", "data": data}
}

func SuccessResponseWithCount(data interface{}, count interface{}) gin.H {
	return gin.H{"Message": "success", "Total": count, "data": data}
}
