package awsssmcontacts


// Experimental.
type AwsSsmcontactsPlan_StageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#duration_in_minutes AwsSsmcontactsPlan#duration_in_minutes}.
	// Experimental.
	DurationInMinutes *float64 `field:"required" json:"durationInMinutes" yaml:"durationInMinutes"`
	// target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#target AwsSsmcontactsPlan#target}
	// Experimental.
	Target interface{} `field:"optional" json:"target" yaml:"target"`
}

