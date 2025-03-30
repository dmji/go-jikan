package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetMangaByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMangaByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetMangaByIDResponseBody - Returns pictures related to the entry
type GetMangaByIDResponseBody struct {
	// Manga Resource
	Data *components.Manga `json:"data,omitempty"`
}

func (o *GetMangaByIDResponseBody) GetData() *components.Manga {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetMangaByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns pictures related to the entry
	Object *GetMangaByIDResponseBody
}

func (o *GetMangaByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaByIDResponse) GetObject() *GetMangaByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
