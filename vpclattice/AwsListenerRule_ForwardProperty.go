package vpclattice


// Experimental.
type AwsListenerRule_ForwardProperty struct {
	// target_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#target_groups AwsListenerRule#target_groups}
	// Experimental.
	TargetGroups interface{} `field:"required" json:"targetGroups" yaml:"targetGroups"`
}

