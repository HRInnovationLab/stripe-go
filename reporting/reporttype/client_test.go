package reporttype

import (
	"testing"

	stripe "github.com/HRInnovationLab/stripe-go/v72"
	_ "github.com/HRInnovationLab/stripe-go/v72/testing"
	assert "github.com/stretchr/testify/require"
)

func TestReportTestGet(t *testing.T) {
	reporttype, err := Get("activity.summary.1", nil)
	assert.Nil(t, err)
	assert.NotNil(t, reporttype)
	assert.Equal(t, "reporting.report_type", reporttype.Object)
}

func TestReportTestList(t *testing.T) {
	i := List(&stripe.ReportTypeListParams{})

	// Verify that we can get at least one reporttype
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.ReportType())
	assert.Equal(t, "reporting.report_type", i.ReportType().Object)
	assert.NotNil(t, i.ReportTypeList())
}
