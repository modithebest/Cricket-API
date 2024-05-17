package controllers

import "github.com/gin-gonic/gin"

func HealthCheck(ctx *gin.Context) {
	ctx.String(200, "Hello World")
}
