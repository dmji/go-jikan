package components

import (
	"encoding/json"
	"fmt"
)

// AnimeSearchQueryStatus - Available Anime statuses
type AnimeSearchQueryStatus string

const (
	AnimeSearchQueryStatusAiring   AnimeSearchQueryStatus = "airing"
	AnimeSearchQueryStatusComplete AnimeSearchQueryStatus = "complete"
	AnimeSearchQueryStatusUpcoming AnimeSearchQueryStatus = "upcoming"
)

func (e AnimeSearchQueryStatus) ToPointer() *AnimeSearchQueryStatus {
	return &e
}

func (e *AnimeSearchQueryStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "airing":
		fallthrough
	case "complete":
		fallthrough
	case "upcoming":
		*e = AnimeSearchQueryStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnimeSearchQueryStatus: %v", v)
	}
}
