package main

import (
    "QBlockVegMess_Backend_go/api"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    //initialize config
    app := api.Config{Router: router}

    //routes
    app.Routes()

    router.Run(":80")
}
