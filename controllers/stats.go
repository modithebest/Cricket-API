package controllers

import (
	"cricket_stats_api/initializers"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GameType struct {
	ID           int    `json:"id" db:"id"`
	FormatTypeID int    `json:"format_type_id" db:"format_type_id"`
	Name         string `json:"name" db:"name"`
}

func Stats(c *gin.Context) {
	supabase := initializers.Supabase()

	// Ensure the FormatTypeID exists
	formatTypeID := 1                             // Example ID, you can change this based on your requirement
	formatTypeIDStr := strconv.Itoa(formatTypeID) // Convert int to string

	var formatTypeExists []struct {
		ID int `json:"id"`
	}
	err := supabase.DB.From("formattype").Select("id").Eq("id", formatTypeIDStr).Execute(&formatTypeExists)
	if err != nil {
		log.Printf("Error checking FormatTypeID existence: %v\n", err)
		c.JSON(500, gin.H{"error": "Error checking FormatTypeID existence: " + err.Error()})
		return
	}

	if len(formatTypeExists) == 0 {
		log.Println("FormatTypeID does not exist")
		c.JSON(400, gin.H{"error": "FormatTypeID does not exist"})
		return
	}

	// Insert into the GameType table
	gameType := GameType{
		ID:           1,
		FormatTypeID: formatTypeID, // Ensure this ID exists in the FormatType table
		Name:         "Test Game Type",
	}

	var results []GameType
	err = supabase.DB.From("GameType").Insert(gameType).Execute(&results)
	if err != nil {
		log.Printf("Error inserting data: %v\n", err)
		c.JSON(500, gin.H{"error": "Error inserting data: " + err.Error()})
		return
	}

	log.Println("Successfully inserted data into GameType")
	c.JSON(201, gin.H{
		"status": "Successfully created",
		"result": gameType,
	})
}
