package components

type ClubRelationsData struct {
	Anime      []MalURL `json:"anime,omitempty"`
	Manga      []MalURL `json:"manga,omitempty"`
	Characters []MalURL `json:"characters,omitempty"`
}

func (o *ClubRelationsData) GetAnime() []MalURL {
	if o == nil {
		return nil
	}
	return o.Anime
}

func (o *ClubRelationsData) GetManga() []MalURL {
	if o == nil {
		return nil
	}
	return o.Manga
}

func (o *ClubRelationsData) GetCharacters() []MalURL {
	if o == nil {
		return nil
	}
	return o.Characters
}

// ClubRelations - Club Relations
type ClubRelations struct {
	Data *ClubRelationsData `json:"data,omitempty"`
}

func (o *ClubRelations) GetData() *ClubRelationsData {
	if o == nil {
		return nil
	}
	return o.Data
}
