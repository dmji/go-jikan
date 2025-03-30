package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetRecentAnimeRecommendationsRequest struct {
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetRecentAnimeRecommendationsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetRecentAnimeRecommendationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns recent anime recommendations
	Recommendations *components.Recommendations
}

func (o *GetRecentAnimeRecommendationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRecentAnimeRecommendationsResponse) GetRecommendations() *components.Recommendations {
	if o == nil {
		return nil
	}
	return o.Recommendations
}
