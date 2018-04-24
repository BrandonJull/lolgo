package api

// MatchDTO - Structure to hold Match data retrieved using the API.
type MatchDTO struct {
	SeasonID              int                      `json:"seasonId"`
	QueueID               int                      `json:"queueId"`
	GameID                int64                    `json:"gameId"`
	ParticipantIdentities []ParticipantIdentityDTO `json:"participantIdentities"`
	GameVersion           string                   `json:"gameVersion"`
	PlatformID            string                   `json:"platformId"`
	GameMode              string                   `json:"gameMode"`
	MapID                 int                      `json:"mapId"`
	GameType              string                   `json:"gameType"`
	Teams                 []TeamStatsDTO           `json:"teams"`
	Participants          []ParticipantDTO         `json:"participants"`
	GameDuration          int64                    `json:"gameDuration"`
	GameCreation          int64                    `json:"gameCreation"`
}

// ParticipantIdentityDTO - Structure for holding ParticipantIdentity data retrieved using the API.
type ParticipantIdentityDTO struct {
	Player        PlayerDTO `json:"player"`
	ParticipantID int       `json:"participantId"`
}

// PlayerDTO -  Structure for holding Player data retrieved using the API.
type PlayerDTO struct {
	CurrentPlatformID string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	PlatformID        string `json:"platformId"`
	CurrentAccountID  int64  `json:"currentAccountId"`
	ProfileIcon       int    `json:"profileIcon"`
	SummonerID        int64  `json:"summonerId"`
	AccountID         int64  `json:"accountId"`
}

// TeamStatsDTO - Structure for holding TeamStats data retrieved using the API.
type TeamStatsDTO struct {
	FirstDragon          bool          `json:"firstDragon"`
	FirstInhibitor       bool          `json:"firstInhibitor"`
	Bans                 []TeamBansDTO `json:"bans"`
	BaronKills           int           `json:"baronKills"`
	FirstRiftHerald      bool          `json:"firstRiftHerald"`
	FirstBaron           bool          `json:"firstBaron"`
	RiftHeraldKills      int           `json:"riftHeraldKills"`
	FirstBlood           bool          `json:"firstBlood"`
	TeamID               int           `json:"teamId"`
	FirstTower           bool          `json:"firstTower"`
	VilemawKills         int           `json:"vilemawKills"`
	InhibitorKills       int           `json:"inhibitorKills"`
	TowerKills           int           `json:"towerKills"`
	DominionVictoryScore int           `json:"dominionVictoryScore"`
	Win                  string        `json:"win"`
	DragonKills          int           `json:"dragonKills"`
}

// TeamBansDTO - Structure for holding TeamBans data retrieved using the API.
type TeamBansDTO struct {
	PickTurn   int `json:"pickTurn"`
	ChampionID int `json:"championId"`
}

// ParticipantDTO - Structure for holding Participant data retrieved using the API.
type ParticipantDTO struct {
	Stats                     ParticipantStatsDTO    `json:"stats"`
	ParticipantID             int                    `json:"participantId"`
	Runes                     []MatchRuneDTO         `json:"runes"`
	Timeline                  ParticipantTimelineDTO `json:"timeline"`
	TeamID                    int                    `json:"teamid"`
	Spell2ID                  int                    `json:"spell2Id"`
	Masteries                 []MatchMasteryDTO      `json:"masteries"`
	HighestAchievedSeasonTier string                 `json:"highestAchievedSeasonTier"`
	Spell1ID                  int                    `json:"spell1Id"`
	ChampionID                int                    `json:"championId"`
}

