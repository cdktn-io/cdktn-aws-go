package appmesh


// Experimental.
type AwsVirtualNode_SpecBackendDefaultsClientPolicyProperty struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#tls AwsVirtualNode#tls}
	// Experimental.
	Tls *AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

