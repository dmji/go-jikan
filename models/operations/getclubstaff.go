package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetClubStaffRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetClubStaffRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetClubStaffResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns Club Staff
	ClubStaff *components.ClubStaff
}

func (o *GetClubStaffResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetClubStaffResponse) GetClubStaff() *components.ClubStaff {
	if o == nil {
		return nil
	}
	return o.ClubStaff
}
