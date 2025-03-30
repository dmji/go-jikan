package components

import (
	"encoding/json"
	"fmt"
)

// ClubSearchQueryOrderby - Club Search Query OrderBy
type ClubSearchQueryOrderby string

const (
	ClubSearchQueryOrderbyMalID        ClubSearchQueryOrderby = "mal_id"
	ClubSearchQueryOrderbyName         ClubSearchQueryOrderby = "name"
	ClubSearchQueryOrderbyMembersCount ClubSearchQueryOrderby = "members_count"
	ClubSearchQueryOrderbyCreated      ClubSearchQueryOrderby = "created"
)

func (e ClubSearchQueryOrderby) ToPointer() *ClubSearchQueryOrderby {
	return &e
}

func (e *ClubSearchQueryOrderby) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mal_id":
		fallthrough
	case "name":
		fallthrough
	case "members_count":
		fallthrough
	case "created":
		*e = ClubSearchQueryOrderby(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClubSearchQueryOrderby: %v", v)
	}
}
