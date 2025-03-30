package components

type UserFriendsData struct {
	User *UserMeta `json:"user,omitempty"`
	// Last Online Date ISO8601 format
	LastOnline *string `json:"last_online,omitempty"`
	// Friends Since Date ISO8601 format
	FriendsSince *string `json:"friends_since,omitempty"`
}

func (o *UserFriendsData) GetUser() *UserMeta {
	if o == nil {
		return nil
	}
	return o.User
}

func (o *UserFriendsData) GetLastOnline() *string {
	if o == nil {
		return nil
	}
	return o.LastOnline
}

func (o *UserFriendsData) GetFriendsSince() *string {
	if o == nil {
		return nil
	}
	return o.FriendsSince
}

type UserFriendsPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *UserFriendsPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *UserFriendsPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

// UserFriends - User Friends
type UserFriends struct {
	Data       []UserFriendsData      `json:"data,omitempty"`
	Pagination *UserFriendsPagination `json:"pagination,omitempty"`
}

func (o *UserFriends) GetData() []UserFriendsData {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *UserFriends) GetPagination() *UserFriendsPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}
