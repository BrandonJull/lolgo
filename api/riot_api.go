package api

// RiotAPI - API connection structure.
type RiotAPI struct {
	APIKey  string
	Region  string
	Version string
}

// NewRiotAPI constructs the connection to the API given
// the APIKey in string format.
func NewRiotAPI(apiKey string, region string, version string) *RiotAPI {
	return &RiotAPI{apiKey, region, version}
}

// GetSummonerByName fetches a Summoner given the SummonerName
// in string format.
func (api *RiotAPI) GetSummonerByName(summonerName string) *Summoner {
	return SummonerByName(api, summonerName)
}

// GetSummonerByID fetches a Summoner given the SummonerName
// in string format.
func (api *RiotAPI) GetSummonerByID(summonerID string) *Summoner {
	return SummonerByID(api, summonerID)
}

// GetSummonerByAccountID fetches a Summoner given the SummonerName
// in string format.
func (api *RiotAPI) GetSummonerByAccountID(summonerAccountID string) *Summoner {
	return SummonerByAccountID(api, summonerAccountID)
}

// GetChampionByID fetches a Champion given the ChampionID
// in string format.
func (api *RiotAPI) GetChampionByID(championID string) *Champion {
	return ChampionByID(api, championID)
}

// GetChampions fetches a list of all Champions.
func (api *RiotAPI) GetChampions() *Champions {
	return AllChampions(api)
}

// GetFreeToPlayChampions fetches a list of all Free to Play Champions.
func (api *RiotAPI) GetFreeToPlayChampions() *Champions {
	return FreeToPlayChampions(api)
}

// GetNonFreeToPlayChampions fetches a list of all non Free to Play Champions.
func (api *RiotAPI) GetNonFreeToPlayChampions() *Champions {
	return NonFreeToPlayChampions(api)
}

// GetChampionMasteryByID fetches the champion mastery of a specific champion given the summoner's id and champion id.
func (api *RiotAPI) GetChampionMasteryByID(summonerID string, championID string) *ChampionMastery {
	return ChampionMasteryByID(api, summonerID, championID)
}

// GetChampionMasteries fetches a list of all champion masteries for a specifc Summoner.
func (api *RiotAPI) GetChampionMasteries(summonerID string) *ChampionMasteries {
	return AllChampionMasteries(api, summonerID)
}

// GetChampionMasteryScore fetches the champion mastery score of a specific Summoner.
func (api *RiotAPI) GetChampionMasteryScore(summonerID string) MasteryScore {
	return ChampionMasteryScore(api, summonerID)
}

// GetActiveGame fetches current active game spectator mode information for a specific Summoner.
func (api *RiotAPI) GetActiveGame(summonerID string) *GameInfo {
	return ActiveGame(api, summonerID)
}

// GetFeaturedGames fetches all current featured games open for spectating.
func (api *RiotAPI) GetFeaturedGames() *FeaturedGames {
	return AllFeaturedGames(api)
}

// GetRunes fetches all rune page and slot data for a specific summoner.
func (api *RiotAPI) GetRunes(SummonerID string) *RunePages {
	return Runes(api, SummonerID)
}

// GetChampionStaticData fetches all champion static data for every champion.
func (api *RiotAPI) GetChampionStaticData() *ChampionList {
	return ChampionStaticData(api)
}

// GetChampionStaticDataByID fetches all champion static data for a specific champion.
func (api *RiotAPI) GetChampionStaticDataByID(championID string) *ChampionDTO {
	return ChampionStaticDataByID(api, championID)
}

// GetItemStaticData fetches all item static data for every item.
func (api *RiotAPI) GetItemStaticData() *ItemListDTO {
	return ItemStaticData(api)
}

// GetItemStaticDataByID fetches all item static data for a specific item.
func (api *RiotAPI) GetItemStaticDataByID(itemID string) *ItemDTO {
	return ItemStaticDataByID(api, itemID)
}

// GetLanguageStrings fetches all language string static data.
func (api *RiotAPI) GetLanguageStrings() *LanguageStringsDTO {
	return LanguageStrings(api)
}

// GetLanguages fetches the language list static data.
func (api *RiotAPI) GetLanguages() LanguageList {
	return Languages(api)
}

// GetMapStaticData fetches the map static data for all maps.
func (api *RiotAPI) GetMapStaticData() *MapDataDTO {
	return MapStaticData(api)
}

// GetMasteryStaticData fetches the static data for all masteries.
func (api *RiotAPI) GetMasteryStaticData() *MasteryListDTO {
	return MasteryStaticData(api)
}

