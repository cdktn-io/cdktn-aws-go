package awsappmesh


// Experimental.
type TfVirtualNode_SpecBackendVirtualServiceClientPolicyTlsValidationTrustAcmProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_node#certificate_authority_arns TfVirtualNode#certificate_authority_arns}.
	// Experimental.
	CertificateAuthorityArns *[]*string `field:"required" json:"certificateAuthorityArns" yaml:"certificateAuthorityArns"`
}

