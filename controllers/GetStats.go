package controllers

import (
	"cricket_stats_api/initializers"

	"github.com/gin-gonic/gin"
)

func GetStats(ctx *gin.Context) {
	supabase := initializers.Supabase()

	var results map[string]interface{}
	err := supabase.DB.From("playersTable").Select("*").Execute(&results)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{

		"result": results,
	})
}
