package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetClubRelationsRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetClubRelationsRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetClubRelationsResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns Club Relations
	ClubRelations *components.ClubRelations
}

func (o *GetClubRelationsResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetClubRelationsResponse) GetClubRelations() *components.ClubRelations {
	if o == nil {
		return nil
	}
	return o.ClubRelations
}
