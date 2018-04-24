package api

// ShardStatus - structure for ShardStatus retrieved by the API.
type ShardStatus struct {
	Name      string    `json:"name"`
	RegionTag string    `json:"region_tag"`
	Hostname  string    `json:"hostname"`
	Services  []Service `json:"services"`
	Slug      string    `json:"slug"`
	Locales   []string  `json:"locales"`
}

// Service - Strucutre for Service retrieved by the API.
type Service struct {
	Status    string     `json:"status"`
	Incidents []Incident `json:"incidents"`
	Name      string     `json:"name"`
	Slug      string     `json:"slug"`
}

// Incident - Structure for Incident retrieved by the API.
type Incident struct {
	Active    bool      `json:"active"`
	CreatedAt string    `json:"created_at"`
	ID        int64     `json:"id"`
	Updates   []Message `json:"updatees"`
}

// Message - Structure for Message retrieved by the API.
type Message struct {
	Severity     string        `json:"severity"`
	Author       string        `json:"author"`
	CreatedAt    string        `json:"created_at"`
	Translations []Translation `json:"translations"`
	UpdatedAt    string        `json:"updated_at"`
	Content      string        `json:"content"`
	ID           string        `json:"id"`
}

// Translation - Structure for Translation retrieved by the API.
type Translation struct {
	Locale    string `json:"locale"`
	Content   string `json:"content"`
	UpdatedAt string `json:"updated_at"`
}

// ShardData fetches shard data for a specific server.
func ShardData(api *RiotAPI) *ShardStatus {
	url := "https://" + api.Region + ".api.riotgames.com/lol/status/" + api.Version + "/shard-data/" + "?api_key=" + api.APIKey
	shardStatus := new(ShardStatus)

	NewAPICall(url, &shardStatus)

	return shardStatus
}
