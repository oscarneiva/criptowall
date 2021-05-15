package entities

type User struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Wallets []Wallet `json:"wallets"`
}
