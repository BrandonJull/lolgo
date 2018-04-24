package api

// Champions - Structure for a list of Champions retrieved using the API.
type Champions struct {
	Champion []Champion `json:"champions"`
}

// Champion - Structure for a Champion retrieved using the API.
type Champion struct {
	RankedPlayEnabled bool  `json:"rankedPlayEnabled"`
	BotEnabled        bool  `json:"botEnabled"`
	BotMMEnabled      bool  `json:"botMmEnabled"`
	Active            bool  `json:"active"`
	FreeToPlay        bool  `json:"freeToPlay"`
	ID                int64 `json:"id"`
}

// AllChampions fetches the list of all champions.
func AllChampions(api *RiotAPI) *Champions {
	url := "https://" + api.Region + ".api.riotgames.com/lol/platform/" + api.Version + "/champions" + "?api_key=" + api.APIKey

	champions := new(Champions)

	NewAPICall(url, &champions)

	return champions
}

// FreeToPlayChampions fetches a list of all current Free to Play champions.
func FreeToPlayChampions(api *RiotAPI) *Champions {
	url := "https://" + api.Region + ".api.riotgames.com/lol/platform/" + api.Version + "/champions" + "?freeToPlay=true&api_key=" + api.APIKey

	champions := new(Champions)

	NewAPICall(url, &champions)

	return champions
}

// NonFreeToPlayChampions fetches a list of all current non Free to Play champions.
func NonFreeToPlayChampions(api *RiotAPI) *Champions {
	url := "https://" + api.Region + ".api.riotgames.com/lol/platform/" + api.Version + "/champions" + "?freeToPlay=false&api_key=" + api.APIKey

	champions := new(Champions)

	NewAPICall(url, &champions)

	return champions
}

// ChampionByID fetches the champion by champion id.
func ChampionByID(api *RiotAPI, championID string) *Champion {
	url := "https://" + api.Region + ".api.riotgames.com/lol/platform/" + api.Version + "/champions/" + championID + "?api_key=" + api.APIKey

	champion := new(Champion)

	NewAPICall(url, &champion)

	return champion
}
