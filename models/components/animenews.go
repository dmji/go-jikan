package components

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

type AnimeNewsData struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
	// Title
	Title *string `json:"title,omitempty"`
	// Post Date ISO8601
	Date *string `json:"date,omitempty"`
	// Author MyAnimeList Username
	AuthorUsername *string `json:"author_username,omitempty"`
	// Author Profile URL
	AuthorURL *string `json:"author_url,omitempty"`
	// Forum topic URL
	ForumURL *string       `json:"forum_url,omitempty"`
	Images   *CommonImages `json:"images,omitempty"`
	// Comment count
	Comments *int64 `json:"comments,omitempty"`
	// Excerpt
	Excerpt *string `json:"excerpt,omitempty"`
}

func (o *AnimeNewsData) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *AnimeNewsData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *AnimeNewsData) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *AnimeNewsData) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *AnimeNewsData) GetAuthorUsername() *string {
	if o == nil {
		return nil
	}
	return o.AuthorUsername
}

func (o *AnimeNewsData) GetAuthorURL() *string {
	if o == nil {
		return nil
	}
	return o.AuthorURL
}

func (o *AnimeNewsData) GetForumURL() *string {
	if o == nil {
		return nil
	}
	return o.ForumURL
}

func (o *AnimeNewsData) GetImages() *CommonImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *AnimeNewsData) GetComments() *int64 {
	if o == nil {
		return nil
	}
	return o.Comments
}

func (o *AnimeNewsData) GetExcerpt() *string {
	if o == nil {
		return nil
	}
	return o.Excerpt
}

// AnimeNews - Anime News Resource
type AnimeNews struct {
	Pagination *Pagination     `json:"pagination,omitempty"`
	Data       []AnimeNewsData `json:"data,omitempty"`
}

func (o *AnimeNews) GetPagination() *Pagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

func (o *AnimeNews) GetData() []AnimeNewsData {
	if o == nil {
		return nil
	}
	return o.Data
}
