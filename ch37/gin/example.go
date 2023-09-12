package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/webhook", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, gin.H{
				"message": "Failed to read body",
			})
			return
		}

		// 打印 body 内容
		fmt.Println(string(bodyBytes))

		// 如果您还想继续处理该请求，您可能需要将 body 内容重新写回，因为 ioutil.ReadAll 已经读取了它。
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		c.JSON(200, gin.H{
			"message": "Received",
		})
	})
	r.Run()
}
