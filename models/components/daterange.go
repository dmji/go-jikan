package components

// From - Date Prop From
type From struct {
	// Day
	Day *int64 `json:"day,omitempty"`
	// Month
	Month *int64 `json:"month,omitempty"`
	// Year
	Year *int64 `json:"year,omitempty"`
}

func (o *From) GetDay() *int64 {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *From) GetMonth() *int64 {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *From) GetYear() *int64 {
	if o == nil {
		return nil
	}
	return o.Year
}

// To - Date Prop To
type To struct {
	// Day
	Day *int64 `json:"day,omitempty"`
	// Month
	Month *int64 `json:"month,omitempty"`
	// Year
	Year *int64 `json:"year,omitempty"`
}

func (o *To) GetDay() *int64 {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *To) GetMonth() *int64 {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *To) GetYear() *int64 {
	if o == nil {
		return nil
	}
	return o.Year
}

// Date Prop
type Prop struct {
	// Date Prop From
	From *From `json:"from,omitempty"`
	// Date Prop To
	To *To `json:"to,omitempty"`
	// Raw parsed string
	String *string `json:"string,omitempty"`
}

func (o *Prop) GetFrom() *From {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *Prop) GetTo() *To {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *Prop) GetString() *string {
	if o == nil {
		return nil
	}
	return o.String
}

// Daterange - Date range
type Daterange struct {
	// Date ISO8601
	From *string `json:"from,omitempty"`
	// Date ISO8601
	To *string `json:"to,omitempty"`
	// Date Prop
	Prop *Prop `json:"prop,omitempty"`
}

func (o *Daterange) GetFrom() *string {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *Daterange) GetTo() *string {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *Daterange) GetProp() *Prop {
	if o == nil {
		return nil
	}
	return o.Prop
}
