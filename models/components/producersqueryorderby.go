package components

import (
	"encoding/json"
	"fmt"
)

// ProducersQueryOrderby - Producers Search Query Order By
type ProducersQueryOrderby string

const (
	ProducersQueryOrderbyMalID       ProducersQueryOrderby = "mal_id"
	ProducersQueryOrderbyCount       ProducersQueryOrderby = "count"
	ProducersQueryOrderbyFavorites   ProducersQueryOrderby = "favorites"
	ProducersQueryOrderbyEstablished ProducersQueryOrderby = "established"
)

func (e ProducersQueryOrderby) ToPointer() *ProducersQueryOrderby {
	return &e
}

func (e *ProducersQueryOrderby) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mal_id":
		fallthrough
	case "count":
		fallthrough
	case "favorites":
		fallthrough
	case "established":
		*e = ProducersQueryOrderby(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ProducersQueryOrderby: %v", v)
	}
}
