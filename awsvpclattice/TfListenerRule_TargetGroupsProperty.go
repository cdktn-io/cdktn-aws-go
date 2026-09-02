package awsvpclattice


// Experimental.
type TfListenerRule_TargetGroupsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#target_group_identifier TfListenerRule#target_group_identifier}.
	// Experimental.
	TargetGroupIdentifier *string `field:"required" json:"targetGroupIdentifier" yaml:"targetGroupIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#weight TfListenerRule#weight}.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

