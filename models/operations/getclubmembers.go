package operations

import (
	"github.com/dmji/jikan/models/components"
)

type GetClubMembersRequest struct {
	ID   int64  `pathParam:"style=simple,explode=false,name=id"`
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
}

func (o *GetClubMembersRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

func (o *GetClubMembersRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

type Pagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *Pagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *Pagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

type Data struct {
	// User's username
	Username *string `json:"username,omitempty"`
	// User URL
	URL    *string                `json:"url,omitempty"`
	Images *components.UserImages `json:"images,omitempty"`
}

func (o *Data) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *Data) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Data) GetImages() *components.UserImages {
	if o == nil {
		return nil
	}
	return o.Images
}

// GetClubMembersResponseBody - Club Member
type GetClubMembersResponseBody struct {
	Pagination *Pagination `json:"pagination,omitempty"`
	Data       []Data      `json:"data,omitempty"`
}

func (o *GetClubMembersResponseBody) GetPagination() *Pagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

func (o *GetClubMembersResponseBody) GetData() []Data {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetClubMembersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Returns Club Members Resource
	Object *GetClubMembersResponseBody
}

func (o *GetClubMembersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetClubMembersResponse) GetObject() *GetClubMembersResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
