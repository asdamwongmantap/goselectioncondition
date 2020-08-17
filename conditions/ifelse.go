package condition

import (
	Valuetype "goselectioncondition/struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ifelse(c *gin.Context) {
	// param input
	var inputvalue Valuetype.Value
	c.BindJSON(&inputvalue)
	var birthskill = inputvalue.B
	gender := ""
	if birthskill == "Yes" {
		gender = "Female"
	} else {
		gender = "Male"
	}
	c.JSON(http.StatusOK, gin.H{"message": "You are a " + gender})
}
