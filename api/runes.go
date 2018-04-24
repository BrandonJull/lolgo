package api

// RunePages - Structure for the list of rune pages retrieved using the API.
type RunePages struct {
	Pages       []RunePage `json:"pages"`
	SummonderID int64      `json:"summonerId"`
}

// RunePage - Structure for a rune page retrieved using the API.
type RunePage struct {
	Current bool       `json:"current"`
	Slots   []RuneSlot `json:"slots"`
	Name    string     `json:"name"`
	ID      int64      `json:"id"`
}

// RuneSlot - Structure for a rune slot retrieved using the API.
type RuneSlot struct {
	RuneSlotID int `json:"runeSlotId"`
	RuneID     int `json:"runeId"`
}

// Runes fetches all rune page and slot data for a specific summoner.
func Runes(api *RiotAPI, summonerID string) *RunePages {
	url := "https://" + api.Region + ".api.riotgames.com/lol/platform/" + api.Version + "/runes/by-summoner/" + summonerID + "?api_key=" + api.APIKey

	runePages := new(RunePages)

	NewAPICall(url, runePages)

	return runePages
}
