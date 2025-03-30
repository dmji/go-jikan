package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeEpisodeByIDRequest struct {
	ID      int64 `pathParam:"style=simple,explode=false,name=id"`
	Episode int64 `pathParam:"style=simple,explode=false,name=episode"`
}

func (o *GetAnimeEpisodeByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetAnimeEpisodeByIDRequest) GetEpisode() int64 {
	if o == nil {
		return 0
	}
	return o.Episode
}

// GetAnimeEpisodeByIDResponseBody - Returns a single anime episode resource
type GetAnimeEpisodeByIDResponseBody struct {
	// Anime Episode Resource
	Data *components.AnimeEpisode `json:"data,omitempty"`
}

func (o *GetAnimeEpisodeByIDResponseBody) GetData() *components.AnimeEpisode {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetAnimeEpisodeByIDResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns a single anime episode resource
	Object *GetAnimeEpisodeByIDResponseBody
}

func (o *GetAnimeEpisodeByIDResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeEpisodeByIDResponse) GetObject() *GetAnimeEpisodeByIDResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
