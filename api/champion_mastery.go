package api

// ChampionMasteriesRequest - JSON Array of champion masteries retrieved using the API.
type ChampionMasteriesRequest []ChampionMastery

// ChampionMasteries - Structure JSON array contents of ChampionMasteriesRequest
type ChampionMasteries struct {
	ChampionMastery ChampionMasteriesRequest
}

// MasteryScore - Holds mastery score retrieved using the API.
type MasteryScore int

// ChampionMastery - Structure for a ChampionMastery retrieved using the API.
type ChampionMastery struct {
	ChestGranted                 bool  `json:"chestGranted"`
	ChampionLevel                int   `json:"championLevel"`
	ChampionPoints               int   `json:"championPoints"`
	ChampionID                   int64 `json:"championID"`
	PlayerID                     int64 `json:"playerID"`
	ChampionPointsUntilNextLevel int64 `json:"championPointsUntilNextLevel"`
	ChampionPointsSinceLastLevel int64 `json:"championPointsSinceLastLevel"`
	LastPlayTime                 int64 `json:"LastPlayTime"`
}

// AllChampionMasteries fetches the all the champion masteries by summoner id.
func AllChampionMasteries(api *RiotAPI, summonerID string) *ChampionMasteries {
	url := "https://" + api.Region + ".api.riotgames.com/lol/champion-mastery/" + api.Version + "/champion-masteries/by-summoner/" + summonerID + "?api_key=" + api.APIKey

	championMasteriesRequest := new(ChampionMasteriesRequest)

	NewAPICall(url, &championMasteriesRequest)

	championMasteries := new(ChampionMasteries)
	championMasteries.ChampionMastery = *championMasteriesRequest

	return championMasteries
}

// ChampionMasteryByID fetches the champion mastery by summoner and champion id.
func ChampionMasteryByID(api *RiotAPI, summonerID string, championID string) *ChampionMastery {
	url := "https://" + api.Region + ".api.riotgames.com/lol/champion-mastery/" + api.Version + "/champion-masteries/by-summoner/" + summonerID + "/by-champion/" + championID + "?api_key=" + api.APIKey

	championMastery := new(ChampionMastery)

	NewAPICall(url, &championMastery)

	return championMastery
}

// ChampionMasteryScore fetches the mastery score of a summoner given the summoner id.
func ChampionMasteryScore(api *RiotAPI, summonerID string) MasteryScore {
	url := "https://" + api.Region + ".api.riotgames.com/lol/champion-mastery/" + api.Version + "/scores/by-summoner/" + summonerID + "?api_key=" + api.APIKey

	masteryScore := new(MasteryScore)

	NewAPICall(url, masteryScore)

	return *masteryScore
}
