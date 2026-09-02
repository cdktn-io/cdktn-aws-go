package awsvpclattice


// Experimental.
type TfListenerRule_ForwardProperty struct {
	// target_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpclattice_listener_rule#target_groups TfListenerRule#target_groups}
	// Experimental.
	TargetGroups interface{} `field:"required" json:"targetGroups" yaml:"targetGroups"`
}

