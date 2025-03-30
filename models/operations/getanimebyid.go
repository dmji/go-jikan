package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetAnimeByIDResponseBody - Returns anime resource
type GetAnimeByIDResponseBody struct {
	// Anime Resource
	Data *components.Anime `json:"data,omitempty"`
}

func (o *GetAnimeByIDResponseBody) GetData() *components.Anime {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetAnimeByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime resource
	Object *GetAnimeByIDResponseBody
}

func (o *GetAnimeByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeByIDResponse) GetObject() *GetAnimeByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
