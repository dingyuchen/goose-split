package handlers

type Proportion struct {
	User   User
	Amount uint64
}

type Transaction struct {
	PaidBy      User
	Amount      uint64
	Proportions []Proportion
}

type TransactionService interface {
	AddTransaction(id string, transaction Transaction) error
	RemoveTransaction(id string, transaction Transaction) error
	UpdateTransaction(id string, transaction Transaction) error
	GetTransactionByUser(email string) (Transaction, error)
	GetTransactionByParty(id string) (Transaction, error)
}
