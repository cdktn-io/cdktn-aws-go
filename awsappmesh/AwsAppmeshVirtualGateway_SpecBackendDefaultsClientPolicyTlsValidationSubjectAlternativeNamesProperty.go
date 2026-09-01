package awsappmesh


// Experimental.
type AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesProperty struct {
	// match block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appmesh_virtual_gateway#match AwsAppmeshVirtualGateway#match}
	// Experimental.
	Match *AwsAppmeshVirtualGateway_SpecBackendDefaultsClientPolicyTlsValidationSubjectAlternativeNamesMatchProperty `field:"required" json:"match" yaml:"match"`
}

