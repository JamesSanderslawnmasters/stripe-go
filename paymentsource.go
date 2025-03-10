//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"fmt"
	"github.com/stripe/stripe-go/v72/form"
)

// PaymentSourceType consts represent valid payment sources.
type PaymentSourceType string

// List of values that PaymentSourceType can take.
const (
	PaymentSourceTypeAccount     PaymentSourceType = "account"
	PaymentSourceTypeBankAccount PaymentSourceType = "bank_account"
	PaymentSourceTypeCard        PaymentSourceType = "card"
	PaymentSourceTypeObject      PaymentSourceType = "source"
)

// SourceParams is a union struct used to describe an
// arbitrary payment source.
type SourceParams struct {
	Card  *CardParams `form:"-"`
	Token *string     `form:"source"`
}

// AppendTo implements custom encoding logic for SourceParams.
func (p *SourceParams) AppendTo(body *form.Values, keyParts []string) {
	if p.Card != nil {
		p.Card.AppendToAsCardSourceOrExternalAccount(body, keyParts)
	}
}

// CustomerSourceParams are used to manipulate a given Stripe
// Customer object's payment sources.
// For more details see https://stripe.com/docs/api#sources
type CustomerSourceParams struct {
	Params            `form:"*"`
	Customer          *string                   `form:"-"` // Included in URL
	AccountHolderName *string                   `form:"account_holder_name"`
	AccountHolderType *string                   `form:"account_holder_type"`
	AddressCity       *string                   `form:"address_city"`
	AddressCountry    *string                   `form:"address_country"`
	AddressLine1      *string                   `form:"address_line1"`
	AddressLine2      *string                   `form:"address_line2"`
	AddressState      *string                   `form:"address_state"`
	AddressZip        *string                   `form:"address_zip"`
	ExpMonth          *string                   `form:"exp_month"`
	ExpYear           *string                   `form:"exp_year"`
	Name              *string                   `form:"name"`
	Owner             *PaymentSourceOwnerParams `form:"owner"`
	Source            *SourceParams             `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
}
type PaymentSourceOwnerParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// SourceVerifyParams are used to verify a customer source
// For more details see https://stripe.com/docs/guides/ach-beta
type SourceVerifyParams struct {
	Params   `form:"*"`
	Customer *string   `form:"-"`       // Included in URL
	Amounts  [2]int64  `form:"amounts"` // Amounts is used when verifying bank accounts
	Values   []*string `form:"values"`  // Values is used when verifying sources
}

// SourceListParams are used to enumerate the payment sources that are attached
// to a Customer.
type SourceListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"-"` // Included in URL
	Object     *string `form:"object"`
}

// SetSource adds valid sources to a CustomerSourceParams object,
// returning an error for unsupported sources.
func (cp *CustomerSourceParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	cp.Source = source
	return err
}

// SourceParamsFor creates SourceParams objects around supported
// payment sources, returning errors if not.
//
// Currently supported source types are Card (CardParams) and
// Tokens/IDs (string), where Tokens could be single use card
// tokens
func SourceParamsFor(obj interface{}) (*SourceParams, error) {
	var sp *SourceParams
	var err error
	switch p := obj.(type) {
	case *CardParams:
		sp = &SourceParams{
			Card: p,
		}
	case string:
		sp = &SourceParams{
			Token: &p,
		}
	default:
		err = fmt.Errorf("Unsupported source type %s", p)
	}
	return sp, err
}

// PaymentSource describes the payment source used to make a Charge.
// The Type should indicate which object is fleshed out (eg. BankAccount or Card)
// For more details see https://stripe.com/docs/api#retrieve_charge
type PaymentSource struct {
	APIResource
	BankAccount  *BankAccount      `json:"-"`
	Card         *Card             `json:"-"`
	Deleted      bool              `json:"deleted"`
	ID           string            `json:"id"`
	SourceObject *Source           `json:"-"`
	Type         PaymentSourceType `json:"object"`
}

// SourceList is a list object for cards.
type SourceList struct {
	APIResource
	ListMeta
	Data []*PaymentSource `json:"data"`
}

// UnmarshalJSON handles deserialization of a PaymentSource.
// This custom unmarshaling is needed because the specific
// type of payment instrument it refers to is specified in the JSON
func (s *PaymentSource) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type paymentSource PaymentSource
	var v paymentSource
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var err error
	*s = PaymentSource(v)

	switch s.Type {
	case PaymentSourceTypeBankAccount:
		err = json.Unmarshal(data, &s.BankAccount)
	case PaymentSourceTypeCard:
		err = json.Unmarshal(data, &s.Card)
	case PaymentSourceTypeObject:
		err = json.Unmarshal(data, &s.SourceObject)
	}

	return err
}

// MarshalJSON handles serialization of a PaymentSource.
// This custom marshaling is needed because the specific type
// of payment instrument it represents is specified by the Type
func (s *PaymentSource) MarshalJSON() ([]byte, error) {
	var target interface{}

	switch s.Type {
	case PaymentSourceTypeCard:
		var customerID *string
		if s.Card.Customer != nil {
			customerID = &s.Card.Customer.ID
		}

		target = struct {
			*Card
			Customer *string           `json:"customer"`
			Type     PaymentSourceType `json:"object"`
		}{
			Card:     s.Card,
			Customer: customerID,
			Type:     s.Type,
		}
	case PaymentSourceTypeAccount:
		target = struct {
			ID   string            `json:"id"`
			Type PaymentSourceType `json:"object"`
		}{
			ID:   s.ID,
			Type: s.Type,
		}
	case PaymentSourceTypeBankAccount:
		var customerID *string
		if s.BankAccount.Customer != nil {
			customerID = &s.BankAccount.Customer.ID
		}

		target = struct {
			*BankAccount
			Customer *string           `json:"customer"`
			Type     PaymentSourceType `json:"object"`
		}{
			BankAccount: s.BankAccount,
			Customer:    customerID,
			Type:        s.Type,
		}
	case "":
		target = s.ID
	}

	return json.Marshal(target)
}