// ParticipantStatsDTO - Structure for holding ParticipantStats data retrieved using the API.
type ParticipantStatsDTO struct {
	PhysicalDamageDealt             int64 `json:"physicalDamageDealt"`
	NeutralMinionsKilledTeamJungle  int   `json:"neutralMinionsKilledTeamJungle"`
	MagicDamageDealt                int64 `json:"magicDamageDealt"`
	TotalPlayerScore                int   `json:"totalPlayerScore"`
	Deaths                          int   `json:"deaths"`
	Win                             bool  `json:"win"`
	NeutralMinionsKilledEnemyJungle int   `json:"neutralMinionsKilledEnemyJungle"`
	AltarsCaptured                  int   `json:"altarsCaptured"`
	LargestCriticalStrike           int   `json:"largestCriticalStrike"`
	TotalDamageDealt                int64 `json:"totalDamageDealt"`
	MagicDamageDealtToChampions     int64 `json:"magicDamageDealtToChampions"`
	VisionWardsBoughtInGame         int   `json:"visionWardsBoughtInGame"`
	DamageDealtToObjectives         int64 `json:"damageDealtToObjectives"`
	LargestKillingSpree             int   `json:"largestKillingSpree"`
	Item1                           int   `json:"item1"`
	QuadraKills                     int   `json:"quadraKills"`
	TeamObjective                   int   `json:"teamObjective"`
	TotalTimeCrowdControlDealt      int   `json:"totalTimeCrowdControlDealt"`
	LongestTimeSpentLiving          int   `json:"longestTimeSpentLiving"`
	WardsKilled                     int   `json:"wardsKilled"`
	FirstTowerAsset                 bool  `json:"firstTowerAssist"`
	FirstTowerKill                  bool  `json:"firstTowerKill"`
	Item2                           int   `json:"item2"`
	Item3                           int   `json:"item3"`
	Item0                           int   `json:"item0"`
	FirstBloodAssist                bool  `json:"firstBloodAssist"`
	VisionScore                     int64 `json:"visionScore"`
	WardsPlaced                     int   `json:"wardsPlaced"`
	Item4                           int   `json:"item4"`
	Item5                           int   `json:"item5"`
	Item6                           int   `json:"item6"`
	TurretKills                     int   `json:"turretKills"`
	TripleKills                     int   `json:"tripleKills"`
	DamageSelfMitigated             int64 `json:"damageSelfMitigated"`
	ChampLevel                      int   `json:"champLevel"`
	NodeNeutralizeAssist            int   `json:"nodeNeutralizeAssist"`
	FirstInhibitorKill              bool  `json:"firstInhibitorKill"`
	GoldEarned                      int   `json:"goldEarned"`
	MagicalDamageTaken              int64 `json:"magicalDamageTaken"`
	Kills                           int   `json:"kills"`
	DoubleKills                     int   `json:"doubleKills"`
	NodeCaptureAssist               int   `json:"nodeCaptureAssist"`
	TrueDamageTaken                 int64 `json:"trueDamageTaken"`
	NodeNeutralize                  int   `json:"nodeNeutralize"`
	FirstInhibitorAssist            bool  `json:"firstInhibitorAssist"`
	Assists                         int   `json:"assists"`
	UnrealKills                     int   `json:"unrealKills"`
	NeutralMinionsKilled            int   `json:"neutralMinionsKilled"`
	ObjectivePlayerScore            int   `json:"objectivePlayerScore"`
	CombatPlayerScore               int   `json:"combatPlayerScore"`
	DamageDealtToTurrets            int64 `json:"damageDealtToTurrets"`
	AltarsNeutralized               int   `json:"altarsNeutralized"`
	PhysicalDamageDealtToChampions  int64 `json:"physicalDamageDealtToChampions"`
	GoldSpent                       int   `json:"goldSpent"`
	TrueDamageDealt                 int64 `json:"trueDamageDealt"`
	TrueDamageDealtToChampions      int64 `json:"trueDamageDealtToChampions"`
	ParticipantID                   int   `json:"participantId"`
	PentaKills                      int   `json:"pentaKills"`
	TotalHeal                       int64 `json:"totalHeal"`
	TotalMinionsKilled              int   `json:"totalMinionsKilled"`
	FirstBloodKill                  bool  `json:"firstBloodKill"`
	NodeCapture                     int   `json:"nodeCapture"`
	LargestMultiKill                int   `json:"largestMultiKill"`
	SightWardsBoughtInGame          int   `json:"sightWardsBoughtInGame"`
	TotalDamageDealtToChampions     int64 `json:"totalDamageDealtToChampions"`
	TotalUnitsHealed                int   `json:"totalUnitsHealed"`
	InhibitorKills                  int   `json:"inhibitorKills"`
	TotalScoreRank                  int   `json:"totalScoreRank"`
	TotalDamageTaken                int64 `json:"totalDamageTaken"`
	KillingSprees                   int   `json:"killingSprees"`
	TimeCCingOthers                 int64 `json:"timeCCingOthers"`
	PhysicalDamageTaken             int64 `json:"physicalDamageTaken"`
}

// MatchRuneDTO - Structure for holding MatchRune data retrieved using the API.
type MatchRuneDTO struct {
	RuneID int `json:"runeID"`
	Rank   int `json:"rank"`
}

// ParticipantTimelineDTO - Structure for holding ParticipantTimeline data retrieved using the API.
type ParticipantTimelineDTO struct {
	Lane                        string             `json:"lane"`
	ParticipantID               int                `json:"participantId"`
	CSDiffPerMinDeltas          map[string]float64 `json:"csDiffPerMinDeltas"`
	GoldPerMinDeltas            map[string]float64 `json:"goldPerMinDeltas"`
	XPDiffPerMinDeltas          map[string]float64 `json:"xpDiffPerMinDeltas"`
	CreepsPerMinDeltas          map[string]float64 `json:"creepsPerMinDeltas"`
	XPPerMinDeltas              map[string]float64 `json:"xpPerMinDeltas"`
	Role                        string             `json:"role"`
	DamageTakenDiffPerMinDeltas map[string]float64 `json:"damageTakenDiffPerMinDeltas"`
	DamageTakenPerMinDeltas     map[string]float64 `json:"damageTakenPerMinDeltas"`
}

// MatchMasteryDTO - Structure for holding MatchMastery data retrieved using the API.
type MatchMasteryDTO struct {
	MasteryID int `json:"masteryId"`
	Rank      int `json:"rank"`
}

