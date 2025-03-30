package components

import (
	"encoding/json"
	"fmt"
)

// AnimeSearchQueryOrderby - Available Anime order_by properties
type AnimeSearchQueryOrderby string

const (
	AnimeSearchQueryOrderbyMalID      AnimeSearchQueryOrderby = "mal_id"
	AnimeSearchQueryOrderbyTitle      AnimeSearchQueryOrderby = "title"
	AnimeSearchQueryOrderbyStartDate  AnimeSearchQueryOrderby = "start_date"
	AnimeSearchQueryOrderbyEndDate    AnimeSearchQueryOrderby = "end_date"
	AnimeSearchQueryOrderbyEpisodes   AnimeSearchQueryOrderby = "episodes"
	AnimeSearchQueryOrderbyScore      AnimeSearchQueryOrderby = "score"
	AnimeSearchQueryOrderbyScoredBy   AnimeSearchQueryOrderby = "scored_by"
	AnimeSearchQueryOrderbyRank       AnimeSearchQueryOrderby = "rank"
	AnimeSearchQueryOrderbyPopularity AnimeSearchQueryOrderby = "popularity"
	AnimeSearchQueryOrderbyMembers    AnimeSearchQueryOrderby = "members"
	AnimeSearchQueryOrderbyFavorites  AnimeSearchQueryOrderby = "favorites"
)

func (e AnimeSearchQueryOrderby) ToPointer() *AnimeSearchQueryOrderby {
	return &e
}

func (e *AnimeSearchQueryOrderby) UnmarshalJSON(data []byte) error {
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
	case "episodes":
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
		*e = AnimeSearchQueryOrderby(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnimeSearchQueryOrderby: %v", v)
	}
}
