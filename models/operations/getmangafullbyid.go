package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetMangaFullByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMangaFullByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

// GetMangaFullByIDResponseBody - Returns complete manga resource data
type GetMangaFullByIDResponseBody struct {
	// Manga Resource
	Data *components.MangaFull `json:"data,omitempty"`
}

func (o *GetMangaFullByIDResponseBody) GetData() *components.MangaFull {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetMangaFullByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns complete manga resource data
	Object *GetMangaFullByIDResponseBody
}

func (o *GetMangaFullByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaFullByIDResponse) GetObject() *GetMangaFullByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
