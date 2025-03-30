package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetAnimeVideosEpisodesRequest struct {
	ID   int64  `pathParam:"style=simple,explode=false,name=id"`
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetAnimeVideosEpisodesRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetAnimeVideosEpisodesRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetAnimeVideosEpisodesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns episode videos related to the entry
	AnimeVideosEpisodes *components.AnimeVideosEpisodes
}

func (o *GetAnimeVideosEpisodesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeVideosEpisodesResponse) GetAnimeVideosEpisodes() *components.AnimeVideosEpisodes {
	if o == nil {
		return nil
	}
	return o.AnimeVideosEpisodes
}
