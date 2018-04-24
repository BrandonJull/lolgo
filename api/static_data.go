package api

// ChampionList - Structure for a list consisting of static data for all champions retrieved using the API.
type ChampionList struct {
	Keys    map[string]string      `json:"keys"`
	Data    map[string]ChampionDTO `json:"data"`
	Version string                 `json:"version"`
	Type    string                 `json:"type"`
	Format  string                 `json:"format"`
}

// ChampionDTO - Structure for champion static data retrieved using the API.
type ChampionDTO struct {
	Info        InfoDTO            `json:"info"`
	EnemyTips   []string           `json:"enemytips"`
	Stats       StatsDTO           `json:"stats"`
	Name        string             `json:"name"`
	Title       string             `json:"title"`
	Image       ImageDTO           `json:"image"`
	Tags        []string           `json:"tags"`
	ParType     string             `json:"partype"`
	Skins       []SkinDTO          `json:"skins"`
	Passive     PassiveDTO         `json:"passive"`
	Recommended []RecommendedDTO   `json:"recommended"`
	AllyTips    []string           `json:"allytips"`
	Key         string             `json:"key"`
	Lore        string             `json:"lore"`
	ID          int                `json:"id"`
	Blurb       string             `json:"blurb"`
	Spells      []ChampionSpellDTO `json:"spells"`
}

// InfoDTO - Structure for champion info retrieved using the API.
type InfoDTO struct {
	Difficulty int `json:"difficulty"`
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
}

// StatsDTO - Structure for champion stats retrieved using the API.
type StatsDTO struct {
	ArmorPerLevel        float64 `json:"armorperlevel"`
	HPPerLevel           float64 `json:"hpperlevel"`
	AttackDamage         float64 `json:"attackdamage"`
	MPPerLevel           float64 `json:"mpperlevel"`
	AttackSpeedOffset    float64 `json:"attackspeedoffset"`
	Armor                float64 `json:"armor"`
	HP                   float64 `json:"hp"`
	HPRegenPerLevel      float64 `json:"hpregenperlevel"`
	SpellBlock           float64 `json:"spellblock"`
	AttackRange          float64 `json:"attackrange"`
	MoveSpeed            float64 `json:"movespeed"`
	AttackDamagePerLevel float64 `json:"attackdamageperlevel"`
	MPRegenPerLevel      float64 `json:"mpregenperlevel"`
	MP                   float64 `json:"mp"`
	SpellBlockPerLevel   float64 `json:"spellblockperlevel"`
	Crit                 float64 `json:"crit"`
	MPRegen              float64 `json:"mpregen"`
	AttackSpeedPerLevel  float64 `json:"attackspeedperlevel"`
	HPRegen              float64 `json:"hpregen"`
	CritPerLevel         float64 `json:"critperlevel"`
}

// ImageDTO - Structure for static data images retrieved using the API.
type ImageDTO struct {
	Full   string `json:"full"`
	Group  string `json:"group"`
	Sprite string `json:"sprite"`
	H      int    `json:"h"`
	W      int    `json:"w"`
	Y      int    `json:"y"`
	X      int    `json:"x"`
}

