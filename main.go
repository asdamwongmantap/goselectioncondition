package main

import (
	"log"
	"net/http"

	Ccondition "goselectioncondition/conditions"
	Conf "goselectioncondition/config"

	"github.com/gin-gonic/gin"

	"github.com/rs/cors"
)

func main() {

	addr, err := Conf.MyPort()
	if err != nil {
		log.Fatal(err)
	}

	router := gin.Default()

	v1 := router.Group("/api/v1/goselectioncondition")
	{
		v1.POST("/nestedif", Ccondition.Nestedif)
		v1.POST("/ifelse", Ccondition.Ifelse)
		v1.POST("/switchcase", Ccondition.Switchcase)
	}
	c := cors.AllowAll()

	handler := c.Handler(router)
	log.Fatal(http.ListenAndServe(addr, handler))

}
