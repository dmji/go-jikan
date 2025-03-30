package components

// LastComment - Last comment details
type LastComment struct {
	// Last comment URL
	URL *string `json:"url,omitempty"`
	// Author MyAnimeList Username
	AuthorUsername *string `json:"author_username,omitempty"`
	// Author Profile URL
	AuthorURL *string `json:"author_url,omitempty"`
	// Last comment date posted ISO8601
	Date *string `json:"date,omitempty"`
}

func (o *LastComment) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *LastComment) GetAuthorUsername() *string {
	if o == nil {
		return nil
	}
	return o.AuthorUsername
}

func (o *LastComment) GetAuthorURL() *string {
	if o == nil {
		return nil
	}
	return o.AuthorURL
}

func (o *LastComment) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

type ForumData struct {
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
	// Comment count
	Comments *int64 `json:"comments,omitempty"`
	// Last comment details
	LastComment *LastComment `json:"last_comment,omitempty"`
}

func (o *ForumData) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *ForumData) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *ForumData) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *ForumData) GetDate() *string {
	if o == nil {
		return nil
	}
	return o.Date
}

func (o *ForumData) GetAuthorUsername() *string {
	if o == nil {
		return nil
	}
	return o.AuthorUsername
}

func (o *ForumData) GetAuthorURL() *string {
	if o == nil {
		return nil
	}
	return o.AuthorURL
}

func (o *ForumData) GetComments() *int64 {
	if o == nil {
		return nil
	}
	return o.Comments
}

func (o *ForumData) GetLastComment() *LastComment {
	if o == nil {
		return nil
	}
	return o.LastComment
}

// Forum Resource
type Forum struct {
	Data []ForumData `json:"data,omitempty"`
}

func (o *Forum) GetData() []ForumData {
	if o == nil {
		return nil
	}
	return o.Data
}
