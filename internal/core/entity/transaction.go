package entity

type Transaction struct {
	ID                int64
	Customer          *Customer
	FinalizedCustomer *Customer
}
