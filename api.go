package main

import "github.com/gin-gonic/gin"
import "github.com/picsoung/go3scale"

var client = go3scale.New("PROVIDER_KEY")
var usage = go3scale.Usage{Name:"hits",Value:1}

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        user_key := c.Query("user_key")
        if(user_key == ""){
          c.JSON(403, gin.H{
              "message": "unauthorized",
          })
        }else{
          authorized := client.Authrep_with_user_key(user_key,usage)
          if(authorized.IsSuccess()){
            c.JSON(200, gin.H{
                "message": "hello world",
            })
          }else{
            c.JSON(403, gin.H{
                "message": "unauthorized",
            })
          }
        }
    })
    r.Run() // listen and server on 0.0.0.0:8080
}
