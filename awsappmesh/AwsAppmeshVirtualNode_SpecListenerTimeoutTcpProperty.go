package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecListenerTimeoutTcpProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#idle AwsAppmeshVirtualNode#idle}
	// Experimental.
	Idle *AwsAppmeshVirtualNode_SpecListenerTimeoutTcpIdleProperty `field:"optional" json:"idle" yaml:"idle"`
}