// SkinDTO - Structure for a champion skin retrieved using the API.
type SkinDTO struct {
	Num  int    `json:"num"`
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// PassiveDTO - Structure for a champion passive retrieved using the API.
type PassiveDTO struct {
	Image                ImageDTO `json:"image"`
	SanitizedDescription string   `json:"sanitizedDescription"`
	Name                 string   `json:"name"`
	Description          string   `json:"description"`
}

// RecommendedDTO - Structure for champion recommended data retrieved using the API.
type RecommendedDTO struct {
	Map      string     `json:"map"`
	Blocks   []BlockDTO `json:"blocks"`
	Champion string     `json:"champion"`
	Title    string     `json:"title"`
	Priority bool       `json:"priority"`
	Mode     string     `json:"mode"`
	Type     string     `json:"type"`
}

// BlockDTO - Structure for champion recommended block data retrieved using the API.
type BlockDTO struct {
	Items   []BlockItemDTO `json:"items"`
	RecMath bool           `json:"recMath"`
	Type    string         `json:"type"`
}

// BlockItemDTO - Structure for champion recommended block item data retrieved using the API.
type BlockItemDTO struct {
	Count int `json:"count"`
	ID    int `json:"int"`
}

// ChampionSpellDTO - Structure for champion spell data retrieved using the API.
type ChampionSpellDTO struct {
	CooldownBurn         string         `json:"cooldownBurn"`
	Resource             string         `json:"resource"`
	LevelTip             LevelTipDTO    `json:"leveltip"`
	Vars                 []SpellVarsDTO `json:"vars"`
	CostType             string         `json:"costType"`
	Image                ImageDTO       `json:"image"`
	SanitizedDescription string         `json:"sanitizedDescription"`
	SanitizedTooltip     string         `json:"sanitizedTooltip"`
	Effect               []([]float64)  `json:"effect"`
	Tooltip              string         `json:"tooltip"`
	MaxRank              int            `json:"maxrank"`
	CostBurn             string         `json:"costBurn"`
	RangeBurn            string         `json:"rangeBurn"`
	Range                interface{}    `json:"range"`
	Cooldown             []float64      `json:"cooldown"`
	Cost                 []int          `json:"cost"`
	Key                  string         `json:"key"`
	Description          string         `json:"description"`
	EffectBurn           []string       `json:"effectBurn"`
	AltImages            []ImageDTO     `json:"altimages"`
	Name                 string         `json:"name"`
}

// LevelTipDTO - Structure for champion level tip data retrieved using the API.
type LevelTipDTO struct {
	Effect []string `json:"effect"`
	Label  []string `json:"label"`
}

// SpellVarsDTO - Structure for spell vars data retrieved using the API.
type SpellVarsDTO struct {
	RanksWith string    `json:"ranksWith"`
	Dyn       string    `json:"dyn"`
	Link      string    `json:"link"`
	Coeff     []float64 `json:"coeff"`
	Key       string    `json:"key"`
}

// ItemListDTO - Structure for a list of all item data retrieved using the API.
type ItemListDTO struct {
	Data    map[string]ItemDTO `json:"data"`
	Version string             `json:"version"`
	Tree    []ItemTreeDTO      `json:"tree"`
	Groups  []GroupDTO         `json:"groups"`
	Type    string             `json:"type"`
}

// ItemTreeDTO - Structure for item tree data retrieved using the API.
type ItemTreeDTO struct {
	Header string   `json:"header"`
	Tags   []string `json:"tags"`
}

// ItemDTO - Structure for item data retrieved using the API.
type ItemDTO struct {
	Gold                 GoldDTO               `json:"gold"`
	Plaintext            string                `json:"plaintext"`
	HideFromAll          bool                  `json:"hideFromAll"`
	InStore              bool                  `json:"inStore"`
	Into                 []string              `json:"into"`
	ID                   int                   `json:"id"`
	Stats                InventoryDataStatsDTO `json:"stats"`
	Colloq               string                `json:"colloq"`
	Maps                 map[string]bool       `json:"maps"`
	SpecialRecipe        int                   `json:"specialRecipe"`
	Image                ImageDTO              `json:"image"`
	Description          string                `json:"description"`
	Tags                 []string              `json:"tags"`
	Effect               map[string]string     `json:"effect"`
	RequiredChampion     string                `json:"requiredChampion"`
	From                 []string              `json:"from"`
	Group                string                `json:"group"`
	ConsumeOnFull        bool                  `json:"consumeOnFull"`
	Name                 string                `json:"name"`
	Consumed             bool                  `json:"consumed"`
	SanitizedDescription string                `json:"sanitizedDescription"`
	Depth                int                   `json:"depth"`
	Stacks               int                   `json:"stacks"`
}

// GoldDTO - Structure for item gold data retrieved using the API.
type GoldDTO struct {
	Sell        int  `json:"sell"`
	Total       int  `json:"total"`
	Base        int  `json:"base"`
	Purchasable bool `json:"purchasable"`
}

// InventoryDataStatsDTO - Structure for inventory data stats retrieved using the API.
type InventoryDataStatsDTO struct {
	PercentCritDamageMod     float64 `json:"PercentCritDamageMod"`
	PercentSpellBlockMod     float64 `json:"PercentSpellBlockMod"`
	PercentHPRegenMod        float64 `json:"PercentHPRegenMod"`
	PercentMovementSpeedMod  float64 `json:"PercentMovementSpeedMod"`
	FlatSpellBlockMod        float64 `json:"FlatSpellBlockMod"`
	FlatCritDamageMod        float64 `json:"FlatCritDamageMod"`
	FlatEnergyPoolMod        float64 `json:"FlatEnergyPoolMod"`
	PercentLifeStealMod      float64 `json:"PercentLifeStealMod"`
	FlatMPPoolMod            float64 `json:"FlatMPPoolMod"`
	FlatMovementSpeedMod     float64 `json:"FlatMovementSpeedMod"`
	PercentAttackSpeedMod    float64 `json:"PercentAttackSpeedMod"`
	FlatBlockMod             float64 `json:"FlatBlockMod"`
	PercentBlockMod          float64 `json:"PercentBlockMod"`
	FlatEnergyRegenMod       float64 `json:"FlatEnergyRegenMod"`
	PercentSpellVampMod      float64 `json:"PercentSpellVampMod"`
	FlatMPRegenMod           float64 `json:"FlatMPRegenMod"`
	PercentDodgeMod          float64 `json:"PercentDodgeMod"`
	FlatAttackSpeedMod       float64 `json:"FlatAttackSpeedMod"`
	FlatArmorMod             float64 `json:"FlatArmorMod"`
	FlatHPRegenMod           float64 `json:"FlatHPRegenMod"`
	PercentMagicDamageMod    float64 `json:"PercentMagicDamageMod"`
	PercentMPPoolMod         float64 `json:"PercentMPPoolMod"`
	FlatMagicDamageMod       float64 `json:"FlatMagicDamageMod"`
	PercentMPRegenMod        float64 `json:"PercentMPRegenMod"`
	PercentPhysicalDamageMod float64 `json:"PercentPhysicalDamageMod"`
	FlatPhysicalDamageMod    float64 `json:"FlatPhysicalDamageMod"`
	PercentHPPoolMod         float64 `json:"PercentHPPoolMod"`
	PercentArmorMod          float64 `json:"PercentArmorMod"`
	PercentCritChanceMod     float64 `json:"PercentCritChanceMod"`
	PercentEXPBonus          float64 `json:"PercentEXPBonus"`
	FlatHPPoolMod            float64 `json:"FlatHPPoolMod"`
	FlatCritChanceMod        float64 `json:"FlatCritChanceMod"`
	FlatEXPBonus             float64 `json:"FlatEXPBonus"`
}

// GroupDTO - Structure to hold item group data retrieved using the API.
type GroupDTO struct {
	MaxGroupOwnable string `json:"MaxGroupOwnable"`
	Key             string `json:"key"`
}

// LanguageStringsDTO - Structure to hold language string data retrieved using the API.
type LanguageStringsDTO struct {
	Data    map[string]string `json:"data"`
	Version string            `json:"version"`
	Type    string            `json:"type"`
}

// LanguageList - A list of strings consisting of languages retrieved using the API.
type LanguageList []string

// MapDataDTO -  Structure to hold map data retrieved using the API.
type MapDataDTO struct {
	Data    map[string]MapDetailsDTO `json:"data"`
	Version string                   `json:"version"`
	Type    string                   `json:"type"`
}

// MapDetailsDTO - Structure to hold map details data retrieved using the API.
type MapDetailsDTO struct {
	MapName               string   `json:"mapName"`
	Image                 ImageDTO `json:"image"`
	MapID                 int64    `json:"mapID"`
	UnpurchasableItemList []int64  `json:"unpurchasableItemList"`
}

// MasteryListDTO - Structure to hold a list of mastery data retrieved using the API.
type MasteryListDTO struct {
	Data    map[string]MasteryDTO `json:"data"`
	Version string                `json:"version"`
	Tree    MasteryTreeDTO        `json:"tree"`
	Type    string                `json:"type"`
}

// MasteryTreeDTO - Structure to hold mastery tree data retrieved using the API.
type MasteryTreeDTO struct {
	Resolve  []MasteryTreeListDTO `json:"Resolve"`
	Ferocity []MasteryTreeListDTO `json:"Ferocity"`
	Cunning  []MasteryTreeListDTO `json:"Cunning"`
}

// MasteryTreeListDTO - Structure to hold a mastery tree list data retrieved using the API.
type MasteryTreeListDTO struct {
	MasteryTreeItems []MasteryTreeItemDTO `json:"masteryTreeItems"`
}

// MasteryTreeItemDTO - Structure to hold mastery tree item data retrieved using the API.
type MasteryTreeItemDTO struct {
	MasteryID int    `json:"masteryId"`
	Prereq    string `json:"prereq"`
}

// MasteryDTO - Structure to hold mastery data retrieved using the API.
type MasteryDTO struct {
	Prereq               string   `json:"prereq"`
	MasteryTree          string   `json:"masteryTree"`
	Name                 string   `json:"name"`
	Ranks                int      `json:"ranks"`
	Image                ImageDTO `json:"image"`
	SanitizedDescription []string `json:"sanitizedDescription"`
	ID                   int      `json:"id"`
	Description          []string `json:"description"`
}

// ProfileIconDataDTO - Structure to hold profile icon data retrieved using the API.
type ProfileIconDataDTO struct {
	Data    map[string]ProfileIconDetailsDTO `json:"data"`
	Version string                           `json:"version"`
	Type    string                           `json:"type"`
}

// ProfileIconDetailsDTO - Structure to hold profile icon details retrieved using the API.
type ProfileIconDetailsDTO struct {
	Image ImageDTO `json:"image"`
	ID    int64    `json:"id"`
}

// RealmDTO - Structure to hold realm data retrieved using the API.
type RealmDTO struct {
	LG             string            `json:"lg"`
	DD             string            `json:"dd"`
	L              string            `json:"l"`
	N              map[string]string `json:"n"`
	ProfileIconMax int               `json:"profileiconmax"`
	Store          string            `json:"store"`
	V              string            `json:"v"`
	CDN            string            `json:"cdn"`
	CSS            string            `json:"css"`
}

// ReforgedRunePaths - Holds a list of ReforgedRunePathDTO retrieved using the API.
type ReforgedRunePaths []ReforgedRunePathDTO

// ReforgedRunePathDTO - Structure to hold ReforgedRunePath data retrieved using the API.
type ReforgedRunePathDTO struct {
	Slots []ReforgedRuneSlotDTO `json:"slots"`
	Icon  string                `json:"icon"`
	ID    int                   `json:"id"`
	Key   string                `json:"key"`
	Name  string                `json:"name"`
}

// ReforgedRuneSlotDTO - Structure to hold ReforgedRuneSlot data retrieved using the API.
type ReforgedRuneSlotDTO struct {
	Runes []ReforgedRuneDTO `json:"runes"`
}

// ReforgedRunes - Structure to hold a list of ReforgedRuneDTO data retrieved using the API.
type ReforgedRunes []ReforgedRuneDTO

// ReforgedRuneDTO - Structure to hold ReforgedRune data retrieved using the API.
type ReforgedRuneDTO struct {
	RunePathName string `json:"runePathName"`
	RunePathID   int    `json:"runePathId"`
	Name         string `json:"name"`
	ID           int    `json:"id"`
	Key          string `json:"key"`
	ShortDesc    string `json:"shortDesc"`
	LongDesc     string `json:"longDesc"`
	Icon         string `json:"icon"`
}

// RuneListDTO - Structure to hold rune list data retrieved using the API.
type RuneListDTO struct {
	Data    map[string]RuneDTO `json:"data"`
	Version string             `json:"version"`
	Type    string             `json:"type"`
}

// RuneDTO - Structure to hold rune data retrieved using the API.
type RuneDTO struct {
	Stats                RuneStatsDTO `json:"stats"`
	Name                 string       `json:"name"`
	Tags                 []string     `json:"tags"`
	Image                ImageDTO     `json:"image"`
	SanitizedDescription string       `json:"sanitizedDescription"`
	Rune                 MetaDataDTO  `json:"rune"`
	ID                   int          `json:"id"`
	Description          string       `json:"description"`
}

// RuneStatsDTO - Structure to hold rune stats data retrieved using the API.
type RuneStatsDTO struct {
	PercentTimeDeadModPerLevel         float64 `json:"PercentTimeDeadModPerLevel"`
	PercentArmorPenetrationModPerLevel float64 `json:"PercentArmorPenetrationModPerLevel"`
	PercentCritDamageMod               float64 `json:"PercentCritDamageMod"`
	PercentSpellBlockMod               float64 `json:"PercentSpellBlockMod"`
	PercentHPRegenMod                  float64 `json:"PercentHPRegenMod"`
	PercentMovementSpeedMod            float64 `json:"PercentMovementSpeedMod"`
	FlatSpellBlockMod                  float64 `json:"FlatSpellBlockMod"`
	FlatEnergyRegenModPerLevel         float64 `json:"FlatEnergyRegenModPerLevel"`
	FlatEnergyPoolMod                  float64 `json:"FlatEnergyPoolMod"`
	FlatMagicPenetrationModPerLevel    float64 `json:"FlatMagicPenetrationModPerLevel"`
	PercentLifeStealMod                float64 `json:"PercentLifeStealMod"`
	FlatMPPoolMod                      float64 `json:"FlatMPPoolMod"`
	PercentCooldownMod                 float64 `json:"PercentCooldownMod"`
	PercentMagicPenetrationMod         float64 `json:"PercentMagicPenetrationMod"`
	FlatArmorPenetrationModPerLevel    float64 `json:"FlatArmorPenetrationModPerLevel"`
	FlatMovementSpeedMod               float64 `json:"FlatMovementSpeedMod"`
	FlatTimeDeadModPerLevel            float64 `json:"FlatTimeDeadModPerLevel"`
	FlatArmorModPerLevel               float64 `json:"FlatArmorModPerLevel"`
	PercentAttackSpeedMod              float64 `json:"PercentAttackSpeedMod"`
	FlatDodgeModPerLevel               float64 `json:"FlatDodgeModPerLevel"`
	PercentMagicDamageMod              float64 `json:"PercentMagicDamageMod"`
	PercentBlockMod                    float64 `json:"PercentBlockMod"`
	FlatDodgeMod                       float64 `json:"FlatDodgeMod"`
	FlatEnergyRegenMod                 float64 `json:"FlatEnergyRegenMod"`
	FlatHPModPerLevel                  float64 `json:"FlatHPModPerLevel"`
	PercentAttackSpeedModPerLevel      float64 `json:"PercentAttackSpeedModPerLevel"`
	PercentSpellVampMod                float64 `json:"PercentSpellVampMod"`
	FlatMPRegenMod                     float64 `json:"FlatMPRegenMod"`
	PercentHPPoolMod                   float64 `json:"PercentHPPoolMod"`
	PercentDodgeMod                    float64 `json:"PercentDodgeMod"`
	FlatAttackSpeedMod                 float64 `json:"FlatAttackSpeedMod"`
	FlatArmorMod                       float64 `json:"FlatArmorMod"`
	FlatMagicDamageModPerLevel         float64 `json:"FlatMagicDamageModPerLevel"`
	FlatHPRegenMod                     float64 `json:"FlatHPRegenMod"`
	PercentPhysicalDamageMod           float64 `json:"PercentPhysicalDamageMod"`
	FlatCritChanceModPerLevel          float64 `json:"FlatCritChanceModPerLevel"`
	FlatSpellBlockModPerLevel          float64 `json:"FlatSpellBlockModPerLevel"`
	PercentTimeDeadMod                 float64 `json:"PercentTimeDeadMod"`
	FlatBlockMod                       float64 `json:"FlatBlockMod"`
	PercentMPPoolMod                   float64 `json:"PercentMPPoolMod"`
	FlatMagicDamageMod                 float64 `json:"FlatMagicDamageMod"`
	PercentMPRegenMod                  float64 `json:"PercentMPRegenMod"`
	PercentMovementSpeedModPerLevel    float64 `json:"PercentMovementSpeedModPerLevel"`
	PercentCooldownModPerLevel         float64 `json:"PercentCooldownModPerLevel"`
	FlatMPModPerLevel                  float64 `json:"FlatMPModPerLevel"`
	FlatEnergyModPerLevel              float64 `json:"FlatEnergyModPerLevel"`
	FlatPhysicalDamageMod              float64 `json:"FlatPhysicalDamageMod"`
	FlatHPRegenModPerLevel             float64 `json:"FlatHPRegenModPerLevel"`
	FlatCritDamageMod                  float64 `json:"FlatCritDamageMod"`
	PercentArmorMod                    float64 `json:"PercentArmorMod"`
	FlatMagicPenetrationMod            float64 `json:"FlatMagicPenetrationMod"`
	PercentCritChanceMod               float64 `json:"PercentCritChanceMod"`
	FlatPhysicalDamageModPerLevel      float64 `json:"FlatPhysicalDamageModPerLevel"`
	PercentArmorPenetrationMod         float64 `json:"PercentArmorPenetrationMod"`
	PercentEXPBonus                    float64 `json:"PercentEXPBonus"`
	FlatMPRegenModPerLevel             float64 `json:"FlatMPRegenModPerLevel"`
	PercentMagicPenetrationModPerLevel float64 `json:"PercentMagicPenetrationModPerLevel"`
	FlatTimeDeadMod                    float64 `json:"FlatTimeDeadMod"`
	FlatMovementSpeedModPerLevel       float64 `json:"FlatMovementSpeedModPerLevel"`
	FlatGoldPer10Mod                   float64 `json:"FlatGoldPer10Mod"`
	FlatArmorPenetrationMod            float64 `json:"FlatArmorPenetrationMod"`
	FlatCritDamageModPerLevel          float64 `json:"FlatCritDamageModPerLevel"`
	FlatHPPoolMod                      float64 `json:"FlatHPPoolMod"`
	FlatCritChanceMod                  float64 `json:"FlatCritChanceMod"`
	FlatEXPBonus                       float64 `json:"FlatEXPBonus"`
}

// MetaDataDTO - Structure to hold meta data retrieved using the API.
type MetaDataDTO struct {
	Tier   string `json:"tier"`
	Type   string `json:"type"`
	IsRune bool   `json:"isRune"`
}

// SummonerSpellListDTO - Structure to hold summoner spell list data retrieved using the API.
type SummonerSpellListDTO struct {
	Data    map[string]SummonerSpellDTO `json:"data"`
	Version string                      `json:"version"`
	Type    string                      `json:"type"`
}

// SummonerSpellDTO - Structure to hold summoner spell data retrieved using the API.
type SummonerSpellDTO struct {
	Vars                 []SpellVarsDTO `json:"vars"`
	Image                ImageDTO       `json:"image"`
	CostBurn             string         `json:"costBurn"`
	Cooldown             []float64      `json:"cooldown"`
	EffectBurn           []string       `json:"effectBurn"`
	ID                   int            `json:"id"`
	CooldownBurn         string         `json:"cooldownBurn"`
	Tooltip              string         `json:"tooltip"`
	Maxrank              int            `json:"maxrank"`
	RangeBurn            string         `json:"rangeBurn"`
	Description          string         `json:"description"`
	Effect               []([]float64)  `json:"effect"`
	Key                  string         `json:"key"`
	Leveltip             LevelTipDTO    `json:"leveltip"`
	Modes                []string       `json:"modes"`
	Resource             string         `json:"resource"`
	Name                 string         `json:"name"`
	CostType             string         `json:"costType"`
	SanitizedDescription string         `json:"sanitizedDescription"`
	SanitizedTooltip     string         `json:"sanitizedTooltip"`
	Range                interface{}    `json:"range"`
	Cost                 []int          `json:"cost"`
	SummonerLevel        int            `json:"summonerLevel"`
}

// VersionList - A list of strings consisting of versions retrieved using the API.
type VersionList []string

// ChampionStaticData fetches static data for all champions.
func ChampionStaticData(api *RiotAPI) *ChampionList {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/champions?locale=en_US&tags=all&dataById=false&api_key=" + api.APIKey

	championList := new(ChampionList)

	NewAPICall(url, &championList)

	return championList
}

// ChampionStaticDataByID fetches static data for a specific champion given the id.
func ChampionStaticDataByID(api *RiotAPI, championID string) *ChampionDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/champions/" + championID + "?locale=en_US&tags=all&api_key=" + api.APIKey

	champion := new(ChampionDTO)

	NewAPICall(url, &champion)

	return champion
}

