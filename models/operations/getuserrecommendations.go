package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetUserRecommendationsRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
	Page     *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetUserRecommendationsRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

func (o *GetUserRecommendationsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type GetUserRecommendationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns Recent Anime Recommendations
	Recommendations *components.Recommendations
}

func (o *GetUserRecommendationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserRecommendationsResponse) GetRecommendations() *components.Recommendations {
	if o == nil {
		return nil
	}
	return o.Recommendations
}
