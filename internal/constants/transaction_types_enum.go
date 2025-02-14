// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package constants

import (
	"errors"
	"fmt"
)

const (
	// TransactionTypeINBOUND is a TransactionType of type INBOUND.
	TransactionTypeINBOUND TransactionType = "INBOUND"
	// TransactionTypeOUTBOUND is a TransactionType of type OUTBOUND.
	TransactionTypeOUTBOUND TransactionType = "OUTBOUND"
	// TransactionTypeOUTGOINGFEE is a TransactionType of type OUTGOING_FEE.
	TransactionTypeOUTGOINGFEE TransactionType = "OUTGOING_FEE"
)

var ErrInvalidTransactionType = errors.New("not a valid TransactionType")

// String implements the Stringer interface.
func (x TransactionType) String() string {
	return string(x)
}

// String implements the Stringer interface.
func (x TransactionType) IsValid() bool {
	_, err := ParseTransactionType(string(x))
	return err == nil
}

var _TransactionTypeValue = map[string]TransactionType{
	"INBOUND":      TransactionTypeINBOUND,
	"OUTBOUND":     TransactionTypeOUTBOUND,
	"OUTGOING_FEE": TransactionTypeOUTGOINGFEE,
}

// ParseTransactionType attempts to convert a string to a TransactionType.
func ParseTransactionType(name string) (TransactionType, error) {
	if x, ok := _TransactionTypeValue[name]; ok {
		return x, nil
	}
	return TransactionType(""), fmt.Errorf("%s is %w", name, ErrInvalidTransactionType)
}
