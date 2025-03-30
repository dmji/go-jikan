package components

type MangaNewsPagination struct {
	LastVisiblePage *int64 `json:"last_visible_page,omitempty"`
	HasNextPage     *bool  `json:"has_next_page,omitempty"`
}

func (o *MangaNewsPagination) GetLastVisiblePage() *int64 {
	if o == nil {
		return nil
	}
	return o.LastVisiblePage
}

func (o *MangaNewsPagination) GetHasNextPage() *bool {
	if o == nil {
		return nil
	}
	return o.HasNextPage
}

type MangaNewsData struct {
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

func (o *MangaNewsData) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *MangaNewsData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *MangaNewsData) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *MangaNewsData) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *MangaNewsData) GetAuthorUsername() *string {
	if o == nil {
		return nil
	}
	return o.AuthorUsername
}

func (o *MangaNewsData) GetAuthorURL() *string {
	if o == nil {
		return nil
	}
	return o.AuthorURL
}

func (o *MangaNewsData) GetForumURL() *string {
	if o == nil {
		return nil
	}
	return o.ForumURL
}

func (o *MangaNewsData) GetImages() *CommonImages {
	if o == nil {
		return nil
	}
	return o.Images
}

func (o *MangaNewsData) GetComments() *int64 {
	if o == nil {
		return nil
	}
	return o.Comments
}

func (o *MangaNewsData) GetExcerpt() *string {
	if o == nil {
		return nil
	}
	return o.Excerpt
}

// MangaNews - Manga News Resource
type MangaNews struct {
	Pagination *MangaNewsPagination `json:"pagination,omitempty"`
	Data       []MangaNewsData      `json:"data,omitempty"`
}

func (o *MangaNews) GetPagination() *MangaNewsPagination {
	if o == nil {
		return nil
	}
	return o.Pagination
}

func (o *MangaNews) GetData() []MangaNewsData {
	if o == nil {
		return nil
	}
	return o.Data
}
