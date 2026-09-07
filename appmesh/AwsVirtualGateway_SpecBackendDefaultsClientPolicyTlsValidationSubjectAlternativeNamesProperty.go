package appmesh


// Experimental.
type AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#match AwsVirtualGateway#match}
	// Experimental.
	Match *AwsVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty `field:"required" json:"match" yaml:"match"`
}

