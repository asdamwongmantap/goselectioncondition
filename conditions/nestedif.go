package condition

import (
	Valuetype "goselectioncondition/struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Nestedif(c *gin.Context) {
	// param input
	var inputvalue Valuetype.Value
	c.BindJSON(&inputvalue)
	var value1 = inputvalue.A
	predicate := ""
	if value1 >= 80 && value1 <= 100 {
		predicate = "Cumlaude"
	} else if value1 >= 50 && value1 <= 79 {
		predicate = "Good"
	} else if value1 >= 0 && value1 <= 49 {
		predicate = "Doesn't Pass"
	} else {
		predicate = "Nothing Predicate"
	}
	c.JSON(http.StatusOK, gin.H{"predicate": predicate})
}
