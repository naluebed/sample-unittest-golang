package handlers

import(
  "github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
  r := gin.Default()
	r.GET("/home", home)
  return r
}
