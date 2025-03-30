package components

import (
	"encoding/json"
	"fmt"
)

// AnimeSearchQueryType - Available Anime types
type AnimeSearchQueryType string

const (
	AnimeSearchQueryTypeTv        AnimeSearchQueryType = "tv"
	AnimeSearchQueryTypeMovie     AnimeSearchQueryType = "movie"
	AnimeSearchQueryTypeOva       AnimeSearchQueryType = "ova"
	AnimeSearchQueryTypeSpecial   AnimeSearchQueryType = "special"
	AnimeSearchQueryTypeOna       AnimeSearchQueryType = "ona"
	AnimeSearchQueryTypeMusic     AnimeSearchQueryType = "music"
	AnimeSearchQueryTypeCm        AnimeSearchQueryType = "cm"
	AnimeSearchQueryTypePv        AnimeSearchQueryType = "pv"
	AnimeSearchQueryTypeTvSpecial AnimeSearchQueryType = "tv_special"
)

func (e AnimeSearchQueryType) ToPointer() *AnimeSearchQueryType {
	return &e
}

func (e *AnimeSearchQueryType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "tv":
		fallthrough
	case "movie":
		fallthrough
	case "ova":
		fallthrough
	case "special":
		fallthrough
	case "ona":
		fallthrough
	case "music":
		fallthrough
	case "cm":
		fallthrough
	case "pv":
		fallthrough
	case "tv_special":
		*e = AnimeSearchQueryType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnimeSearchQueryType: %v", v)
	}
}
