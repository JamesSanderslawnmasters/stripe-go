//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// TransferSourceType is the list of allowed values for the transfer's source_type field.
type TransferSourceType string

// List of values that TransferSourceType can take.
const (
	TransferSourceTypeBankAccount TransferSourceType = "bank_account"
	TransferSourceTypeCard        TransferSourceType = "card"
	TransferSourceTypeFPX         TransferSourceType = "fpx"
)

// TransferParams is the set of parameters that can be used when creating or updating a transfer.
// For more details see https://stripe.com/docs/api#create_transfer and https://stripe.com/docs/api#update_transfer.
type TransferParams struct {
	Params            `form:"*"`
	Amount            *int64  `form:"amount"`
	Currency          *string `form:"currency"`
	Description       *string `form:"description"`
	Destination       *string `form:"destination"`
	SourceTransaction *string `form:"source_transaction"`
	SourceType        *string `form:"source_type"`
	TransferGroup     *string `form:"transfer_group"`
}

// TransferListParams is the set of parameters that can be used when listing transfers.
// For more details see https://stripe.com/docs/api#list_transfers.
type TransferListParams struct {
	ListParams    `form:"*"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Destination   *string           `form:"destination"`
	TransferGroup *string           `form:"transfer_group"`
}

// TransferDestination describes the destination of a Transfer.
// The Type should indicate which object is fleshed out
// For more details see https://stripe.com/docs/api/?lang=go#transfer_object
type TransferDestination struct {
	ID string `json:"id"`

	Account *Account `json:"-"`
}

// Transfer is the resource representing a Stripe transfer.
// For more details see https://stripe.com/docs/api#transfers.
type Transfer struct {
	APIResource
	Amount             int64                     `json:"amount"`
	AmountReversed     int64                     `json:"amount_reversed"`
	BalanceTransaction *BalanceTransaction       `json:"balance_transaction"`
	Created            int64                     `json:"created"`
	Currency           Currency                  `json:"currency"`
	Description        string                    `json:"description"`
	Destination        *TransferDestination      `json:"destination"`
	DestinationPayment *Charge                   `json:"destination_payment"`
	ID                 string                    `json:"id"`
	Livemode           bool                      `json:"livemode"`
	Metadata           map[string]string         `json:"metadata"`
	Object             string                    `json:"object"`
	Reversals          *ReversalList             `json:"reversals"`
	Reversed           bool                      `json:"reversed"`
	SourceTransaction  *BalanceTransactionSource `json:"source_transaction"`
	SourceType         TransferSourceType        `json:"source_type"`
	TransferGroup      string                    `json:"transfer_group"`
}

// TransferList is a list of transfers as retrieved from a list endpoint.
type TransferList struct {
	APIResource
	ListMeta
	Data []*Transfer `json:"data"`
}

// UnmarshalJSON handles deserialization of a Transfer.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *Transfer) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type transfer Transfer
	var v transfer
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = Transfer(v)
	return nil
}

// UnmarshalJSON handles deserialization of a TransferDestination.
// This custom unmarshaling is needed because the specific type of
// TransferDestination it refers to is specified in the JSON
func (t *TransferDestination) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type transferDestination TransferDestination
	var v transferDestination
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = TransferDestination(v)
	err := json.Unmarshal(data, &t.Account)
	return err
}