// ItemStaticData fetches static data for all items.
func ItemStaticData(api *RiotAPI) *ItemListDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/items?locale=en_US&tags=all&api_key=" + api.APIKey

	itemList := new(ItemListDTO)

	NewAPICall(url, &itemList)

	return itemList
}

// ItemStaticDataByID fetches static data for a specific item.
func ItemStaticDataByID(api *RiotAPI, itemID string) *ItemDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/items/" + itemID + "?locale=en_US&tags=all&api_key=" + api.APIKey

	item := new(ItemDTO)

	NewAPICall(url, &item)

	return item
}

// LanguageStrings fetches language string static data.
func LanguageStrings(api *RiotAPI) *LanguageStringsDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/language-strings?locale=en_US&api_key=" + api.APIKey

	languageStrings := new(LanguageStringsDTO)

	NewAPICall(url, &languageStrings)

	return languageStrings
}

// Languages fetches the language list static data.
func Languages(api *RiotAPI) LanguageList {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/languages?api_key=" + api.APIKey

	languages := new(LanguageList)

	NewAPICall(url, languages)

	return *languages
}

// MapStaticData fetches the map static data for all maps.
func MapStaticData(api *RiotAPI) *MapDataDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/maps?locale=en_US&api_key=" + api.APIKey

	mapData := new(MapDataDTO)

	NewAPICall(url, &mapData)

	return mapData
}

