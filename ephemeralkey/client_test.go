package ephemeralkey

import (
	"testing"

	stripe "github.com/HRInnovationLab/stripe-go/v72"
	_ "github.com/HRInnovationLab/stripe-go/v72/testing"
	assert "github.com/stretchr/testify/require"
)

func TestEphemeralKeyDel(t *testing.T) {
	key, err := Del("ephkey_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, key)
}

func TestEphemeralKeyNew(t *testing.T) {
	key, err := New(&stripe.EphemeralKeyParams{
		Customer:      stripe.String("cus_123"),
		StripeVersion: stripe.String(stripe.APIVersion),
	})
	assert.Nil(t, err)
	assert.NotNil(t, key)
}
