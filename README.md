# lolgo
A wrapper for the League of Legends API written in Go using the [Official API](http://developer.riotgames.com/)

## Setup

### 1. Installation

Install and import the lolgo package:

```
go get github.com/brandonjull/lolgo
```

### 2. Obtain a Riot Games API key

Sign in with your account at: http://developer.riotgames.com/ and generate an API key.

## Usage Example

```
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
```
