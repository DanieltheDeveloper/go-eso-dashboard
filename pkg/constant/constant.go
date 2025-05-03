package constant

import "time"

// fetchTimeout is the duration for which the data is cached.
const FetchTimeout = 10 * time.Second
const Unreachable = "Unreachable"
// PlayerCountCacheDuration is the duration for which the data is cached.
const PlayerCountCacheDuration = 3 * time.Minute