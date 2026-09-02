package awsappmesh


// Experimental.
type TfVirtualGateway_SpecBackendDefaultsClientPolicyTlsCertificateFileProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#certificate_chain TfVirtualGateway#certificate_chain}.
	// Experimental.
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#private_key TfVirtualGateway#private_key}.
	// Experimental.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
}

