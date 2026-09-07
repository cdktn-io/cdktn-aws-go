package appmesh


// Experimental.
type AwsVirtualGateway_ClientPolicyProperty struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#tls AwsVirtualGateway#tls}
	// Experimental.
	Tls *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

