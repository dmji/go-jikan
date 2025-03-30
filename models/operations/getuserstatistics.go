package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetUserStatisticsRequest struct {
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

func (o *GetUserStatisticsRequest) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type GetUserStatisticsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns user statistics
	UserStatistics *components.UserStatistics
}

func (o *GetUserStatisticsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUserStatisticsResponse) GetUserStatistics() *components.UserStatistics {
	if o == nil {
		return nil
	}
	return o.UserStatistics
}
