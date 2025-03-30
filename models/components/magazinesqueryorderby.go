package components

import (
	"encoding/json"
	"fmt"
)

// MagazinesQueryOrderby - Order by magazine data
type MagazinesQueryOrderby string

const (
	MagazinesQueryOrderbyMalID MagazinesQueryOrderby = "mal_id"
	MagazinesQueryOrderbyName  MagazinesQueryOrderby = "name"
	MagazinesQueryOrderbyCount MagazinesQueryOrderby = "count"
)

func (e MagazinesQueryOrderby) ToPointer() *MagazinesQueryOrderby {
	return &e
}

func (e *MagazinesQueryOrderby) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "mal_id":
		fallthrough
	case "name":
		fallthrough
	case "count":
		*e = MagazinesQueryOrderby(v)
		return nil
	default:
		return fmt.Errorf("invalid value for MagazinesQueryOrderby: %v", v)
	}
}
