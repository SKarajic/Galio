package urls

// URL is a list of url constants
type URL string

const (
	// API is a formatting URL for the RiotGames league of legends API, the first value should be the region and the second
	// region should be the endpoint
	API = "https://%s.api.riotgames.com/lol/%v"
)
