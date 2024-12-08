package domain

const (
	StatusInvalid Status = iota
	StatusNew
	StatusAwaitingPayment
	StatusFailed
	StatusPayed
	StatusCancelled
)
