package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetAnimeRecommendationsRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetAnimeRecommendationsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetAnimeRecommendationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns anime recommendations
	EntryRecommendations *components.EntryRecommendations
}

func (o *GetAnimeRecommendationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetAnimeRecommendationsResponse) GetEntryRecommendations() *components.EntryRecommendations {
	if o == nil {
		return nil
	}
	return o.EntryRecommendations
}
