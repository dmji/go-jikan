package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetMangaStatisticsRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMangaStatisticsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetMangaStatisticsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime statistics
	MangaStatistics *components.MangaStatistics
}

func (o *GetMangaStatisticsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaStatisticsResponse) GetMangaStatistics() *components.MangaStatistics {
	if o == nil {
		return nil
	}
	return o.MangaStatistics
}
