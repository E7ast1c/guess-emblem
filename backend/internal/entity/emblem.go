package entity

type Emblem struct {
	Id          int16  `json:"id,omitempty" db:"id"`
	Name        string `json:"name,omitempty" db:"name"`
	FullName    string `json:"full_name,omitempty" db:"full_name"`
	Title       string `json:"title,omitempty" db:"title"`
	Description string `json:"description,omitempty" db:"description"`
	Emblem      []byte `json:"emblem,omitempty" db:"emblem"`
	Link        string `json:"link,omitempty" db:"link"`
}
