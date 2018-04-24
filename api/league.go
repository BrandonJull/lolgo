package api

import "fmt"

// LeagueListDTO - Structure for a list consisting of static data for leagues retreived using the API.
type LeagueListDTO struct {
	LeagueID string          `json:"leagueId"`
	Tier     string          `json:"tier"`
	Entries  []LeagueItemDTO `json:"entries"`
	Queue    string          `json:"queue"`
	Name     string          `json:"name"`
}

// LeagueItemDTO - Structure for league item static data retrieved using the API.
type LeagueItemDTO struct {
	Rank             string        `json:"rank"`
	HotStreak        bool          `json:"hotStreak"`
	MiniSeries       MiniSeriesDTO `json:"miniSeries"`
	Wins             int           `json:"wins"`
	Veteran          bool          `json:"veteran"`
	Losses           int           `json:"losses"`
	FreshBlood       bool          `json:"freshBlood"`
	PlayerOrTeamName string        `json:"playerOrTeamName"`
	Inactive         bool          `json:"inactive"`
	PlayerOrTeamID   string        `json:"playerOrTeamId"`
	LeaguePoints     int           `json:"leaguePoints"`
}

// MiniSeriesDTO - Structure for league mini series static data retrieved using the API.
type MiniSeriesDTO struct {
	Wins     int    `json:"wins"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Progress string `json:"progress"`
}

// LeaguePositionRequest - JSON Array of league positions for a specific summoner.
type LeaguePositionRequest []LeaguePositionDTO

// LeaguePositions - Structure JSON Array contents of LeaguePositionRequest.
type LeaguePositions struct {
	LeaguePosition LeaguePositionRequest
}

// LeaguePositionDTO - Structure for league position static data retrieved using the API.
type LeaguePositionDTO struct {
	Rank             string        `json:"rank"`
	HotStreak        bool          `json:"hotStreak"`
	MiniSeries       MiniSeriesDTO `json:"miniSeries"`
	Wins             int           `json:"wins"`
	Veteran          bool          `json:"veteran"`
	Losses           int           `json:"losses"`
	FreshBlood       bool          `json:"freshBlood"`
	PlayerOrTeamName string        `json:"playerOrTeamName"`
	Inactive         bool          `json:"inactive"`
	PlayerOrTeamID   string        `json:"playerOrTeamId"`
	LeagueID         string        `json:"leagueId"`
	LeaguePoints     int           `json:"leaguePoints"`
}

// ChallengerLeaguesByQueue fetches all league data for a specific queue.
func ChallengerLeaguesByQueue(api *RiotAPI, queue string) *LeagueListDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/league/" + api.Version + "/challengerleagues/by-queue/" + queue + "?api_key=" + api.APIKey

	fmt.Println(url)

	leagueList := new(LeagueListDTO)

	NewAPICall(url, &leagueList)

	return leagueList
}

// LeagueByID fetches all league data given the id.
func LeagueByID(api *RiotAPI, leagueID string) *LeagueListDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/league/" + api.Version + "/leagues/" + leagueID + "?api_key=" + api.APIKey

	leagueList := new(LeagueListDTO)

	NewAPICall(url, &leagueList)

	return leagueList
}

// MasterLeaguesByQueue fetches all league data for a specific queue.
func MasterLeaguesByQueue(api *RiotAPI, queue string) *LeagueListDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/league/" + api.Version + "/masterleagues/by-queue/" + queue + "?api_key=" + api.APIKey

	leagueList := new(LeagueListDTO)

	NewAPICall(url, &leagueList)

	return leagueList
}

// LeaguePositionsBySummonerID fetches the league positions of a summoner given the id.
func LeaguePositionsBySummonerID(api *RiotAPI, summonerID string) *LeaguePositions {
	url := "https://" + api.Region + ".api.riotgames.com/lol/league/" + api.Version + "/positions/by-summoner/" + summonerID + "?api_key=" + api.APIKey

	leaguePositionRequest := new(LeaguePositionRequest)

	NewAPICall(url, &leaguePositionRequest)

	leaguePositions := new(LeaguePositions)
	leaguePositions.LeaguePosition = *leaguePositionRequest

	return leaguePositions
}
