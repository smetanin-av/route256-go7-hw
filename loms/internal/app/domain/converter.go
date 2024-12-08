package domain

import (
	"fmt"
)

func (s Status) String() string {
	switch s {
	case StatusInvalid:
		return "Invalid"
	case StatusNew:
		return "New"
	case StatusAwaitingPayment:
		return "AwaitingPayment"
	case StatusFailed:
		return "Failed"
	case StatusPayed:
		return "Payed"
	case StatusCancelled:
		return "Cancelled"
	}
	return fmt.Sprintf("%d", s)
}

func (s Status) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", s)), nil
}
