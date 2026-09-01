package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecBackendVirtualServiceClientPolicyTlsCertificateFileProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#certificate_chain AwsAppmeshVirtualNode#certificate_chain}.
	// Experimental.
	CertificateChain *string `field:"required" json:"certificateChain" yaml:"certificateChain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#private_key AwsAppmeshVirtualNode#private_key}.
	// Experimental.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
}

