package player

import "time"

type Player struct {
	IPL  IPL
	T20  T20
	ODI  ODI
	TEST Test
}

type Playerinfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	TeamName    string `json:"teamName"`
	FaceImageID string `json:"faceImageId"`
	Age         string
	DateOfBirth time.Time
}

type Batting struct {
	MatchesPlayed string
	BattingAvg    string
	HighestScore  string
	StrikeRate    string
	Centuries     string
	HalfCenturies string
	Fours         string
	Sixes         string
}
type Bowling struct {
	MatchesPlayed  string
	InningsBowled  string
	Wickets        string
	BowlingAverage string
	EconomyRate    string
	StrikeRate     string
	BestBowling    string
	FiveWicket     string
	Maiden         string
}
type IPL struct {
	IPL string
	Batting
	Bowling
}
type T20 struct {
	T20 string
	Batting
	Bowling
}

type ODI struct {
	ODI string
	Batting
	Bowling
}

type Test struct {
	Test string
	Batting
	Bowling
}
