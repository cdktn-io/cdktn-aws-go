package awssecurityhub


// Experimental.
type TfInsight_FiltersLastObservedAtDateRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#unit TfInsight#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securityhub_insight#value TfInsight#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

