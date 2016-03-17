package main

import "github.com/gin-gonic/gin"
import "github.com/picsoung/go3scale"

var client = go3scale.New("YOUR_PROVIDER_KEY")
var usage = go3scale.Usage{Name:"hits",Value:1}

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        user_key := c.Query("user_key")
        if(user_key == ""){
          c.JSON(401, gin.H{
              "error": "401 - Unauthorized, missing user_key",
          })
        }else{
          var arr []go3scale.Usage
          arr = append(arr,usage)
          authorized := client.AuthrepUserKey(user_key)
          if(authorized.IsSuccess()){
            c.JSON(200, gin.H{
                "message": "Hello World!",
            })
          }else{
            c.JSON(401, gin.H{
                "message": "401 - Unauthorized",
            })
          }
        }
    })
    r.Run() // listen and server on 0.0.0.0:8080
}
