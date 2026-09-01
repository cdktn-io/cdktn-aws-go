package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_BackendDefaultsProperty struct {
	// client_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#client_policy AwsAppmeshVirtualGateway#client_policy}
	// Experimental.
	ClientPolicy *AwsAppmeshVirtualGateway_ClientPolicyProperty `field:"optional" json:"clientPolicy" yaml:"clientPolicy"`
}

