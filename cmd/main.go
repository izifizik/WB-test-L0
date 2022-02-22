package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
)

/* TODO
* handlers:
1. POST like pub
parse body -> ch(chanel of struct(model)) (to nats) ??
           -> map[string(uuid)]struct(model) (cache)
           -> insert to psql 				(database)

2. GET /:uuid like sub
ch(chanel of struct(model)) ->??
from cache -> to response

*Logic
Nats-streaming

Cache

repo->cache ??

repo->psql
*/

func main() {
	stanC, err := stan.Connect("testCluster", "testID")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer stanC.Close()
	fmt.Println(&stanC)

	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	router.GET("/", Handle)

	log.Fatal(router.Run("localhost:8080"))
}

func Handle(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func NPubSend() {

}
