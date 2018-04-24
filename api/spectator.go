package api

// FeaturedGames - Structure for FeaturedGames retrieved using the API.
type FeaturedGames struct {
	ClientRefreshInterval int64      `json:"clientRefreshInterval"`
	GameList              []GameInfo `json:"gameList"`
}

// GameInfo - Structure for GameInfo retrieved using the API.
type GameInfo struct {
	GameID            int64                    `json:"gameId"`
	GameStartTime     int64                    `json:"gameStartTime"`
	PlatformID        string                   `json:"platformId"`
	GameMode          string                   `json:"gameMode"`
	MapID             int64                    `json:"mapId"`
	GameType          string                   `json:"gameType"`
	BannedChampions   []BannedChampion         `json:"bannedChampions"`
	Observers         Observer                 `json:"observers"`
	Participants      []CurrentGameParticipant `json:"participants"`
	GameLength        int64                    `json:"gameLength"`
	GameQueueConfigID int64                    `json:"gameQueueConfigId"`
}

// BannedChampion - Structure for a BannedChampion retrieved using the API.
type BannedChampion struct {
	PickTurn   int   `json:"pickTurn"`
	ChampionID int64 `json:"championId"`
	TeamID     int64 `json:"teamId"`
}

// Observer - Structure for an Observer retrieved using the API.
type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}

// CurrentGameParticipant - Structure for a CurrentGameParticipant retrieved using the API.
type CurrentGameParticipant struct {
	ProfileIconID int64     `json:"profileIconId"`
	ChampionID    int64     `json:"championId"`
	SummonerName  string    `json:"summonerName"`
	Runes         []Rune    `json:"runes"`
	Bot           bool      `json:"bot"`
	TeamID        int64     `json:"teamId"`
	Spell2ID      int64     `json:"spell2Id"`
	Masteries     []Mastery `json:"masteries"`
	Spell1ID      int64     `json:"spell1Id"`
	SummonerID    int64     `json:"summonerId"`
}

// Rune - Structure for a Rune retrieved using the API.
type Rune struct {
	Count  int   `json:"count"`
	RuneID int64 `json:"runeId"`
}

// Mastery - Structure for a Mastery retrieved using the API.
type Mastery struct {
	MasteryID int64 `json:"masteryId"`
	Rank      int   `json:"rank"`
}

// ActiveGame fetches game spectator mode information of an active game given the summoner id.
func ActiveGame(api *RiotAPI, summonerID string) *GameInfo {
	url := "https://" + api.Region + ".api.riotgames.com/lol/spectator/" + api.Version + "/active-games/by-summoner/" + summonerID + "?api_key=" + api.APIKey

	gameInfo := new(GameInfo)

	NewAPICall(url, &gameInfo)

	return gameInfo
}

// AllFeaturedGames fetches all featured games open for spectacting.
func AllFeaturedGames(api *RiotAPI) *FeaturedGames {
	url := "https://" + api.Region + ".api.riotgames.com/lol/spectator/" + api.Version + "/featured-games?api_key=" + api.APIKey

	featuredGames := new(FeaturedGames)

	NewAPICall(url, &featuredGames)

	return featuredGames
}
