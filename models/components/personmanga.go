package components

type PersonMangaData struct {
	// Person's position
	Position *string    `json:"position,omitempty"`
	Manga    *MangaMeta `json:"manga,omitempty"`
}

func (o *PersonMangaData) GetPosition() *string {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *PersonMangaData) GetManga() *MangaMeta {
	if o == nil {
		return nil
	}
	return o.Manga
}

// PersonManga - Person's mangaography
type PersonManga struct {
	Data []PersonMangaData `json:"data,omitempty"`
}

func (o *PersonManga) GetData() []PersonMangaData {
	if o == nil {
		return nil
	}
	return o.Data
}
