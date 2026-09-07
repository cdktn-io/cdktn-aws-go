package appmesh


// Experimental.
type AwsVirtualNode_BackendDefaultsProperty struct {
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#client_policy AwsVirtualNode#client_policy}
	// Experimental.
	ClientPolicy *AwsVirtualNode_SpecBackendDefaultsClientPolicyProperty `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

