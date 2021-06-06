package items

import (
	"dev-hack-backend/app/db"
	"dev-hack-backend/app/handlers/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(c *gin.Context) {

	_, done := user.ParseBearer(c)
	if done {
		return
	}

	items := db.GetItemsList()

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})

}