// MatchlistDTO - Structure for holding Matchlist data retrieved using the API.
type MatchlistDTO struct {
	Matches    []MatchReferenceDTO `json:"matches"`
	TotalGames int                 `json:"totalGames"`
	StartIndex int                 `json:"startIndex"`
	EndIndex   int                 `json:"endIndex"`
}

// MatchReferenceDTO - Structure for holding MatchReference data retrieved using the API.
type MatchReferenceDTO struct {
	Lane       string `json:"lane"`
	GameID     int64  `json:"gameId"`
	Champion   int    `json:"champion"`
	PlatformID string `json:"platformId"`
	Season     int    `json:"season"`
	Queue      int    `json:"queue"`
	Role       string `json:"role"`
	Timestamp  int64  `json:"timestamp"`
}

// MatchTimelineDTO - Structure for holding MatchTimeline data retrieved using the API.
type MatchTimelineDTO struct {
	Frames        []MatchFrameDTO `json:"frames"`
	FrameInterval int64           `json:"frameInterval"`
}

// MatchFrameDTO - Structure for holding MatchFrame data retrieved using the API.
type MatchFrameDTO struct {
	Timestamp         int64                            `json:"timestamp"`
	ParticipantFrames map[int]MatchParticipantFrameDTO `json:"participantFrames"`
	Events            []MatchEventDTO                  `json:"events"`
}

// MatchParticipantFrameDTO - Structure for holding MatchParticipantFrame data retrieved using the API.
type MatchParticipantFrameDTO struct {
	TotalGold           int              `json:"totalGold"`
	TeamScore           int              `json:"teamScore"`
	ParticipantID       int              `json:"participantId"`
	Level               int              `json:"level"`
	CurrentGold         int              `json:"currentGold"`
	MinionsKilled       int              `json:"minionsKilled"`
	DominionScore       int              `json:"dominionScore"`
	Position            MatchPositionDTO `json:"position"`
	XP                  int              `json:"xp"`
	JungleMinionsKilled int              `json:"jungleMinionsKilled"`
}

// MatchPositionDTO - Structure for holding MatchPosition data retrieved using the API.
type MatchPositionDTO struct {
	Y int `json:"y"`
	X int `json:"x"`
}

// MatchEventDTO - Structure for holding MatchEvent data retrieved using the API.
type MatchEventDTO struct {
	EventType               string           `json:"eventType"`
	TowerType               string           `json:"towerType"`
	TeamID                  int              `json:"teamId"`
	AscendedType            string           `json:"ascendedType"`
	KillerID                int              `json:"killerId"`
	LevelUpType             string           `json:"levelUpType"`
	PointCaptured           string           `json:"pointCaptured"`
	AssistingParticipantIDs []int            `json:"assistingParticipantIds"`
	WardType                string           `json:"wardType"`
	MonsterType             string           `json:"monsterType"`
	Type                    string           `json:"type"`
	SkillSlot               int              `json:"skillSlot"`
	VictimID                int              `json:"victimId"`
	Timestamp               int64            `json:"timestamp"`
	AfterID                 int              `json:"afterId"`
	MonsterSubType          string           `json:"monsterSubType"`
	LaneType                string           `json:"laneType"`
	ItemID                  int              `json:"itemId"`
	ParticipantID           int              `json:"participantId"`
	BuildingType            string           `json:"buildingType"`
	CreatorID               int              `json:"creatorId"`
	Position                MatchPositionDTO `json:"position"`
	BeforeID                int              `json:"beforeId"`
}

// Match fetches all match data for a given match.
func Match(api *RiotAPI, matchID string) *MatchDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/match/" + api.Version + "/matches/" + matchID + "?api_key=" + api.APIKey

	match := new(MatchDTO)

	NewAPICall(url, match)

	return match
}

// MatchlistByAccountID fetches a list of match data for a given account.
func MatchlistByAccountID(api *RiotAPI, accountID string) *MatchlistDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/match/" + api.Version + "/matchlists/by-account/" + accountID + "?api_key=" + api.APIKey

	matchlist := new(MatchlistDTO)

	NewAPICall(url, matchlist)

	return matchlist
}

// RecentMatchlistByAccountID fetches a list of recent match data for a given account.
func RecentMatchlistByAccountID(api *RiotAPI, accountID string) *MatchlistDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/match/" + api.Version + "/matchlists/by-account/" + accountID + "/recent?api_key=" + api.APIKey

	matchlist := new(MatchlistDTO)

	NewAPICall(url, matchlist)

	return matchlist
}

// MatchTimeline fetches the timeline data for a specific match.
func MatchTimeline(api *RiotAPI, matchID string) *MatchTimelineDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/match/" + api.Version + "/timelines/by-match/" + matchID + "?api_key=" + api.APIKey

	matchTimeline := new(MatchTimelineDTO)

	NewAPICall(url, matchTimeline)

	return matchTimeline
}
