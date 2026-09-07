package billing


// Experimental.
type AwsView_TimeRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#begin_date_inclusive AwsView#begin_date_inclusive}.
	// Experimental.
	BeginDateInclusive *string `field:"required" json:"beginDateInclusive" yaml:"beginDateInclusive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/billing_view#end_date_inclusive AwsView#end_date_inclusive}.
	// Experimental.
	EndDateInclusive *string `field:"required" json:"endDateInclusive" yaml:"endDateInclusive"`
}

