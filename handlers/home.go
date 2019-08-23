package handlers

import(
  "github.com/gin-gonic/gin"
)

func home(c *gin.Context) {
  c.String(200, "Welcome home")
}
