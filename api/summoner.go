package api

// Summoner - Structure for a Summoner retrieved using the API.
type Summoner struct {
	ProfileIconID int    `json:"profileIconId"`
	Name          string `json:"name"`
	SummonerLevel int64  `json:"summonerLevel"`
	RevisionDate  int64  `json:"revisionDate"`
	ID            int64  `json:"id"`
	AccountID     int64  `json:"accountId"`
}

// SummonerByName fetches the summoner by name.
func SummonerByName(api *RiotAPI, summonerName string) *Summoner {
	url := "https://" + api.Region + ".api.riotgames.com/lol/summoner/" + api.Version + "/summoners/by-name/" + summonerName + "?api_key=" + api.APIKey
	summoner := new(Summoner)

	NewAPICall(url, &summoner)

	return summoner
}

// SummonerByID fetches the summoner by summoner id.
func SummonerByID(api *RiotAPI, summonerID string) *Summoner {
	url := "https://" + api.Region + ".api.riotgames.com/lol/summoner/" + api.Version + "/summoners/" + summonerID + "?api_key=" + api.APIKey

	summoner := new(Summoner)

	NewAPICall(url, &summoner)

	return summoner
}

// SummonerByAccountID fetches the summoner by account id.
func SummonerByAccountID(api *RiotAPI, accountID string) *Summoner {
	url := "https://" + api.Region + ".api.riotgames.com/lol/summoner/" + api.Version + "/summoners/by-account/" + accountID + "?api_key=" + api.APIKey

	summoner := new(Summoner)

	NewAPICall(url, &summoner)

	return summoner
}
