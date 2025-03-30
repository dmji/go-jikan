package components

import (
	"encoding/json"
	"fmt"
)

// ClubSearchQueryCategory - Club Search Query Category
type ClubSearchQueryCategory string

const (
	ClubSearchQueryCategoryAnime                  ClubSearchQueryCategory = "anime"
	ClubSearchQueryCategoryManga                  ClubSearchQueryCategory = "manga"
	ClubSearchQueryCategoryActorsAndArtists       ClubSearchQueryCategory = "actors_and_artists"
	ClubSearchQueryCategoryCharacters             ClubSearchQueryCategory = "characters"
	ClubSearchQueryCategoryCitiesAndNeighborhoods ClubSearchQueryCategory = "cities_and_neighborhoods"
	ClubSearchQueryCategoryCompanies              ClubSearchQueryCategory = "companies"
	ClubSearchQueryCategoryConventions            ClubSearchQueryCategory = "conventions"
	ClubSearchQueryCategoryGames                  ClubSearchQueryCategory = "games"
	ClubSearchQueryCategoryJapan                  ClubSearchQueryCategory = "japan"
	ClubSearchQueryCategoryMusic                  ClubSearchQueryCategory = "music"
	ClubSearchQueryCategoryOther                  ClubSearchQueryCategory = "other"
	ClubSearchQueryCategorySchools                ClubSearchQueryCategory = "schools"
)

func (e ClubSearchQueryCategory) ToPointer() *ClubSearchQueryCategory {
	return &e
}

func (e *ClubSearchQueryCategory) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "anime":
		fallthrough
	case "manga":
		fallthrough
	case "actors_and_artists":
		fallthrough
	case "characters":
		fallthrough
	case "cities_and_neighborhoods":
		fallthrough
	case "companies":
		fallthrough
	case "conventions":
		fallthrough
	case "games":
		fallthrough
	case "japan":
		fallthrough
	case "music":
		fallthrough
	case "other":
		fallthrough
	case "schools":
		*e = ClubSearchQueryCategory(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClubSearchQueryCategory: %v", v)
	}
}
