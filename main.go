package main

import (
	"cricket_stats_api/initializers"
	"cricket_stats_api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	initializers.LoadEnv()

	router := gin.Default()

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
	// 	// AllowHeaders: []string{"Content-Type", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Access-Control-Allow-Methods", "Authorization"},
	// 	AllowHeaders:  []string{"Content-Type", "Authorization"},
	// 	ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin"},
	// }))

	routes.SetRoutes(router)

	router.Run(":8081") // listen and serve on 0.0.0.0:8080
}

// import (
// 	"bufio"
// 	"context"
// 	"cricket_stats_api/player"
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Enter player's first and last name: ")
// 	fullName, err := reader.ReadString('\n')

// 	if err != nil {
// 		fmt.Println("Error reading name:", err)
// 		return
// 	}

// 	fullName = strings.TrimSpace(fullName)
// 	names := strings.Split(fullName, " ")

// 	var firsName, lastName string
// 	if len(names) > 0 {
// 		firsName = names[0]
// 	}
// 	if len(names) > 1 {
// 		lastName = strings.Join(names[1:], " ")
// 	}

// 	url := fmt.Sprintf("https://cricbuzz-cricket.p.rapidapi.com/stats/v1/player/search?plrN=%s%%20%s", firsName, lastName)
// 	Timoutctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	req, err := http.NewRequestWithContext(Timoutctx, "GET", url, nil)
// 	if err != nil {
// 		fmt.Printf("Error creating request: %v\n", err)
// 	}
// 	apiKey := "35a1484e2dmsh245a0be5d118a20p1779d2jsn3c2d4bf768a0"
// 	req.Header.Add("X-RapidAPI-Key", apiKey)
// 	req.Header.Add("X-RapidAPI-Host", "cricbuzz-cricket.p.rapidapi.com")

// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		fmt.Printf("Error perfirming request: %v\n", err)
// 		return
// 	}
// 	defer res.Body.Close()

// 	if res.StatusCode != http.StatusOK {
// 		fmt.Printf("HTTP Request Failed with Status: %s\n", res.Status)
// 		return
// 	}

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		fmt.Printf("Error reading response body: %v\n", err)
// 		return
// 	}

// 	//Unmarshall json data for player
// 	jsonData := string(body)

// 	var Response struct {
// 		Players  []player.Playerinfo `json:"player"`
// 		Category string              `json:"category"`
// 	}

// 	if err := json.Unmarshal([]byte(jsonData), &Response); err != nil {
// 		fmt.Println("Error parsing JSON: ", err)
// 		return
// 	}

// 	for _, p := range Response.Players {
// 		if p.Name == fullName {
// 			fmt.Printf("Player ID :%s, Name: %s\n", p.ID, p.Name)
// 			url1 := fmt.Sprintf("https://cricbuzz-cricket.p.rapidapi.com/stats/v1/player/%s/batting", p.ID)
// 			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 			defer cancel()

// 			req1, err := http.NewRequestWithContext(ctx, "GET", url1, nil)
// 			if err != nil {
// 				fmt.Printf("Error creating request: %v\n", err)
// 			}
// 			apiKey := "35a1484e2dmsh245a0be5d118a20p1779d2jsn3c2d4bf768a0"
// 			req1.Header.Add("X-RapidAPI-Key", apiKey)
// 			req1.Header.Add("X-RapidAPI-Host", "cricbuzz-cricket.p.rapidapi.com")

// 			res1, err := http.DefaultClient.Do(req1)
// 			if err != nil {
// 				fmt.Printf("Error perfirming request: %v\n", err)
// 				return
// 			}
// 			defer res1.Body.Close()

// 			if res1.StatusCode != http.StatusOK {
// 				fmt.Printf("HTTP Request Failed with Status: %s\n", res1.Status)
// 				return
// 			}

// 			body1, err := io.ReadAll(res1.Body)
// 			if err != nil {
// 				fmt.Printf("Error reading response body: %v\n", err)
// 				return
// 			}

// 			var temp interface{}
// 			err1 := json.Unmarshal(body1, &temp)
// 			if err1 != nil {
// 				fmt.Println("Error unmarshaling:", err)
// 				return
// 			}

// 			// MarshalIndent for pretty print
// 			prettyJSON, err1 := json.MarshalIndent(temp, "", "    ") // Prefix "", Indent "    "
// 			if err1 != nil {
// 				fmt.Println("Error marshaling:", err1)
// 				return
// 			}

// 			return prettyJSON;
// 		}
// 	}
// }
