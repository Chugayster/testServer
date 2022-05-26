package event

type Event struct {
	Id   uint	`json:"id" db:"id,omitempty"`
	Name string `json:"name" db:"name"`
	Location string `json:"location" db:"location"`
}
