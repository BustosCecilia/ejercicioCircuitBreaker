package main

import (
    "github.com/gin-gonic/gin"
    "github.com/mercadolibre/ejercicioCircuitBreaker/src/api/controllers/myml"
    "github.com/mercadolibre/ejercicioCircuitBreaker/src/api/controllers/ping"
)

const (
    port = ":8084"
)

var (
    router = gin.Default()
)
//http://localhost:8082/myml/1
func main() {
    router.GET("/ping", ping.Ping)
    router.GET("/myml/:userID", myml.GetUser)
    router.Run(port)

}
