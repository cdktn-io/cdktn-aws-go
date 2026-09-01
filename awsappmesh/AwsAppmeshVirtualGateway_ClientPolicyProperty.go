package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_ClientPolicyProperty struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#tls AwsAppmeshVirtualGateway#tls}
	// Experimental.
	Tls *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

