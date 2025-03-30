package components

import (
	"encoding/json"
	"fmt"
)

// TopAnimeFilter - Top items filter types
type TopAnimeFilter string

const (
	TopAnimeFilterAiring       TopAnimeFilter = "airing"
	TopAnimeFilterUpcoming     TopAnimeFilter = "upcoming"
	TopAnimeFilterBypopularity TopAnimeFilter = "bypopularity"
	TopAnimeFilterFavorite     TopAnimeFilter = "favorite"
)

func (e TopAnimeFilter) ToPointer() *TopAnimeFilter {
	return &e
}

func (e *TopAnimeFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "airing":
		fallthrough
	case "upcoming":
		fallthrough
	case "bypopularity":
		fallthrough
	case "favorite":
		*e = TopAnimeFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TopAnimeFilter: %v", v)
	}
}
