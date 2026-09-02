package awsssmcontacts


// Experimental.
type TfPlan_StageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#duration_in_minutes TfPlan#duration_in_minutes}.
	// Experimental.
	DurationInMinutes *float64 `field:"required" json:"durationInMinutes" yaml:"durationInMinutes"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#target TfPlan#target}
	// Experimental.
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

