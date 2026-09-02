package awsappmesh


// Experimental.
type TfVirtualNode_SpecBackendDefaultsClientPolicyProperty struct {
	// tls block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#tls TfVirtualNode#tls}
	// Experimental.
	Tls *TfVirtualNode_SpecBackendDefaultsClientPolicyTlsProperty `field:"optional" json:"tls" yaml:"tls"`
}

