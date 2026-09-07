package appmesh


// Experimental.
type AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationTrustAcmProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#certificate_authority_arns AwsVirtualGateway#certificate_authority_arns}.
	// Experimental.
	CertificateAuthorityArns *[]*string `field:"required" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

