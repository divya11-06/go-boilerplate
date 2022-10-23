package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "fmt"
  "rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	r := gin.Default()
      r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
          "message": "pong",
        })
      })
    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
