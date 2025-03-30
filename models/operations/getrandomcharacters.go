package operations

import (
	"github.com/dmji/jikan/models/components"
)

// GetRandomCharactersResponseBody - Returns a random character resource
type GetRandomCharactersResponseBody struct {
	// Character Resource
	Data *components.Character `json:"data,omitempty"`
}

func (o *GetRandomCharactersResponseBody) GetData() *components.Character {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetRandomCharactersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a random character resource
	Object *GetRandomCharactersResponseBody
}

func (o *GetRandomCharactersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRandomCharactersResponse) GetObject() *GetRandomCharactersResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
