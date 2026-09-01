package awsvpclattice


// Experimental.
type AwsVpclatticeListener_TargetGroupsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener#target_group_identifier AwsVpclatticeListener#target_group_identifier}.
	// Experimental.
	TargetGroupIdentifier *string `field:"optional" json:"targetGroupIdentifier" yaml:"targetGroupIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener#weight AwsVpclatticeListener#weight}.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

