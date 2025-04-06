package transaction

type Transaction struct {
	Transaction_id  string  `json:"transaction_id"`
	SourceAccountID string  `json:"source_account_id"`
	TargetAccountID string  `json:"target_account_id"`
	Amount          float64 `json:"amount"`
}
