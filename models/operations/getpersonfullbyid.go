package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetPersonFullByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetPersonFullByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetPersonFullByIDResponseBody - Returns complete character resource data
type GetPersonFullByIDResponseBody struct {
	// Person Resource
	Data *components.PersonFull `json:"data,omitempty"`
}

func (o *GetPersonFullByIDResponseBody) GetData() *components.PersonFull {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetPersonFullByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns complete character resource data
	Object *GetPersonFullByIDResponseBody
}

func (o *GetPersonFullByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPersonFullByIDResponse) GetObject() *GetPersonFullByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
