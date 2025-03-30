package components

import (
	"encoding/json"
	"fmt"
)

// ClubSearchQueryType - Club Search Query Type
type ClubSearchQueryType string

const (
	ClubSearchQueryTypePublic  ClubSearchQueryType = "public"
	ClubSearchQueryTypePrivate ClubSearchQueryType = "private"
	ClubSearchQueryTypeSecret  ClubSearchQueryType = "secret"
)

func (e ClubSearchQueryType) ToPointer() *ClubSearchQueryType {
	return &e
}

func (e *ClubSearchQueryType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "public":
		fallthrough
	case "private":
		fallthrough
	case "secret":
		*e = ClubSearchQueryType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClubSearchQueryType: %v", v)
	}
}
