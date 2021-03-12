package app

type Transction struct {
	ID int
	AccountID int
	Amount float64 // transaction amount, f.e. if we credit 100 coins to account, Amount = 100
	AccountAmount float64 // account amount after transaction, AccountAmount = Current AccountAmount + Amount
}