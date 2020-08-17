package condition

import (
	Valuetype "goselectioncondition/struct"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Switchcase(c *gin.Context) {
	// param input
	var inputvalue Valuetype.Value
	c.BindJSON(&inputvalue)
	var value1 = inputvalue.A
	month := ""
	switch value1 {
	case 1:
		month = "January"
	case 2:
		month = "February"
	case 3:
		month = "March"
	case 4:
		month = "April"
	case 5:
		month = "May"
	case 6:
		month = "June"
	case 7:
		month = "July"
	case 8:
		month = "August"
	case 9:
		month = "September"
	case 10:
		month = "October"
	case 11:
		month = "November"
	case 12:
		month = "December"
	default:
		month = "The month not found"
	}
	c.JSON(http.StatusOK, gin.H{"month": month})
}
