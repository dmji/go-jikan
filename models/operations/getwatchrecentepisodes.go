package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetWatchRecentEpisodesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns Recently Added Episodes
	WatchEpisodes *components.WatchEpisodes
}

func (o *GetWatchRecentEpisodesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetWatchRecentEpisodesResponse) GetWatchEpisodes() *components.WatchEpisodes {
	if o == nil {
		return nil
	}
	return o.WatchEpisodes
}
