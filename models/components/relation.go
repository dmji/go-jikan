package components

// Relation - Related resources
type Relation struct {
	// Relation type
	Relation *string `json:"relation,omitempty"`
	// Related entries
	Entry []MalURL `json:"entry,omitempty"`
}

func (o *Relation) GetRelation() *string {
	if o == nil {
		return nil
	}
	return o.Relation
}

func (o *Relation) GetEntry() []MalURL {
	if o == nil {
		return nil
	}
	return o.Entry
}
