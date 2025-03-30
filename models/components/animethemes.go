package components

type AnimeThemesData struct {
	Openings []string `json:"openings,omitempty"`
	Endings  []string `json:"endings,omitempty"`
}

func (o *AnimeThemesData) GetOpenings() []string {
	if o == nil {
		return nil
	}
	return o.Openings
}

func (o *AnimeThemesData) GetEndings() []string {
	if o == nil {
		return nil
	}
	return o.Endings
}

// AnimeThemes - Anime Opening and Ending Themes
type AnimeThemes struct {
	Data *AnimeThemesData `json:"data,omitempty"`
}

func (o *AnimeThemes) GetData() *AnimeThemesData {
	if o == nil {
		return nil
	}
	return o.Data
}
