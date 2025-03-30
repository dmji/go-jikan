package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

// GetRandomAnimeResponseBody - Returns a random anime resource
type GetRandomAnimeResponseBody struct {
	// Anime Resource
	Data *components.Anime `json:"data,omitempty"`
}

func (o *GetRandomAnimeResponseBody) GetData() *components.Anime {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetRandomAnimeResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a random anime resource
	Object *GetRandomAnimeResponseBody
}

func (o *GetRandomAnimeResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRandomAnimeResponse) GetObject() *GetRandomAnimeResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
