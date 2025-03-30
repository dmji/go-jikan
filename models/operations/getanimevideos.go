package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetAnimeVideosRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeVideosRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetAnimeVideosResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns videos related to the entry
	AnimeVideos *components.AnimeVideos
}

func (o *GetAnimeVideosResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeVideosResponse) GetAnimeVideos() *components.AnimeVideos {
	if o == nil {
		return nil
	}
	return o.AnimeVideos
}
