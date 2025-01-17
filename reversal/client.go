//
//
// File generated from our OpenAPI spec
//
//

// Package reversal provides the /transfers/{id}/reversals APIs
package reversal

import (
	"fmt"
	"net/http"

	stripe "github.com/HRInnovationLab/stripe-go/v72"
	"github.com/HRInnovationLab/stripe-go/v72/form"
)

// Client is used to invoke /transfers/{id}/reversals APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new reversal.
func New(params *stripe.ReversalParams) (*stripe.Reversal, error) {
	return getC().New(params)
}

// New creates a new reversal.
func (c Client) New(params *stripe.ReversalParams) (*stripe.Reversal, error) {
	path := stripe.FormatURLPath(
		"/v1/transfers/%s/reversals",
		stripe.StringValue(params.Transfer),
	)
	reversal := &stripe.Reversal{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reversal)
	return reversal, err
}

// Get returns the details of a reversal.
func Get(id string, params *stripe.ReversalParams) (*stripe.Reversal, error) {
	return getC().Get(id, params)
}

// Get returns the details of a reversal.
func (c Client) Get(id string, params *stripe.ReversalParams) (*stripe.Reversal, error) {
	if params == nil {
		return nil, fmt.Errorf(
			"params cannnot be nil, and params.Transfer must be set",
		)
	}
	path := stripe.FormatURLPath(
		"/v1/transfers/%s/reversals/%s",
		stripe.StringValue(params.Transfer),
		id,
	)
	reversal := &stripe.Reversal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, reversal)
	return reversal, err
}

// Update updates a reversal's properties.
func Update(id string, params *stripe.ReversalParams) (*stripe.Reversal, error) {
	return getC().Update(id, params)
}

// Update updates a reversal's properties.
func (c Client) Update(id string, params *stripe.ReversalParams) (*stripe.Reversal, error) {
	path := stripe.FormatURLPath(
		"/v1/transfers/%s/reversals/%s",
		stripe.StringValue(params.Transfer),
		id,
	)
	reversal := &stripe.Reversal{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, reversal)
	return reversal, err
}

// List returns a list of reversals.
func List(params *stripe.ReversalListParams) *Iter {
	return getC().List(params)
}

// List returns a list of reversals.
func (c Client) List(listParams *stripe.ReversalListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/transfers/%s/reversals",
		stripe.StringValue(listParams.Transfer),
	)
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.ReversalList{}
		err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for reversals.
type Iter struct {
	*stripe.Iter
}

// Reversal returns the reversal which the iterator is currently pointing to.
func (i *Iter) Reversal() *stripe.Reversal {
	return i.Current().(*stripe.Reversal)
}

// ReversalList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ReversalList() *stripe.ReversalList {
	return i.List().(*stripe.ReversalList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
