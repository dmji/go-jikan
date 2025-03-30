package components

type PersonAnimeData struct {
	// Person's position
	Position *string    `json:"position,omitempty"`
	Anime    *AnimeMeta `json:"anime,omitempty"`
}

func (o *PersonAnimeData) GetPosition() *string {
	if o == nil {
		return nil
	}
	return o.Position
}

func (o *PersonAnimeData) GetAnime() *AnimeMeta {
	if o == nil {
		return nil
	}
	return o.Anime
}

// PersonAnime - Person anime staff positions
type PersonAnime struct {
	Data []PersonAnimeData `json:"data,omitempty"`
}

func (o *PersonAnime) GetData() []PersonAnimeData {
	if o == nil {
		return nil
	}
	return o.Data
}
