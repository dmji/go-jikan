package components

import (
	"encoding/json"
	"fmt"
)

// MangaSearchQueryOrderby - Available Manga order_by properties
type MangaSearchQueryOrderby string

const (
	MangaSearchQueryOrderbyMalID      MangaSearchQueryOrderby = "mal_id"
	MangaSearchQueryOrderbyTitle      MangaSearchQueryOrderby = "title"
	MangaSearchQueryOrderbyStartDate  MangaSearchQueryOrderby = "start_date"
	MangaSearchQueryOrderbyEndDate    MangaSearchQueryOrderby = "end_date"
	MangaSearchQueryOrderbyChapters   MangaSearchQueryOrderby = "chapters"
	MangaSearchQueryOrderbyVolumes    MangaSearchQueryOrderby = "volumes"
	MangaSearchQueryOrderbyScore      MangaSearchQueryOrderby = "score"
	MangaSearchQueryOrderbyScoredBy   MangaSearchQueryOrderby = "scored_by"
	MangaSearchQueryOrderbyRank       MangaSearchQueryOrderby = "rank"
	MangaSearchQueryOrderbyPopularity MangaSearchQueryOrderby = "popularity"
	MangaSearchQueryOrderbyMembers    MangaSearchQueryOrderby = "members"
	MangaSearchQueryOrderbyFavorites  MangaSearchQueryOrderby = "favorites"
)

func (e MangaSearchQueryOrderby) ToPointer() *MangaSearchQueryOrderby {
	return &e
}

func (e *MangaSearchQueryOrderby) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mal_id":
		fallthrough
	case "title":
		fallthrough
	case "start_date":
		fallthrough
	case "end_date":
		fallthrough
	case "chapters":
		fallthrough
	case "volumes":
		fallthrough
	case "score":
		fallthrough
	case "scored_by":
		fallthrough
	case "rank":
		fallthrough
	case "popularity":
		fallthrough
	case "members":
		fallthrough
	case "favorites":
		*e = MangaSearchQueryOrderby(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MangaSearchQueryOrderby: %v", v)
	}
}
