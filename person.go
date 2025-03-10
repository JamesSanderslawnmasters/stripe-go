//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// VerificationDocumentDetailsCode is a machine-readable code specifying the verification state of
// a document associated with a person.
type VerificationDocumentDetailsCode string

// List of values that IdentityVerificationDetailsCode can take.
const (
	VerificationDocumentDetailsCodeDocumentCorrupt               VerificationDocumentDetailsCode = "document_corrupt"
	VerificationDocumentDetailsCodeDocumentFailedCopy            VerificationDocumentDetailsCode = "document_failed_copy"
	VerificationDocumentDetailsCodeDocumentFailedGreyscale       VerificationDocumentDetailsCode = "document_failed_greyscale"
	VerificationDocumentDetailsCodeDocumentFailedOther           VerificationDocumentDetailsCode = "document_failed_other"
	VerificationDocumentDetailsCodeDocumentFailedTestMode        VerificationDocumentDetailsCode = "document_failed_test_mode"
	VerificationDocumentDetailsCodeDocumentFraudulent            VerificationDocumentDetailsCode = "document_fraudulent"
	VerificationDocumentDetailsCodeDocumentIDTypeNotSupported    VerificationDocumentDetailsCode = "document_id_type_not_supported"
	VerificationDocumentDetailsCodeDocumentIDCountryNotSupported VerificationDocumentDetailsCode = "document_id_country_not_supported"
	VerificationDocumentDetailsCodeDocumentManipulated           VerificationDocumentDetailsCode = "document_manipulated"
	VerificationDocumentDetailsCodeDocumentMissingBack           VerificationDocumentDetailsCode = "document_missing_back"
	VerificationDocumentDetailsCodeDocumentMissingFront          VerificationDocumentDetailsCode = "document_missing_front"
	VerificationDocumentDetailsCodeDocumentNotReadable           VerificationDocumentDetailsCode = "document_not_readable"
	VerificationDocumentDetailsCodeDocumentNotUploaded           VerificationDocumentDetailsCode = "document_not_uploaded"
	VerificationDocumentDetailsCodeDocumentTooLarge              VerificationDocumentDetailsCode = "document_too_large"
)

// PersonPoliticalExposure describes the political exposure of a given person.
type PersonPoliticalExposure string

// List of values that IdentityVerificationStatus can take.
const (
	PersonPoliticalExposureExisting PersonPoliticalExposure = "existing"
	PersonPoliticalExposureNone     PersonPoliticalExposure = "none"
)

// PersonVerificationDetailsCode is a machine-readable code specifying the verification state of a
// person.
type PersonVerificationDetailsCode string

// List of values that IdentityVerificationDetailsCode can take.
const (
	PersonVerificationDetailsCodeFailedKeyedIdentity PersonVerificationDetailsCode = "failed_keyed_identity"
	PersonVerificationDetailsCodeFailedOther         PersonVerificationDetailsCode = "failed_other"
	PersonVerificationDetailsCodeScanNameMismatch    PersonVerificationDetailsCode = "scan_name_mismatch"
)

// IdentityVerificationStatus describes the different statuses for identity verification.
type IdentityVerificationStatus string

// List of values that IdentityVerificationStatus can take.
const (
	IdentityVerificationStatusPending    IdentityVerificationStatus = "pending"
	IdentityVerificationStatusUnverified IdentityVerificationStatus = "unverified"
	IdentityVerificationStatusVerified   IdentityVerificationStatus = "verified"
)

// DOBParams represents a DOB during account creation/updates.
type DOBParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}

// One or more documents that demonstrate proof that this person is authorized to represent the company.
type DocumentsCompanyAuthorizationParams struct {
	Files []*string `form:"files"`
}

// One or more documents showing the person's passport page with photo and personal data.
type DocumentsPassportParams struct {
	Files []*string `form:"files"`
}

// One or more documents showing the person's visa required for living in the country where they are residing.
type DocumentsVisaParams struct {
	Files []*string `form:"files"`
}

