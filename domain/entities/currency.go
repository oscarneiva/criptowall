package entities

type Currency struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Value float64 `json:"value"`
}
