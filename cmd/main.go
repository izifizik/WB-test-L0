package main

import (
	"WB-test-L0/internal/app"
	"github.com/gin-gonic/gin"
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
	err := app.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Handle(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "hello",
	})
}

func NPubSend() {

}
