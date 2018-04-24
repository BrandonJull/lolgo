package main

import (
	"fmt"
	"lolgo/api"
)

func main() {
	// Riot API key (http://developer.riotgames.com/)
	api := api.NewRiotAPI("YOUR_KEY_HERE", "na1", "v3")

	// Usage example.
	fmt.Println(api.GetSummonerByName("TGB").SummonerLevel)
	fmt.Println(api.GetReforgedRuneStaticDataByID("8410"))
}
