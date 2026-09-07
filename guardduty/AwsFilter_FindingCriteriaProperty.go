package guardduty


// Experimental.
type AwsFilter_FindingCriteriaProperty struct {
	// criterion block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/guardduty_filter#criterion AwsFilter#criterion}
	// Experimental.
	Criterion interface{} `field:"required" json:"criterion" yaml:"criterion"`
}

