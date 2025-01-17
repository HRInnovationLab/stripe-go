package valuelistitem

import (
	"testing"

	stripe "github.com/HRInnovationLab/stripe-go/v72"
	_ "github.com/HRInnovationLab/stripe-go/v72/testing"
	assert "github.com/stretchr/testify/require"
)

func TestRadarValueListItemDel(t *testing.T) {
	vli, err := Del("rsli_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, vli)
}

func TestRadarValueListItemGet(t *testing.T) {
	vli, err := Get("rsli_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, vli)
}

func TestRadarValueListItemList(t *testing.T) {
	i := List(&stripe.RadarValueListItemListParams{
		RadarValueList: stripe.String("rsl_123"),
	})

	// Verify that we can get at least one value list item
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.RadarValueListItem())
	assert.NotNil(t, i.RadarValueListItemList())
}

func TestRadarValueListItemNew(t *testing.T) {
	vli, err := New(&stripe.RadarValueListItemParams{
		Value:          stripe.String("value"),
		RadarValueList: stripe.String("rsl_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, vli)
}
