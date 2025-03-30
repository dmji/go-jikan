package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeFullByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeFullByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetAnimeFullByIDResponseBody - Returns complete anime resource data
type GetAnimeFullByIDResponseBody struct {
	// Full anime Resource
	Data *components.AnimeFull `json:"data,omitempty"`
}

func (o *GetAnimeFullByIDResponseBody) GetData() *components.AnimeFull {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetAnimeFullByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns complete anime resource data
	Object *GetAnimeFullByIDResponseBody
}

func (o *GetAnimeFullByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeFullByIDResponse) GetObject() *GetAnimeFullByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