// Documents that may be submitted to satisfy various informational requests.
type DocumentsParams struct {
	CompanyAuthorization *DocumentsCompanyAuthorizationParams `form:"company_authorization"`
	Passport             *DocumentsPassportParams             `form:"passport"`
	Visa                 *DocumentsVisaParams                 `form:"visa"`
}

// RelationshipParams is used to set the relationship between an account and a person.
type RelationshipParams struct {
	Director         *bool    `form:"director"`
	Executive        *bool    `form:"executive"`
	Owner            *bool    `form:"owner"`
	PercentOwnership *float64 `form:"percent_ownership"`
	Representative   *bool    `form:"representative"`
	Title            *string  `form:"title"`
}

// PersonVerificationDocumentParams represents the parameters available for the document verifying
// a person's identity.
type PersonVerificationDocumentParams struct {
	Back  *string `form:"back"`
	Front *string `form:"front"`
}

// PersonVerificationParams is used to represent parameters associated with a person's verification
// details.
type PersonVerificationParams struct {
	AdditionalDocument *PersonVerificationDocumentParams `form:"additional_document"`
	Document           *PersonVerificationDocumentParams `form:"document"`
}

// PersonParams is the set of parameters that can be used when creating or updating a person.
// For more details see https://stripe.com/docs/api#create_person.
type PersonParams struct {
	Params            `form:"*"`
	Account           *string                   `form:"-"` // Included in URL
	Address           *AccountAddressParams     `form:"address"`
	AddressKana       *AccountAddressParams     `form:"address_kana"`
	AddressKanji      *AccountAddressParams     `form:"address_kanji"`
	DOB               *DOBParams                `form:"dob"`
	Documents         *DocumentsParams          `form:"documents"`
	Email             *string                   `form:"email"`
	FirstName         *string                   `form:"first_name"`
	FirstNameKana     *string                   `form:"first_name_kana"`
	FirstNameKanji    *string                   `form:"first_name_kanji"`
	FullNameAliases   []*string                 `form:"full_name_aliases"`
	Gender            *string                   `form:"gender"`
	IDNumber          *string                   `form:"id_number"`
	LastName          *string                   `form:"last_name"`
	LastNameKana      *string                   `form:"last_name_kana"`
	LastNameKanji     *string                   `form:"last_name_kanji"`
	MaidenName        *string                   `form:"maiden_name"`
	Nationality       *string                   `form:"nationality"`
	PersonToken       *string                   `form:"person_token"`
	Phone             *string                   `form:"phone"`
	PoliticalExposure *string                   `form:"political_exposure"`
	Relationship      *RelationshipParams       `form:"relationship"`
	SSNLast4          *string                   `form:"ssn_last_4"`
	Verification      *PersonVerificationParams `form:"verification"`
}

// RelationshipListParams is used to filter persons by the relationship
type RelationshipListParams struct {
	Director       *bool `form:"director"`
	Executive      *bool `form:"executive"`
	Owner          *bool `form:"owner"`
	Representative *bool `form:"representative"`
}

// PersonListParams is the set of parameters that can be used when listing persons.
// For more detail see https://stripe.com/docs/api#list_persons.
type PersonListParams struct {
	ListParams   `form:"*"`
	Account      *string                 `form:"-"` // Included in URL
	Relationship *RelationshipListParams `form:"relationship"`
}

// DOB represents a Person's date of birth.
type DOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type PersonFutureRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Fields that are `currently_due` and need to be collected again because validation or verification failed.
type PersonFutureRequirementsError struct {
	Code        string `json:"code"`
	Reason      string `json:"reason"`
	Requirement string `json:"requirement"`
}

