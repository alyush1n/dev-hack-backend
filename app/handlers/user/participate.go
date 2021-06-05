package user

import "github.com/gin-gonic/gin"

func Participate(c *gin.Context)  {
	jsonInput := struct {
		Username  string `json:"username"`

	}
}