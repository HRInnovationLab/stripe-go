package usagerecord

import (
	"testing"
	"time"

	stripe "github.com/HRInnovationLab/stripe-go/v72"
	_ "github.com/HRInnovationLab/stripe-go/v72/testing"
	assert "github.com/stretchr/testify/require"
)

func TestUsageRecordNew(t *testing.T) {
	now := int64(time.Now().Unix())
	usageRecord, err := New(&stripe.UsageRecordParams{
		Quantity:         stripe.Int64(123),
		Timestamp:        stripe.Int64(now),
		Action:           stripe.String(stripe.UsageRecordActionIncrement),
		SubscriptionItem: stripe.String("si_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, usageRecord)
}