// Information about the upcoming new requirements for this person, including what information needs to be collected, and by when.
type PersonFutureRequirements struct {
	Alternatives        []*PersonFutureRequirementsAlternative `json:"alternatives"`
	CurrentlyDue        []string                               `json:"currently_due"`
	Errors              []*PersonFutureRequirementsError       `json:"errors"`
	EventuallyDue       []string                               `json:"eventually_due"`
	PastDue             []string                               `json:"past_due"`
	PendingVerification []string                               `json:"pending_verification"`
}

// Relationship represents how the Person relates to the business.
type Relationship struct {
	Director         bool    `json:"director"`
	Executive        bool    `json:"executive"`
	Owner            bool    `json:"owner"`
	PercentOwnership float64 `json:"percent_ownership"`
	Representative   bool    `json:"representative"`
	Title            string  `json:"title"`
}

// Fields that are due and can be satisfied by providing the corresponding alternative fields instead.
type PersonRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}

// Requirements represents what's missing to verify a Person.
type Requirements struct {
	Alternatives        []*PersonRequirementsAlternative `json:"alternatives"`
	CurrentlyDue        []string                         `json:"currently_due"`
	Errors              []*AccountRequirementsError      `json:"errors"`
	EventuallyDue       []string                         `json:"eventually_due"`
	PastDue             []string                         `json:"past_due"`
	PendingVerification []string                         `json:"pending_verification"`
}

// PersonVerificationDocument represents the documents associated with a Person.
type PersonVerificationDocument struct {
	Back        *File                           `json:"back"`
	Details     string                          `json:"details"`
	DetailsCode VerificationDocumentDetailsCode `json:"details_code"`
	Front       *File                           `json:"front"`
}

// PersonVerification is the structure for a person's verification details.
type PersonVerification struct {
	AdditionalDocument *PersonVerificationDocument   `json:"additional_document"`
	Details            string                        `json:"details"`
	DetailsCode        PersonVerificationDetailsCode `json:"details_code"`
	Document           *PersonVerificationDocument   `json:"document"`
	Status             IdentityVerificationStatus    `json:"status"`
}

// Person is the resource representing a Stripe person.
// For more details see https://stripe.com/docs/api#persons.
type Person struct {
	APIResource
	Account            string                    `json:"account"`
	Address            *AccountAddress           `json:"address"`
	AddressKana        *AccountAddress           `json:"address_kana"`
	AddressKanji       *AccountAddress           `json:"address_kanji"`
	Created            int64                     `json:"created"`
	Deleted            bool                      `json:"deleted"`
	DOB                *DOB                      `json:"dob"`
	Email              string                    `json:"email"`
	FirstName          string                    `json:"first_name"`
	FirstNameKana      string                    `json:"first_name_kana"`
	FirstNameKanji     string                    `json:"first_name_kanji"`
	FullNameAliases    []string                  `json:"full_name_aliases"`
	FutureRequirements *PersonFutureRequirements `json:"future_requirements"`
	Gender             string                    `json:"gender"`
	ID                 string                    `json:"id"`
	IDNumberProvided   bool                      `json:"id_number_provided"`
	LastName           string                    `json:"last_name"`
	LastNameKana       string                    `json:"last_name_kana"`
	LastNameKanji      string                    `json:"last_name_kanji"`
	MaidenName         string                    `json:"maiden_name"`
	Metadata           map[string]string         `json:"metadata"`
	Nationality        string                    `json:"nationality"`
	Object             string                    `json:"object"`
	Phone              string                    `json:"phone"`
	PoliticalExposure  PersonPoliticalExposure   `json:"political_exposure"`
	Relationship       *Relationship             `json:"relationship"`
	Requirements       *Requirements             `json:"requirements"`
	SSNLast4Provided   bool                      `json:"ssn_last_4_provided"`
	Verification       *PersonVerification       `json:"verification"`
}

// PersonList is a list of persons as retrieved from a list endpoint.
type PersonList struct {
	APIResource
	ListMeta
	Data []*Person `json:"data"`
}

// UnmarshalJSON handles deserialization of a Person.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *Person) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type person Person
	var v person
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = Person(v)
	return nil
}
