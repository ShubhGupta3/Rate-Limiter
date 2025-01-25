package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func TokenBucket(ctx *gin.Context) {

	if bucket == 0 {
		ctx.AbortWithStatusJSON(429, "Too many requests, wait for some time")
		return
	}
	bucket--
	fmt.Println("Tokens remaining: ", bucket)
	// ctx.JSON(200, "MIDDLEWARE SAYS HI")
}
