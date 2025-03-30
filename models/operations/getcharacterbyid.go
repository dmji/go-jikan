package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetCharacterByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCharacterByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetCharacterByIDResponseBody - Returns character resource
type GetCharacterByIDResponseBody struct {
	// Character Resource
	Data *components.Character `json:"data,omitempty"`
}

func (o *GetCharacterByIDResponseBody) GetData() *components.Character {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetCharacterByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns character resource
	Object *GetCharacterByIDResponseBody
}

func (o *GetCharacterByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCharacterByIDResponse) GetObject() *GetCharacterByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