// GetMasteryStaticDataByID fetches the static data for all masteries.
func (api *RiotAPI) GetMasteryStaticDataByID(masteryID string) *MasteryDTO {
	return MasteryStaticDataByID(api, masteryID)
}

// GetProfileIconStaticData fetches profile icon static data for all icons.
func (api *RiotAPI) GetProfileIconStaticData() *ProfileIconDataDTO {
	return ProfileIconStaticData(api)
}

// GetRealmStaticData fetches realm static data.
func (api *RiotAPI) GetRealmStaticData() *RealmDTO {
	return RealmStaticData(api)
}

// GetReforgedRunePathsStaticData fetches a list of all reforged rune paths.
func (api *RiotAPI) GetReforgedRunePathsStaticData() *ReforgedRunePaths {
	return ReforgedRunePathsStaticData(api)
}

// GetReforgedRunePathsStaticDataByID fetches a list of all reforged rune paths.
func (api *RiotAPI) GetReforgedRunePathsStaticDataByID(id string) *ReforgedRunePathDTO {
	return ReforgedRunePathsStaticDataByID(api, id)
}

// GetReforgedRuneStaticData fetches all reforged rune static data.
func (api *RiotAPI) GetReforgedRuneStaticData() *ReforgedRunes {
	return ReforgedRuneStaticData(api)
}

// GetReforgedRuneStaticDataByID fetches all reforged rune static data.
func (api *RiotAPI) GetReforgedRuneStaticDataByID(id string) *ReforgedRuneDTO {
	return ReforgedRuneStaticDataByID(api, id)
}

// GetRuneStaticData fetches the static data for all runes.
func (api *RiotAPI) GetRuneStaticData() *RuneListDTO {
	return RuneStaticData(api)
}

// GetRuneStaticDataByID fetches the static data for a specific rune.
func (api *RiotAPI) GetRuneStaticDataByID(runeID string) *RuneDTO {
	return RuneStaticDataByID(api, runeID)
}

// GetSummonerSpellData fetches the static data for all summoner spells.
func (api *RiotAPI) GetSummonerSpellData() *SummonerSpellListDTO {
	return SummonerSpellData(api)
}

// GetSummonerSpellDataByID fetches the static data for a specific summoner spell.
func (api *RiotAPI) GetSummonerSpellDataByID(summonerSpellID string) *SummonerSpellDTO {
	return SummonerSpellDataByID(api, summonerSpellID)
}

// GetVersions fetches the version list static data.
func (api *RiotAPI) GetVersions() VersionList {
	return Versions(api)
}

// GetChallengerLeaguesByQueue fetches all league data for a specific queue type.
func (api *RiotAPI) GetChallengerLeaguesByQueue(queue string) *LeagueListDTO {
	return ChallengerLeaguesByQueue(api, queue)
}

// GetLeagueByID fetches all league data given the league id.
func (api *RiotAPI) GetLeagueByID(leagueID string) *LeagueListDTO {
	return LeagueByID(api, leagueID)
}

// GetMasterLeaguesByQueue fetches all league data for a specific queue type.
func (api *RiotAPI) GetMasterLeaguesByQueue(queue string) *LeagueListDTO {
	return MasterLeaguesByQueue(api, queue)
}

// GetLeaguePositionsBySummonerID fetches summoner league positions static data for a specific summoner.
func (api *RiotAPI) GetLeaguePositionsBySummonerID(summonerID string) *LeaguePositions {
	return LeaguePositionsBySummonerID(api, summonerID)
}

// GetShardData fetches shard data for a specific server.
func (api *RiotAPI) GetShardData() *ShardStatus {
	return ShardData(api)
}

// GetMatch fetches match data for a specific match.
func (api *RiotAPI) GetMatch(matchID string) *MatchDTO {
	return Match(api, matchID)
}

// GetMatchlistByAccountID fetches a list of match data for a specific account ID.
func (api *RiotAPI) GetMatchlistByAccountID(accountID string) *MatchlistDTO {
	return MatchlistByAccountID(api, accountID)
}

// GetRecentMatchlistByAccountID fetches a list of match data for a specific account ID.
func (api *RiotAPI) GetRecentMatchlistByAccountID(accountID string) *MatchlistDTO {
	return RecentMatchlistByAccountID(api, accountID)
}

// GetMatchTimeline fetches timeline data for a specific match ID.
func (api *RiotAPI) GetMatchTimeline(matchID string) *MatchTimelineDTO {
	return MatchTimeline(api, matchID)
}
