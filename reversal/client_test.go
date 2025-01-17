package reversal

import (
	"testing"

	stripe "github.com/HRInnovationLab/stripe-go/v72"
	_ "github.com/HRInnovationLab/stripe-go/v72/testing"
	assert "github.com/stretchr/testify/require"
)

func TestReversalGet(t *testing.T) {
	reversal, err := Get("trr_123", &stripe.ReversalParams{
		Transfer: stripe.String("tr_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}

func TestReversalList(t *testing.T) {
	i := List(&stripe.ReversalListParams{
		Transfer: stripe.String("tr_123"),
	})

	// Verify that we can get at least one reversal
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Reversal())
	assert.NotNil(t, i.ReversalList())
}

func TestReversalNew(t *testing.T) {
	reversal, err := New(&stripe.ReversalParams{
		Amount:   stripe.Int64(123),
		Transfer: stripe.String("tr_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}

func TestReversalUpdate(t *testing.T) {
	reversal, err := Update("trr_123", &stripe.ReversalParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
		Transfer: stripe.String("tr_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reversal)
}
