package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetWatchPopularEpisodesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns Popular Episodes
	WatchEpisodes *components.WatchEpisodes
}

func (o *GetWatchPopularEpisodesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetWatchPopularEpisodesResponse) GetWatchEpisodes() *components.WatchEpisodes {
	if o == nil {
		return nil
	}
	return o.WatchEpisodes
}
