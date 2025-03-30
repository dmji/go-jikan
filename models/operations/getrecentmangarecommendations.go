package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetRecentMangaRecommendationsRequest struct {
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetRecentMangaRecommendationsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetRecentMangaRecommendationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns recent manga recommendations
	Recommendations *components.Recommendations
}

func (o *GetRecentMangaRecommendationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetRecentMangaRecommendationsResponse) GetRecommendations() *components.Recommendations {
	if o == nil {
		return nil
	}
	return o.Recommendations
}
