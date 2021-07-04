package sourcetransaction

import (
	"testing"

	stripe "github.com/HRInnovationLab/stripe-go/v72"
	_ "github.com/HRInnovationLab/stripe-go/v72/testing"
	assert "github.com/stretchr/testify/require"
)

func TestSourceTransactionList(t *testing.T) {
	i := List(&stripe.SourceTransactionListParams{
		Source: stripe.String("src_123"),
	})

	// Verify that we can get at least one transaction
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SourceTransaction())
	assert.NotNil(t, i.SourceTransactionList())
}
