package accountlink

import (
	"testing"

	stripe "github.com/HRInnovationLab/stripe-go/v72"
	_ "github.com/HRInnovationLab/stripe-go/v72/testing"
	assert "github.com/stretchr/testify/require"
)

func TestAccountLinkNew(t *testing.T) {
	params := &stripe.AccountLinkParams{
		Account:    stripe.String("acct_123"),
		Collect:    stripe.String(string(stripe.AccountLinkCollectCurrentlyDue)),
		RefreshURL: stripe.String("https://stripe.com/refresh"),
		ReturnURL:  stripe.String("https://stripe.com/return"),
		Type:       stripe.String(string(stripe.AccountLinkTypeAccountOnboarding)),
	}
	link, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, link)
}
