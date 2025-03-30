package components

type Promo struct {
	// Title
	Title *string `json:"title,omitempty"`
	// Youtube Details
	Trailer *Trailer `json:"trailer,omitempty"`
}

func (o *Promo) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Promo) GetTrailer() *Trailer {
	if o == nil {
		return nil
	}
	return o.Trailer
}

type Episodes struct {
	// MyAnimeList ID
	MalID *int64 `json:"mal_id,omitempty"`
	// MyAnimeList URL
	URL *string `json:"url,omitempty"`
	// Title
	Title *string `json:"title,omitempty"`
	// Episode
	Episode *string       `json:"episode,omitempty"`
	Images  *CommonImages `json:"images,omitempty"`
}

func (o *Episodes) GetMalID() *int64 {
	if o == nil {
		return nil
	}
	return o.MalID
}

func (o *Episodes) GetURL() *string {
	if o == nil {
		return nil
	}
	return o.URL
}

func (o *Episodes) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Episodes) GetEpisode() *string {
	if o == nil {
		return nil
	}
	return o.Episode
}

func (o *Episodes) GetImages() *CommonImages {
	if o == nil {
		return nil
	}
	return o.Images
}

type Meta struct {
	Title  *string `json:"title,omitempty"`
	Author *string `json:"author,omitempty"`
}

func (o *Meta) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *Meta) GetAuthor() *string {
	if o == nil {
		return nil
	}
	return o.Author
}

type MusicVideos struct {
	// Title
	Title *string `json:"title,omitempty"`
	// Youtube Details
	Video *Trailer `json:"video,omitempty"`
	Meta  *Meta    `json:"meta,omitempty"`
}

func (o *MusicVideos) GetTitle() *string {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *MusicVideos) GetVideo() *Trailer {
	if o == nil {
		return nil
	}
	return o.Video
}

func (o *MusicVideos) GetMeta() *Meta {
	if o == nil {
		return nil
	}
	return o.Meta
}

type AnimeVideosData struct {
	Promo       []Promo       `json:"promo,omitempty"`
	Episodes    []Episodes    `json:"episodes,omitempty"`
	MusicVideos []MusicVideos `json:"music_videos,omitempty"`
}

func (o *AnimeVideosData) GetPromo() []Promo {
	if o == nil {
		return nil
	}
	return o.Promo
}

func (o *AnimeVideosData) GetEpisodes() []Episodes {
	if o == nil {
		return nil
	}
	return o.Episodes
}

func (o *AnimeVideosData) GetMusicVideos() []MusicVideos {
	if o == nil {
		return nil
	}
	return o.MusicVideos
}

// AnimeVideos - Anime Videos Resource
type AnimeVideos struct {
	Data *AnimeVideosData `json:"data,omitempty"`
}

func (o *AnimeVideos) GetData() *AnimeVideosData {
	if o == nil {
		return nil
	}
	return o.Data
}
