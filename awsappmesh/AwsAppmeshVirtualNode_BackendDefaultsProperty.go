package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_BackendDefaultsProperty struct {
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#client_policy AwsAppmeshVirtualNode#client_policy}
	// Experimental.
	ClientPolicy *AwsAppmeshVirtualNode_SpecBackendDefaultsClientPolicyProperty `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

