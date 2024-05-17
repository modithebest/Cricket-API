package controllers

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"cricket_stats_api/player"

	"github.com/gin-gonic/gin"
)

type StatsResponse struct {
	Headers       []string        `json:"headers"`
	Values        []StatsValue    `json:"values"`
	AppIndex      AppIndex        `json:"appIndex"`
	SeriesSpinner []SeriesSpinner `json:"seriesSpinner"`
}

type StatsValue struct {
	Values []string `json:"values"`
}

type AppIndex struct {
	SeoTitle string `json:"seoTitle"`
	WebURL   string `json:"webURL"`
}

type SeriesSpinner struct {
	SeriesName string `json:"seriesName"`
	SeriesId   int    `json:"seriesId,omitempty"`
}

func Statisitic(c *gin.Context) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter player's first and last name: ")
	fullName, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading name:", err)
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
		fmt.Printf("Error creating request: %v\n", err)
	}
	apiKey := "35a1484e2dmsh245a0be5d118a20p1779d2jsn3c2d4bf768a0"
	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "cricbuzz-cricket.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error performing request: %v\n", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("HTTP Request Failed with Status: %s\n", res.Status)
		return
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Unmarshal JSON data for player
	var response struct {
		Players  []player.Playerinfo `json:"player"`
		Category string              `json:"category"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return
	}

	// Check if the player exists
	for _, p := range response.Players {
		if p.Name == fullName {
			fmt.Printf("Player ID: %s, Name: %s\n", p.ID, p.Name)
			url1 := fmt.Sprintf("https://cricbuzz-cricket.p.rapidapi.com/stats/v1/player/%s/bowling", p.ID)
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			req1, err := http.NewRequestWithContext(ctx, "GET", url1, nil)
			if err != nil {
				fmt.Printf("Error creating request: %v\n", err)
			}
			req1.Header.Add("X-RapidAPI-Key", apiKey)
			req1.Header.Add("X-RapidAPI-Host", "cricbuzz-cricket.p.rapidapi.com")

			res1, err := http.DefaultClient.Do(req1)
			if err != nil {
				fmt.Printf("Error performing request: %v\n", err)
				return
			}
			defer res1.Body.Close()

			if res1.StatusCode != http.StatusOK {
				fmt.Printf("HTTP Request Failed with Status: %s\n", res1.Status)
				return
			}

			body1, err := io.ReadAll(res1.Body)
			if err != nil {
				fmt.Printf("Error reading response body: %v\n", err)
				return
			}

			// Return the response body
			c.Data(http.StatusOK, "application/json", body1)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Player not found"})
}