// MasteryStaticData fetches mastery static data for all masteries.
func MasteryStaticData(api *RiotAPI) *MasteryListDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/masteries?locale=en_US&tags=all&api_key=" + api.APIKey

	masteries := new(MasteryListDTO)

	NewAPICall(url, &masteries)

	return masteries
}

// MasteryStaticDataByID fetches mastery static data for a specific mastery.
func MasteryStaticDataByID(api *RiotAPI, masteryID string) *MasteryDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/masteries/" + masteryID + "?locale=en_US&tags=all&api_key=" + api.APIKey

	mastery := new(MasteryDTO)

	NewAPICall(url, &mastery)

	return mastery
}

// ProfileIconStaticData fetches profile icon static data for all profile icons.
func ProfileIconStaticData(api *RiotAPI) *ProfileIconDataDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/profile-icons?locale=en_US&api_key=" + api.APIKey

	icons := new(ProfileIconDataDTO)

	NewAPICall(url, &icons)

	return icons
}

// RealmStaticData fetches realm static data.
func RealmStaticData(api *RiotAPI) *RealmDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/realms?api_key=" + api.APIKey

	realm := new(RealmDTO)

	NewAPICall(url, &realm)

	return realm
}

// ReforgedRunePathsStaticData fetches all reforged rune path static data.
func ReforgedRunePathsStaticData(api *RiotAPI) *ReforgedRunePaths {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/reforged-rune-paths?api_key=" + api.APIKey

	reforgedRunePaths := new(ReforgedRunePaths)

	NewAPICall(url, &reforgedRunePaths)

	return reforgedRunePaths
}

