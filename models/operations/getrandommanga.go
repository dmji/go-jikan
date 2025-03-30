package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

// GetRandomMangaResponseBody - Returns a random manga resource
type GetRandomMangaResponseBody struct {
	// Manga Resource
	Data *components.Manga `json:"data,omitempty"`
}

func (o *GetRandomMangaResponseBody) GetData() *components.Manga {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetRandomMangaResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a random manga resource
	Object *GetRandomMangaResponseBody
}

func (o *GetRandomMangaResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRandomMangaResponse) GetObject() *GetRandomMangaResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
