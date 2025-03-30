package components

import (
	"encoding/json"
	"fmt"
)

// SearchQuerySort - Search query sort direction
type SearchQuerySort string

const (
	SearchQuerySortDesc SearchQuerySort = "desc"
	SearchQuerySortAsc  SearchQuerySort = "asc"
)

func (e SearchQuerySort) ToPointer() *SearchQuerySort {
	return &e
}

func (e *SearchQuerySort) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "desc":
		fallthrough
	case "asc":
		*e = SearchQuerySort(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SearchQuerySort: %v", v)
	}
}
