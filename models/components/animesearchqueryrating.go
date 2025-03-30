package components

import (
	"encoding/json"
	"fmt"
)

// AnimeSearchQueryRating - Available Anime audience ratings<br><br><b>Ratings</b><br><ul><li>G - All Ages</li><li>PG - Children</li><li>PG-13 - Teens 13 or older</li><li>R - 17+ (violence & profanity)</li><li>R+ - Mild Nudity</li><li>Rx - Hentai</li></ul>
type AnimeSearchQueryRating string

const (
	AnimeSearchQueryRatingG    AnimeSearchQueryRating = "g"
	AnimeSearchQueryRatingPg   AnimeSearchQueryRating = "pg"
	AnimeSearchQueryRatingPg13 AnimeSearchQueryRating = "pg13"
	AnimeSearchQueryRatingR17  AnimeSearchQueryRating = "r17"
	AnimeSearchQueryRatingR    AnimeSearchQueryRating = "r"
	AnimeSearchQueryRatingRx   AnimeSearchQueryRating = "rx"
)

func (e AnimeSearchQueryRating) ToPointer() *AnimeSearchQueryRating {
	return &e
}

func (e *AnimeSearchQueryRating) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "g":
		fallthrough
	case "pg":
		fallthrough
	case "pg13":
		fallthrough
	case "r17":
		fallthrough
	case "r":
		fallthrough
	case "rx":
		*e = AnimeSearchQueryRating(v)
		return nil
	default:
		return fmt.Errorf("invalid value for AnimeSearchQueryRating: %v", v)
	}
}
