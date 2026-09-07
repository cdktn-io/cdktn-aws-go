package appmesh


// Experimental.
type AwsVirtualNode_SpecBackendDefaultsClientPolicyTlsCertificateFileProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#certificate_chain AwsVirtualNode#certificate_chain}.
	// Experimental.
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#private_key AwsVirtualNode#private_key}.
	// Experimental.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
}

