package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeStatisticsRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeStatisticsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetAnimeStatisticsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime statistics
	AnimeStatistics *components.AnimeStatistics
}

func (o *GetAnimeStatisticsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeStatisticsResponse) GetAnimeStatistics() *components.AnimeStatistics {
	if o == nil {
		return nil
	}
	return o.AnimeStatistics
}
