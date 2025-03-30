package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetSeasonsListResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns available list of seasons
	Seasons *components.Seasons
}

func (o *GetSeasonsListResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetSeasonsListResponse) GetSeasons() *components.Seasons {
	if o == nil {
		return nil
	}
	return o.Seasons
}
