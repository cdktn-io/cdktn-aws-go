package appmesh


// Experimental.
type AwsVirtualNode_SpecListenerTimeoutTcpProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#idle AwsVirtualNode#idle}
	// Experimental.
	Idle *AwsVirtualNode_SpecListenerTimeoutTcpIdleProperty `field:"optional" json:"idle" yaml:"idle"`
}

