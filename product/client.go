//
//
// File generated from our OpenAPI spec
//
//

// Package product provides the /products APIs
package product

import (
	"net/http"

	stripe "github.com/HRInnovationLab/stripe-go/v72"
	"github.com/HRInnovationLab/stripe-go/v72/form"
)

// Client is used to invoke /products APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new product.
func New(params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().New(params)
}

// New creates a new product.
func (c Client) New(params *stripe.ProductParams) (*stripe.Product, error) {
	product := &stripe.Product{}
	err := c.B.Call(http.MethodPost, "/v1/products", c.Key, params, product)
	return product, err
}

// Get returns the details of a product.
func Get(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().Get(id, params)
}

// Get returns the details of a product.
func (c Client) Get(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	path := stripe.FormatURLPath("/v1/products/%s", id)
	product := &stripe.Product{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, product)
	return product, err
}

// Update updates a product's properties.
func Update(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().Update(id, params)
}

// Update updates a product's properties.
func (c Client) Update(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	path := stripe.FormatURLPath("/v1/products/%s", id)
	product := &stripe.Product{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, product)
	return product, err
}

// Del removes a product.
func Del(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	return getC().Del(id, params)
}

// Del removes a product.
func (c Client) Del(id string, params *stripe.ProductParams) (*stripe.Product, error) {
	path := stripe.FormatURLPath("/v1/products/%s", id)
	product := &stripe.Product{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, product)
	return product, err
}

// List returns a list of products.
func List(params *stripe.ProductListParams) *Iter {
	return getC().List(params)
}

// List returns a list of products.
func (c Client) List(listParams *stripe.ProductListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
		list := &stripe.ProductList{}
		err := c.B.CallRaw(http.MethodGet, "/v1/products", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list, err
	})}
}

// Iter is an iterator for products.
type Iter struct {
	*stripe.Iter
}

// Product returns the product which the iterator is currently pointing to.
func (i *Iter) Product() *stripe.Product {
	return i.Current().(*stripe.Product)
}

// ProductList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ProductList() *stripe.ProductList {
	return i.List().(*stripe.ProductList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
