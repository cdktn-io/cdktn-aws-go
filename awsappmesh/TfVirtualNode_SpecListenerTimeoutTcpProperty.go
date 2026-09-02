package awsappmesh


// Experimental.
type TfVirtualNode_SpecListenerTimeoutTcpProperty struct {
	// idle block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#idle TfVirtualNode#idle}
	// Experimental.
	Idle *TfVirtualNode_SpecListenerTimeoutTcpIdleProperty `field:"optional" json:"idle" yaml:"idle"`
}

