package operations

import (
	"encoding/json"
	"fmt"

	"github.com/dmji/go-jikan/models/components"
)

type Type string

const (
	TypeAnime Type = "anime"
	TypeManga Type = "manga"
)

func (e Type) ToPointer() *Type {
	return &e
}

func (e *Type) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "anime":
		fallthrough
	case "manga":
		*e = Type(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Type: %v", v)
	}
}

type GetUserHistoryRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
	Type     *Type  `queryParam:"style=form,explode=true,name=type"`
}

func (o *GetUserHistoryRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *GetUserHistoryRequest) GetType() *Type {
	if o == nil {
		return nil
	}
	return o.Type
}

type GetUserHistoryResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user history (past 30 days)
	UserHistory *components.UserHistory
}

func (o *GetUserHistoryResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserHistoryResponse) GetUserHistory() *components.UserHistory {
	if o == nil {
		return nil
	}
	return o.UserHistory
}
