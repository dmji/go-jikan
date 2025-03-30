package components

import (
	"encoding/json"
	"fmt"
)

// GenreQueryFilter - Filter genres by type
type GenreQueryFilter string

const (
	GenreQueryFilterGenres         GenreQueryFilter = "genres"
	GenreQueryFilterExplicitGenres GenreQueryFilter = "explicit_genres"
	GenreQueryFilterThemes         GenreQueryFilter = "themes"
	GenreQueryFilterDemographics   GenreQueryFilter = "demographics"
)

func (e GenreQueryFilter) ToPointer() *GenreQueryFilter {
	return &e
}

func (e *GenreQueryFilter) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "genres":
		fallthrough
	case "explicit_genres":
		fallthrough
	case "themes":
		fallthrough
	case "demographics":
		*e = GenreQueryFilter(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GenreQueryFilter: %v", v)
	}
}
