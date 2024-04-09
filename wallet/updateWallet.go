package wallet

type UpdateWallet struct {
	ID         int     `json:"id" example:"1"`
	WalletName string  `json:"wallet_name" example:"John's Wallet"`
	WalletType string  `json:"wallet_type" example:"Create Card"`
	Balance    float64 `json:"balance" example:"100.00"`
}
