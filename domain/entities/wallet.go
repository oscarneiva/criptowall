package entities

type Wallet struct {
	ID string `json:"id"`
	Currencies []Currency `json:"currencies"`
}