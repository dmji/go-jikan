package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetPersonByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetPersonByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetPersonByIDResponseBody - Returns pictures related to the entry
type GetPersonByIDResponseBody struct {
	// Person Resource
	Data *components.Person `json:"data,omitempty"`
}

func (o *GetPersonByIDResponseBody) GetData() *components.Person {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetPersonByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns pictures related to the entry
	Object *GetPersonByIDResponseBody
}

func (o *GetPersonByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetPersonByIDResponse) GetObject() *GetPersonByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
