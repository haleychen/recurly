package recurly

import (
	"fmt"
	"net/http"
)

var _ BillingService = &billingImpl{}

// billingImpl handles all interaction with the billing info portion
// of the recurly API.
type billingImpl struct {
	client *Client
}

// NewBillingImpl returns a new instance of billingImpl.
func NewBillingImpl(client *Client) *billingImpl {
	return &billingImpl{client: client}
}

// Get returns only the account's current billing information.
// https://docs.recurly.com/api/billing-info#lookup-billing-info
func (s *billingImpl) Get(accountCode string) (*Response, *Billing, error) {
	action := fmt.Sprintf("accounts/%s/billing_info", accountCode)
	req, err := s.client.newRequest("GET", action, nil, nil)
	if err != nil {
		return nil, nil, err
	}

	var dst Billing
	resp, err := s.client.do(req, &dst)
	if err != nil || resp.StatusCode >= http.StatusBadRequest {
		return resp, nil, err
	}

	return resp, &dst, err
}

// Create creates the account's billing information with credit card or
// bank account info. It is recommended you use recurly.js and a token with the CreateWithToken
// method instead.
// https://dev.recurly.com/docs/create-an-accounts-billing-info-credit-card
// https://dev.recurly.com/docs/create-an-accounts-billing-info-bank-account
func (s *billingImpl) Create(accountCode string, b Billing) (*Response, *Billing, error) {
	action := fmt.Sprintf("accounts/%s/billing_info", accountCode)
	req, err := s.client.newRequest("POST", action, nil, b)
	if err != nil {
		return nil, nil, err
	}

	var dst Billing
	resp, err := s.client.do(req, &dst)

	return resp, &dst, err
}

// CreateWithToken creates an account's billing information using a token
// generated by Recurly.js. Returns the account's created Billing Information.
// https://docs.recurly.com/api/billing-info#create-billing-info-token
func (s *billingImpl) CreateWithToken(accountCode string, token string) (*Response, *Billing, error) {
	action := fmt.Sprintf("accounts/%s/billing_info", accountCode)
	req, err := s.client.newRequest("POST", action, nil, Billing{Token: token})
	if err != nil {
		return nil, nil, err
	}

	var dst Billing
	resp, err := s.client.do(req, &dst)

	return resp, &dst, err
}

// Update updates the account's billing information with credit card or
// bank info.It is recommended you use recurly.js and a token with the
// UpdateWithToken method instead.
// https://dev.recurly.com/docs/update-an-accounts-billing-info-credit-card
// https://dev.recurly.com/docs/update-an-accounts-billing-info-bank-account
func (s *billingImpl) Update(accountCode string, b Billing) (*Response, *Billing, error) {
	// Create clean billing object with write-only fields to avoid errors
	// like sending additional/unknown/read-only fields.
	clean := Billing{
		FirstName:         b.FirstName,
		LastName:          b.LastName,
		Address:           b.Address,
		Address2:          b.Address2,
		City:              b.City,
		State:             b.State,
		Zip:               b.Zip,
		Country:           b.Country,
		Phone:             b.Phone,
		VATNumber:         b.VATNumber,
		IPAddress:         b.IPAddress,
		Number:            b.Number,
		Month:             b.Month,
		Year:              b.Year,
		VerificationValue: b.VerificationValue,
		NameOnAccount:     b.NameOnAccount,
		RoutingNumber:     b.RoutingNumber,
		AccountNumber:     b.AccountNumber,
		AccountType:       b.AccountType,
	}

	action := fmt.Sprintf("accounts/%s/billing_info", accountCode)
	req, err := s.client.newRequest("PUT", action, nil, clean)
	if err != nil {
		return nil, nil, err
	}

	var dst Billing
	resp, err := s.client.do(req, &dst)

	return resp, &dst, err
}

// UpdateWithToken updates an account's billing information using a token
// generated by Recurly.js. Returns the account's created Billing Information.
// https://docs.recurly.com/api/billing-info#update-billing-info-token
func (s *billingImpl) UpdateWithToken(accountCode string, token string) (*Response, *Billing, error) {
	action := fmt.Sprintf("accounts/%s/billing_info", accountCode)
	req, err := s.client.newRequest("PUT", action, nil, Billing{Token: token})
	if err != nil {
		return nil, nil, err
	}

	var dst Billing
	resp, err := s.client.do(req, &dst)

	return resp, &dst, err
}

// Clear removes any stored billing information for an account. If the account
// has a subscription, the renewal will go into past due unless you update the
// billing info before the renewal occurs.
// https://docs.recurly.com/api/billing-info#clear-billing-info
func (s *billingImpl) Clear(accountCode string) (*Response, error) {
	action := fmt.Sprintf("accounts/%s/billing_info", accountCode)
	req, err := s.client.newRequest("DELETE", action, nil, nil)
	if err != nil {
		return nil, err
	}

	return s.client.do(req, nil)
}