package awsappmesh


// Experimental.
type TfVirtualNode_BackendDefaultsProperty struct {
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#client_policy TfVirtualNode#client_policy}
	// Experimental.
	ClientPolicy *TfVirtualNode_SpecBackendDefaultsClientPolicyProperty `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

