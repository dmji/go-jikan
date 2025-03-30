package components

import (
	"encoding/json"
	"fmt"
)

// TopMangaFilter - Top items filter types
type TopMangaFilter string

const (
	TopMangaFilterPublishing   TopMangaFilter = "publishing"
	TopMangaFilterUpcoming     TopMangaFilter = "upcoming"
	TopMangaFilterBypopularity TopMangaFilter = "bypopularity"
	TopMangaFilterFavorite     TopMangaFilter = "favorite"
)

func (e TopMangaFilter) ToPointer() *TopMangaFilter {
	return &e
}

func (e *TopMangaFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "publishing":
		fallthrough
	case "upcoming":
		fallthrough
	case "bypopularity":
		fallthrough
	case "favorite":
		*e = TopMangaFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for TopMangaFilter: %v", v)
	}
}
