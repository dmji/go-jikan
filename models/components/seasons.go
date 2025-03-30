package components

type SeasonsData struct {
	// Year
	Year *int64 `json:"year,omitempty"`
	// List of available seasons
	Seasons []string `json:"seasons,omitempty"`
}

func (o *SeasonsData) GetYear() *int64 {
	if o == nil {
		return nil
	}
	return o.Year
}

func (o *SeasonsData) GetSeasons() []string {
	if o == nil {
		return nil
	}
	return o.Seasons
}

// Seasons - List of available seasons
type Seasons struct {
	Data []SeasonsData `json:"data,omitempty"`
}

func (o *Seasons) GetData() []SeasonsData {
	if o == nil {
		return nil
	}
	return o.Data
}
