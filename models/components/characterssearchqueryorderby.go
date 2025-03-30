package components

import (
	"encoding/json"
	"fmt"
)

// CharactersSearchQueryOrderby - Available Character order_by properties
type CharactersSearchQueryOrderby string

const (
	CharactersSearchQueryOrderbyMalID     CharactersSearchQueryOrderby = "mal_id"
	CharactersSearchQueryOrderbyName      CharactersSearchQueryOrderby = "name"
	CharactersSearchQueryOrderbyFavorites CharactersSearchQueryOrderby = "favorites"
)

func (e CharactersSearchQueryOrderby) ToPointer() *CharactersSearchQueryOrderby {
	return &e
}

func (e *CharactersSearchQueryOrderby) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mal_id":
		fallthrough
	case "name":
		fallthrough
	case "favorites":
		*e = CharactersSearchQueryOrderby(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CharactersSearchQueryOrderby: %v", v)
	}
}
