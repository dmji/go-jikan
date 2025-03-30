package components

import (
	"encoding/json"
	"fmt"
)

// PeopleSearchQueryOrderby - Available People order_by properties
type PeopleSearchQueryOrderby string

const (
	PeopleSearchQueryOrderbyMalID     PeopleSearchQueryOrderby = "mal_id"
	PeopleSearchQueryOrderbyName      PeopleSearchQueryOrderby = "name"
	PeopleSearchQueryOrderbyBirthday  PeopleSearchQueryOrderby = "birthday"
	PeopleSearchQueryOrderbyFavorites PeopleSearchQueryOrderby = "favorites"
)

func (e PeopleSearchQueryOrderby) ToPointer() *PeopleSearchQueryOrderby {
	return &e
}

func (e *PeopleSearchQueryOrderby) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mal_id":
		fallthrough
	case "name":
		fallthrough
	case "birthday":
		fallthrough
	case "favorites":
		*e = PeopleSearchQueryOrderby(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PeopleSearchQueryOrderby: %v", v)
	}
}
