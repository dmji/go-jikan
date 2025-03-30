package operations

import (
	"github.com/dmji/go-jikan/models/components"
)

type GetUsersSearchRequest struct {
	Page  *int64  `queryParam:"style=form,explode=true,name=page"`
	Limit *int64  `queryParam:"style=form,explode=true,name=limit"`
	Q     *string `queryParam:"style=form,explode=true,name=q"`
	// Users Search Query Gender.
	Gender   *components.UsersSearchQueryGender `queryParam:"style=form,explode=true,name=gender"`
	Location *string                            `queryParam:"style=form,explode=true,name=location"`
	MaxAge   *int64                             `queryParam:"style=form,explode=true,name=maxAge"`
	MinAge   *int64                             `queryParam:"style=form,explode=true,name=minAge"`
}

func (o *GetUsersSearchRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetUsersSearchRequest) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *GetUsersSearchRequest) GetQ() *string {
	if o == nil {
		return nil
	}
	return o.Q
}

func (o *GetUsersSearchRequest) GetGender() *components.UsersSearchQueryGender {
	if o == nil {
		return nil
	}
	return o.Gender
}

func (o *GetUsersSearchRequest) GetLocation() *string {
	if o == nil {
		return nil
	}
	return o.Location
}

func (o *GetUsersSearchRequest) GetMaxAge() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxAge
}

func (o *GetUsersSearchRequest) GetMinAge() *int64 {
	if o == nil {
		return nil
	}
	return o.MinAge
}

type GetUsersSearchResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns search results for users
	UsersSearch *components.UsersSearch
}

func (o *GetUsersSearchResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetUsersSearchResponse) GetUsersSearch() *components.UsersSearch {
	if o == nil {
		return nil
	}
	return o.UsersSearch
}
