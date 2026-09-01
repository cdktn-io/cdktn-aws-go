package awsvpclattice


// Experimental.
type AwsVpclatticeListenerRule_ForwardProperty struct {
	// target_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#target_groups AwsVpclatticeListenerRule#target_groups}
	// Experimental.
	TargetGroups interface{} `field:"required" json:"targetGroups" yaml:"targetGroups"`
}

