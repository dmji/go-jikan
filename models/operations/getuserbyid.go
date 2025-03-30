package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetUserByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetUserByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetUserByIDResponseBody - Returns username by ID search
type GetUserByIDResponseBody struct {
	// User Meta By ID
	Data *components.UserByID `json:"data,omitempty"`
}

func (o *GetUserByIDResponseBody) GetData() *components.UserByID {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetUserByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns username by ID search
	Object *GetUserByIDResponseBody
}

func (o *GetUserByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserByIDResponse) GetObject() *GetUserByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