// ReforgedRunePathsStaticDataByID fetches all reforged rune path static data.
func ReforgedRunePathsStaticDataByID(api *RiotAPI, id string) *ReforgedRunePathDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/reforged-rune-paths/" + id + "?api_key=" + api.APIKey

	reforgedRunePath := new(ReforgedRunePathDTO)

	NewAPICall(url, &reforgedRunePath)

	return reforgedRunePath
}

// ReforgedRuneStaticData fetches all reforged rune static data.
func ReforgedRuneStaticData(api *RiotAPI) *ReforgedRunes {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/reforged-runes?api_key=" + api.APIKey

	reforgedRunes := new(ReforgedRunes)

	NewAPICall(url, &reforgedRunes)

	return reforgedRunes
}

// ReforgedRuneStaticDataByID fetches all reforged rune static data.
func ReforgedRuneStaticDataByID(api *RiotAPI, id string) *ReforgedRuneDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/reforged-runes/" + id + "?api_key=" + api.APIKey

	reforgedRune := new(ReforgedRuneDTO)

	NewAPICall(url, &reforgedRune)

	return reforgedRune
}

// RuneStaticData feches all rune static data for all runes.
func RuneStaticData(api *RiotAPI) *RuneListDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/runes?locale=en_US&tags=all&api_key=" + api.APIKey

	runes := new(RuneListDTO)

	NewAPICall(url, &runes)

	return runes
}

