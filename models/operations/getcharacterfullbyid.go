package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetCharacterFullByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetCharacterFullByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetCharacterFullByIDResponseBody - Returns complete character resource data
type GetCharacterFullByIDResponseBody struct {
	// Character Resource
	Data *components.CharacterFull `json:"data,omitempty"`
}

func (o *GetCharacterFullByIDResponseBody) GetData() *components.CharacterFull {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetCharacterFullByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns complete character resource data
	Object *GetCharacterFullByIDResponseBody
}

func (o *GetCharacterFullByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetCharacterFullByIDResponse) GetObject() *GetCharacterFullByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
