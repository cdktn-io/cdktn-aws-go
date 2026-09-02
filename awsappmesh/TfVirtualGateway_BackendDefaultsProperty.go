package awsappmesh


// Experimental.
type TfVirtualGateway_BackendDefaultsProperty struct {
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#client_policy TfVirtualGateway#client_policy}
	// Experimental.
	ClientPolicy *TfVirtualGateway_ClientPolicyProperty `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