// RuneStaticDataByID fetches all rune static data for a specific rune.
func RuneStaticDataByID(api *RiotAPI, runeID string) *RuneDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/runes/" + runeID + "?locale=en_US&tags=all&api_key=" + api.APIKey

	rune := new(RuneDTO)

	NewAPICall(url, &rune)

	return rune
}

// SummonerSpellData fetches all summoner spell static data for all summoner spells.
func SummonerSpellData(api *RiotAPI) *SummonerSpellListDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/summoner-spells?locale=en_US&dataById=true&tags=all&api_key=" + api.APIKey

	summonerSpells := new(SummonerSpellListDTO)

	NewAPICall(url, &summonerSpells)

	return summonerSpells
}

// SummonerSpellDataByID fetches all summoner spell static data for a specific summoner spell.
func SummonerSpellDataByID(api *RiotAPI, summonerSpellID string) *SummonerSpellDTO {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/summoner-spells/" + summonerSpellID + "?locale=en_US&tags=all&api_key=" + api.APIKey

	summonerSpell := new(SummonerSpellDTO)

	NewAPICall(url, &summonerSpell)

	return summonerSpell
}

// Versions fetches the version list static data.
func Versions(api *RiotAPI) VersionList {
	url := "https://" + api.Region + ".api.riotgames.com/lol/static-data/" + api.Version + "/versions?api_key=" + api.APIKey

	versions := new(VersionList)

	NewAPICall(url, versions)

	return *versions
}
