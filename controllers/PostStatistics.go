package controllers

import (
	"bufio"
	"context"
	"cricket_stats_api/initializers"
	"cricket_stats_api/player"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type InputJSON struct {
	Headers []string      `json:"headers"`
	Values  []InputValues `json:"values"`
}

type InputValues struct {
	Values []string `json:"values"`
}

type FormatStatsBatting struct {
	PlayerID       string
	Name           string
	Matches        string `json:"Matches"`
	Innings        string `json:"Innings"`
	Runs           string `json:"Runs"`
	Balls          string `json:"Balls"`
	Highest        string `json:"Highest"`
	Average        string `json:"Average"`
	SR             string `json:"SR"`
	NotOut         string `json:"Not Out"`
	Fours          string `json:"Fours"`
	Sixes          string `json:"Sixes"`
	Ducks          string `json:"Ducks"`
	Fiftys         string `json:"50s"`
	Hundreds       string `json:"100s"`
	Two_hundreds   string `json:"200s"`
	Three_hunderds string `json:"300s"`
	Four_hundreds  string `json:"400s"`
}

type FormatStatsBowling struct {
	PlayerID           string
	Name               string
	Matches            string `json:"Matches"`
	Innings            string `json:"Innings"`
	Balls              string `json:"Balls"`
	Runs               string `json:"Runs"`
	Maidens            string `json:"Maidens"`
	Wickets            string `json:"Wickets"`
	Average            string `json:"Avg"`
	Economy            string `json:"Eco"`
	StrikeRate         string `json:"SR"`
	BestBowlingInnings string `json:"BBI"`
	BestBowlingMatch   string `json:"BBM"`
	FourWickets        string `json:"4w"`  // "4w" is not a valid identifier, hence using "FourWickets"
	FiveWickets        string `json:"5w"`  // "5w" is not a valid identifier, hence using "FiveWickets"
	TenWickets         string `json:"10w"` // "10w" is not a valid identifier, hence using "TenWickets"
}

type Batting struct {
	Test FormatStatsBatting `json:"Test"`
	ODI  FormatStatsBatting `json:"ODI"`
	T20  FormatStatsBatting `json:"T20"`
	IPL  FormatStatsBatting `json:"IPL"`
}

type BattingID struct {
	TestbatID int `json:"Test"`
	ODIbatID  int `json:"ODI"`
	T20batID  int `json:"T20"`
	IPLbatID  int `json:"IPL"`
}

type BowlingStats struct {
	Test FormatStatsBowling `json:"Test"`
	ODI  FormatStatsBowling `json:"ODI"`
	T20  FormatStatsBowling `json:"T20"`
	IPL  FormatStatsBowling `json:"IPL"`
}
type GameTypes struct {
	ID           int    `json:"id" db:"id"`
	FormatTypeID int    `json:"format_type_id" db:"format_type_id"`
	Name         string `json:"name" db:"name"`
}

func Post(c *gin.Context) {
	// // Parse the request body into playerStats
	// if err := c.BindJSON(&playerStats); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// Transform the playerStats if necessary to fit your database schema

	// Insert into the database

	supabase := initializers.Supabase()

	playerID, err := PlayerName()
	fmt.Println("PlayerID: " + playerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	battingStatsJSON, err := fetchStats(playerID, "batting")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println("Attempting to insert batting stats:", string(battingStatsJSON))

	battingStats, err := transformBattingStats(battingStatsJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println("battingStats struct:", battingStats.Test.Matches)

	// var battingStats Batting
	// if err := json.Unmarshal(battingStatsJSON, &battingStats); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// fmt.Printf("Populated Batting Stats: %+v\n", battingStats)

	battingRows := []FormatStatsBatting{
		{
			PlayerID:       playerID,
			Name:           "Test",
			Matches:        battingStats.Test.Matches,
			Innings:        battingStats.Test.Innings,
			Runs:           battingStats.Test.Runs,
			Balls:          battingStats.Test.Balls,
			Highest:        battingStats.Test.Highest,
			Average:        battingStats.Test.Average,
			SR:             battingStats.Test.SR,
			NotOut:         battingStats.Test.NotOut,
			Fours:          battingStats.Test.Fours,
			Sixes:          battingStats.Test.Sixes,
			Ducks:          battingStats.Test.Ducks,
			Fiftys:         battingStats.Test.Fiftys,
			Hundreds:       battingStats.Test.Hundreds,
			Two_hundreds:   battingStats.Test.Two_hundreds,
			Three_hunderds: battingStats.Test.Three_hunderds,
			Four_hundreds:  battingStats.Test.Four_hundreds,
		},
		{
			PlayerID:       playerID,
			Name:           "ODI",
			Matches:        battingStats.ODI.Matches,
			Innings:        battingStats.ODI.Innings,
			Runs:           battingStats.ODI.Runs,
			Balls:          battingStats.ODI.Balls,
			Highest:        battingStats.ODI.Highest,
			Average:        battingStats.ODI.Average,
			SR:             battingStats.ODI.SR,
			NotOut:         battingStats.ODI.NotOut,
			Fours:          battingStats.ODI.Fours,
			Sixes:          battingStats.ODI.Sixes,
			Ducks:          battingStats.ODI.Ducks,
			Fiftys:         battingStats.ODI.Fiftys,
			Hundreds:       battingStats.ODI.Hundreds,
			Two_hundreds:   battingStats.ODI.Two_hundreds,
			Three_hunderds: battingStats.ODI.Three_hunderds,
			Four_hundreds:  battingStats.ODI.Four_hundreds,
		},
		{
			PlayerID:       playerID,
			Name:           "T20",
			Matches:        battingStats.T20.Matches,
			Innings:        battingStats.T20.Innings,
			Runs:           battingStats.T20.Runs,
			Balls:          battingStats.T20.Balls,
			Highest:        battingStats.T20.Highest,
			Average:        battingStats.T20.Average,
			SR:             battingStats.T20.SR,
			NotOut:         battingStats.T20.NotOut,
			Fours:          battingStats.T20.Fours,
			Sixes:          battingStats.T20.Sixes,
			Ducks:          battingStats.T20.Ducks,
			Fiftys:         battingStats.T20.Fiftys,
			Hundreds:       battingStats.T20.Hundreds,
			Two_hundreds:   battingStats.T20.Two_hundreds,
			Three_hunderds: battingStats.T20.Three_hunderds,
			Four_hundreds:  battingStats.T20.Four_hundreds,
		},
		{
			PlayerID:       playerID,
			Name:           "IPL",
			Matches:        battingStats.IPL.Matches,
			Innings:        battingStats.IPL.Innings,
			Runs:           battingStats.IPL.Runs,
			Balls:          battingStats.IPL.Balls,
			Highest:        battingStats.IPL.Highest,
			Average:        battingStats.IPL.Average,
			SR:             battingStats.IPL.SR,
			NotOut:         battingStats.IPL.NotOut,
			Fours:          battingStats.IPL.Fours,
			Sixes:          battingStats.IPL.Sixes,
			Ducks:          battingStats.IPL.Ducks,
			Fiftys:         battingStats.IPL.Fiftys,
			Hundreds:       battingStats.IPL.Hundreds,
			Two_hundreds:   battingStats.IPL.Two_hundreds,
			Three_hunderds: battingStats.IPL.Three_hunderds,
			Four_hundreds:  battingStats.IPL.Four_hundreds,
		},
	}

	// err = supabase.DB.From("Batting").Insert(rowTest).Execute(nil)

	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	// err = supabase.DB.From("Batting").Insert(rowODI).Execute(nil)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// err = supabase.DB.From("Batting").Insert(rowT20).Execute(nil)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	// err = supabase.DB.From("Batting").Insert(rowIPL).Execute(nil)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	//Fetching ID And inserting them to GameType table
	for _, row := range battingRows {
		err = supabase.DB.From("Batting").Insert(row).Execute(nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	fetchBatID, err := fetchBattingIDs()
	if err != nil {
		fmt.Println("Unable to fetch the Batsman's ID")
	}

	testbatID, odibatID, t20batID, iplbatID := distributeIDs(fetchBatID)

	row := BattingID{
		TestbatID: testbatID,
		ODIbatID:  odibatID,
		T20batID:  t20batID,
		IPLbatID:  iplbatID,
	}
	var results []BattingID
	err1 := supabase.DB.From("GameType").Insert(row).Execute(&results)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err1.Error()})
		return
	}

	bowlingStatsJSON, err := fetchStats(playerID, "bowling")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("Raw Bowling Stats JSON:", string(bowlingStatsJSON))

	// var bowlingStats BowlingStats
	// if err := json.Unmarshal(bowlingStatsJSON, &bowlingStats); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }
	bowlingStats, err := transformBowlingStats(bowlingStatsJSON)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println("bowlingStats struct:", bowlingStats)
	//fmt.Printf("Populated Bowling Stats: %+v\n", bowlingStats)
	bowlingRows := []FormatStatsBowling{
		{
			PlayerID:           playerID,
			Name:               "Test",
			Matches:            bowlingStats.Test.Matches,
			Innings:            bowlingStats.Test.Innings,
			Balls:              bowlingStats.Test.Balls,
			Runs:               bowlingStats.Test.Runs,
			Maidens:            bowlingStats.Test.Maidens,
			Wickets:            bowlingStats.Test.Wickets,
			Average:            bowlingStats.Test.Average,
			Economy:            bowlingStats.Test.Economy,
			StrikeRate:         bowlingStats.Test.StrikeRate,
			BestBowlingInnings: bowlingStats.Test.BestBowlingInnings,
			BestBowlingMatch:   bowlingStats.Test.BestBowlingMatch,
			FourWickets:        bowlingStats.Test.FourWickets,
			FiveWickets:        bowlingStats.Test.FiveWickets,
			TenWickets:         bowlingStats.Test.TenWickets,
		},
		{
			PlayerID:           playerID,
			Name:               "ODI",
			Matches:            bowlingStats.ODI.Matches,
			Innings:            bowlingStats.ODI.Innings,
			Balls:              bowlingStats.ODI.Balls,
			Runs:               bowlingStats.ODI.Runs,
			Maidens:            bowlingStats.ODI.Maidens,
			Wickets:            bowlingStats.ODI.Wickets,
			Average:            bowlingStats.ODI.Average,
			Economy:            bowlingStats.ODI.Economy,
			StrikeRate:         bowlingStats.ODI.StrikeRate,
			BestBowlingInnings: bowlingStats.ODI.BestBowlingInnings,
			BestBowlingMatch:   bowlingStats.ODI.BestBowlingMatch,
			FourWickets:        bowlingStats.ODI.FourWickets,
			FiveWickets:        bowlingStats.ODI.FiveWickets,
			TenWickets:         bowlingStats.ODI.TenWickets,
		},
		{
			PlayerID:           playerID,
			Name:               "T20",
			Matches:            bowlingStats.T20.Matches,
			Innings:            bowlingStats.T20.Innings,
			Balls:              bowlingStats.T20.Balls,
			Runs:               bowlingStats.T20.Runs,
			Maidens:            bowlingStats.T20.Maidens,
			Wickets:            bowlingStats.T20.Wickets,
			Average:            bowlingStats.T20.Average,
			Economy:            bowlingStats.T20.Economy,
			StrikeRate:         bowlingStats.T20.StrikeRate,
			BestBowlingInnings: bowlingStats.T20.BestBowlingInnings,
			BestBowlingMatch:   bowlingStats.T20.BestBowlingMatch,
			FourWickets:        bowlingStats.T20.FourWickets,
			FiveWickets:        bowlingStats.T20.FiveWickets,
			TenWickets:         bowlingStats.T20.TenWickets,
		},
		{
			PlayerID:           playerID,
			Name:               "IPL",
			Matches:            bowlingStats.IPL.Matches,
			Innings:            bowlingStats.IPL.Innings,
			Balls:              bowlingStats.IPL.Balls,
			Runs:               bowlingStats.IPL.Runs,
			Maidens:            bowlingStats.IPL.Maidens,
			Wickets:            bowlingStats.IPL.Wickets,
			Average:            bowlingStats.IPL.Average,
			Economy:            bowlingStats.IPL.Economy,
			StrikeRate:         bowlingStats.IPL.StrikeRate,
			BestBowlingInnings: bowlingStats.IPL.BestBowlingInnings,
			BestBowlingMatch:   bowlingStats.IPL.BestBowlingMatch,
			FourWickets:        bowlingStats.IPL.FourWickets,
			FiveWickets:        bowlingStats.IPL.FiveWickets,
			TenWickets:         bowlingStats.IPL.TenWickets,
		},
	}

	for _, row := range bowlingRows {
		err = supabase.DB.From("Bowling").Insert(row).Execute(nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

}

func transformBattingStats(rawJSON []byte) (Batting, error) {
	var inputJSON InputJSON
	if err := json.Unmarshal(rawJSON, &inputJSON); err != nil {
		return Batting{}, fmt.Errorf("error unmarshalling raw JSON: %v", err)
	}

	battingStats := Batting{}
	formatIndexMap := make(map[string]int)         // Map to hold the index for each format
	for i, header := range inputJSON.Headers[1:] { // Populate the map with format indices
		formatIndexMap[header] = i + 1 // +1 because Values[0] is the stat name, not a format value
	}

	for _, statValue := range inputJSON.Values {
		statName := statValue.Values[0] // e.g., "Matches"
		for format, index := range formatIndexMap {
			value := statValue.Values[index]
			switch format {
			case "Test":
				assignStatBatting(&battingStats.Test, statName, value)
			case "ODI":
				assignStatBatting(&battingStats.ODI, statName, value)
			case "T20":
				assignStatBatting(&battingStats.T20, statName, value)
			case "IPL":
				assignStatBatting(&battingStats.IPL, statName, value)
			}
		}
	}

	return battingStats, nil
}

func assignStatBatting(formatStats *FormatStatsBatting, statName, value string) {
	switch statName {
	case "Matches":
		formatStats.Matches = value
	case "Innings":
		formatStats.Innings = value
	case "Runs":
		formatStats.Runs = value
	case "Balls":
		formatStats.Balls = value
	case "Highest":
		formatStats.Highest = value
	case "Average":
		formatStats.Average = value
	case "SR":
		formatStats.SR = value
	case "Not Out":
		formatStats.NotOut = value
	case "Fours":
		formatStats.Fours = value
	case "Sixes":
		formatStats.Sixes = value
	case "Ducks":
		formatStats.Ducks = value
	case "50s":
		formatStats.Fiftys = value
	case "100s":
		formatStats.Hundreds = value
	case "200s":
		formatStats.Two_hundreds = value
	case "300s":
		formatStats.Three_hunderds = value
	case "400s":
		formatStats.Four_hundreds = value

	}
}

func transformBowlingStats(rawJSON []byte) (BowlingStats, error) {
	var inputJSON InputJSON
	if err := json.Unmarshal(rawJSON, &inputJSON); err != nil {
		return BowlingStats{}, fmt.Errorf("error unmarshalling raw JSON: %v", err)
	}

	bowlingStats := BowlingStats{}
	formatIndexMap := make(map[string]int)         // Map to hold the index for each format
	for i, header := range inputJSON.Headers[1:] { // Populate the map with format indices
		formatIndexMap[header] = i + 1 // +1 because Values[0] is the stat name, not a format value
	}

	for _, statValue := range inputJSON.Values {
		statName := statValue.Values[0] // e.g., "Matches"
		for format, index := range formatIndexMap {
			value := statValue.Values[index]
			switch format {
			case "Test":
				assignStatBowling(&bowlingStats.Test, statName, value)
			case "ODI":
				assignStatBowling(&bowlingStats.ODI, statName, value)
			case "T20":
				assignStatBowling(&bowlingStats.T20, statName, value)
			case "IPL":
				assignStatBowling(&bowlingStats.IPL, statName, value)
			}
		}
	}

	return bowlingStats, nil
}

func assignStatBowling(formatStats *FormatStatsBowling, statName, value string) {
	switch statName {
	case "Matches":
		formatStats.Matches = value
	case "Innings":
		formatStats.Innings = value
	case "Balls":
		formatStats.Balls = value
	case "Runs":
		formatStats.Runs = value
	case "Maidens":
		formatStats.Maidens = value
	case "Wickets":
		formatStats.Wickets = value
	case "Avg":
		formatStats.Average = value
	case "Eco":
		formatStats.Economy = value
	case "SR":
		formatStats.StrikeRate = value
	case "BBI":
		formatStats.BestBowlingInnings = value
	case "BBM":
		formatStats.BestBowlingMatch = value
	case "4w":
		formatStats.FourWickets = value
	case "5w":
		formatStats.FiveWickets = value
	case "10w":
		formatStats.TenWickets = value
		// Add other cases as needed
	}
}

func fetchStats(playerID string, statType string) ([]byte, error) {
	var url string
	apiKey := "35a1484e2dmsh245a0be5d118a20p1779d2jsn3c2d4bf768a0"

	// Determine which statistics to fetch based on statType
	if statType == "batting" {
		url = fmt.Sprintf("https://cricbuzz-cricket.p.rapidapi.com/stats/v1/player/%s/batting", playerID)
	} else if statType == "bowling" {
		url = fmt.Sprintf("https://cricbuzz-cricket.p.rapidapi.com/stats/v1/player/%s/bowling", playerID)
	} else {
		return nil, fmt.Errorf("invalid statType provided: %s", statType)
	}

	// Create a new request with a context and timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "cricbuzz-cricket.p.rapidapi.com")

	// Perform the request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error performing request: %v", err)
	}
	defer res.Body.Close()

	// Check the response status code
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP request failed with status: %s", res.Status)
	}

	// Read the response body
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	return body, nil
}

func PlayerName() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter player's first and last name: ")
	fullName, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading name:", err)
		return "", err
	}

	fullName = strings.TrimSpace(fullName)
	names := strings.Split(fullName, " ")

	var firsName, lastName string
	if len(names) > 0 {
		firsName = names[0]
	}
	if len(names) > 1 {
		lastName = strings.Join(names[1:], " ")
	}

	url := fmt.Sprintf("https://cricbuzz-cricket.p.rapidapi.com/stats/v1/player/search?plrN=%s%%20%s", firsName, lastName)
	Timoutctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(Timoutctx, "GET", url, nil)
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
	}
	apiKey := "35a1484e2dmsh245a0be5d118a20p1779d2jsn3c2d4bf768a0"
	req.Header.Add("X-RapidAPI-Key", apiKey)
	req.Header.Add("X-RapidAPI-Host", "cricbuzz-cricket.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("Error perfirming request: %v\n", err)
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Printf("HTTP Request Failed with Status: %s\n", res.Status)
		return "", fmt.Errorf("http request failed with status: %s", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return "", err
	}

	//Unmarshall json data for player
	jsonData := string(body)

	var Response struct {
		Players  []player.Playerinfo `json:"player"`
		Category string              `json:"category"`
	}

	if err := json.Unmarshal([]byte(jsonData), &Response); err != nil {
		fmt.Println("Error parsing JSON: ", err)
		return "", err
	}

	for _, p := range Response.Players {
		if p.Name == fullName {

			fmt.Printf("Player ID :%s, Name: %s\n", p.ID, p.Name)
			return p.ID, nil
		}
	}
	return "", fmt.Errorf("player not found")
}

func fetchBattingIDs() ([]int, error) {
	supabase := initializers.Supabase()
	var results []map[string]int
	err := supabase.DB.From("Batting").Select("id").Execute(&results)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch ids: %v", err)
	}

	var ids []int
	for _, result := range results {
		if id, ok := result["id"]; ok {
			ids = append(ids, id)
		}
	}

	return ids, nil
}

func distributeIDs(ids []int) (int, int, int, int) {
	var testID, odiID, t20ID, iplID int
	for i, id := range ids {
		switch i % 4 {
		case 0:
			testID = id
		case 1:
			odiID = id
		case 2:
			t20ID = id
		case 3:
			iplID = id
		}
	}
	// Further processing, such as updating other tables with these IDs, can be done here.
	return testID, odiID, t20ID, iplID
}
