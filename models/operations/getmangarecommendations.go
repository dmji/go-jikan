package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetMangaRecommendationsRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetMangaRecommendationsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetMangaRecommendationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns manga recommendations
	EntryRecommendations *components.EntryRecommendations
}

func (o *GetMangaRecommendationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetMangaRecommendationsResponse) GetEntryRecommendations() *components.EntryRecommendations {
	if o == nil {
		return nil
	}
	return o.EntryRecommendations
}
