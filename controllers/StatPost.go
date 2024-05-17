package controllers

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"cricket_stats_api/initializers"
	"cricket_stats_api/player"

	"github.com/gin-gonic/gin"
)

type PlayerStats struct {
	PlayerID     string          `json:"player_id"`
	Name         string          `json:"name"`
	TeamName     string          `json:"team_name"`
	FaceImageID  string          `json:"face_image_id"`
	Age          int             `json:"age"`
	DateOfBirth  *time.Time      `json:"date_of_birth,omitempty"`
	BattingStats json.RawMessage `json:"batting_stats"`
	BowlingStats json.RawMessage `json:"bowling_stats"`
}

func StatPost(c *gin.Context) {
	supabase := initializers.Supabase()

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter player's first and last name: ")
	fullName, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Error reading name: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading name"})
		return
	}

	fullName = strings.TrimSpace(fullName)
	names := strings.Split(fullName, " ")
	var firstName, lastName string
	if len(names) > 0 {
		firstName = names[0]
	}
	if len(names) > 1 {
		lastName = strings.Join(names[1:], " ")
	}

	url := fmt.Sprintf("https://cricbuzz-cricket.p.rapidapi.com/stats/v1/player/search?plrN=%s%%20%s", firstName, lastName)
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(timeoutCtx, "GET", url, nil)
	if err != nil {
		log.Printf("Error creating request: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
		return
	}

	apiKey := "35a1484e2dmsh245a0be5d118a20p1779d2jsn3c2d4bf768a0"
	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "cricbuzz-cricket.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error performing request: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error performing request"})
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Printf("HTTP request failed with status: %s\n", res.Status)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "HTTP request failed with status: " + res.Status})
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Printf("Error reading response body: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response body"})
		return
	}

	var Response struct {
		Players  []player.Playerinfo `json:"player"`
		Category string              `json:"category"`
	}

	if err := json.Unmarshal(body, &Response); err != nil {
		log.Printf("Error parsing JSON: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing JSON"})
		return
	}

	for _, p := range Response.Players {
		if p.Name == fullName {
			log.Printf("Player ID: %s, Name: %s\n", p.ID, p.Name)
			url1 := fmt.Sprintf("https://cricbuzz-cricket.p.rapidapi.com/stats/v1/player/%s/batting", p.ID)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			req1, err := http.NewRequestWithContext(ctx, "GET", url1, nil)
			if err != nil {
				log.Printf("Error creating request: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating request"})
				return
			}
			req1.Header.Add("X-RapidAPI-Key", apiKey)
			req1.Header.Add("X-RapidAPI-Host", "cricbuzz-cricket.p.rapidapi.com")

			res1, err := http.DefaultClient.Do(req1)
			if err != nil {
				log.Printf("Error performing request: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error performing request"})
				return
			}
			defer res1.Body.Close()

			if res1.StatusCode != http.StatusOK {
				log.Printf("HTTP request failed with status: %s\n", res1.Status)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "HTTP request failed with status: " + res1.Status})
				return
			}

			body1, err := io.ReadAll(res1.Body)
			if err != nil {
				log.Printf("Error reading response body: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response body"})
				return
			}

			log.Printf("Data to insert: %s\n", string(body1))

			var playerStats PlayerStats
			if err := json.Unmarshal(body1, &playerStats); err != nil {
				log.Printf("Error unmarshalling data: %v\n", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error unmarshalling data"})
				return
			}

			playerStats.PlayerID = p.ID
			playerStats.Name = p.Name
			// Set other fields accordingly; date_of_birth is optional
			playerStats.DateOfBirth = nil

			var results []PlayerStats
			err1 := supabase.DB.From("playerstable").Insert(playerStats).Execute(&results)
			if err1 != nil {
				log.Printf("Error inserting data: %v\n", err1)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data into the database: " + err1.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Stats inserted successfully"})
		}
	}
}
