package awsappmesh


// Experimental.
type AwsAppmeshVirtualNode_SpecBackendDefaultsClientPolicyTlsValidationTrustAcmProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#certificate_authority_arns AwsAppmeshVirtualNode#certificate_authority_arns}.
	// Experimental.
	CertificateAuthorityArns *[]*string `field:"required" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

