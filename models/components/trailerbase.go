package components

// TrailerBase - Youtube Details
type TrailerBase struct {
	// YouTube ID
	YoutubeID *string `json:"youtube_id,omitempty"`
	// YouTube URL
	URL *string `json:"url,omitempty"`
	// Parsed Embed URL
	EmbedURL *string `json:"embed_url,omitempty"`
}

func (o *TrailerBase) GetYoutubeID() *string {
	if o == nil {
		return nil
	}
	return o.YoutubeID
}

func (o *TrailerBase) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *TrailerBase) GetEmbedURL() *string {
	if o == nil {
		return nil
	}
	return o.EmbedURL
}
